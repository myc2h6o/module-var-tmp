package validation

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Validator struct {
	hclFiles          map[string]*hcl.File
	ignoredVariables  map[string]struct{}
	builtInProperties map[string]struct{}
}

func NewValidator(hclFiles map[string]*hcl.File, ignoredVariables []string) Validator {
	validator := Validator{
		hclFiles:         hclFiles,
		ignoredVariables: map[string]struct{}{},
		builtInProperties: map[string]struct{}{
			// Built-in Properties are from https://www.terraform.io/language/resources/syntax#meta-arguments and https://www.terraform.io/language/resources/syntax#operation-timeouts
			"depends_on":  {},
			"count":       {},
			"for_each":    {},
			"provider":    {},
			"lifecycle":   {},
			"provisioner": {},
			"timeouts":    {},
		},
	}

	for _, variable := range ignoredVariables {
		validator.ignoredVariables[variable] = struct{}{}
	}

	return validator
}

// Validate returns true as successful
func (v Validator) Validate() bool {
	isValid := true

	for _, hclFile := range v.hclFiles {
		body := hclFile.Body.(*hclsyntax.Body)
		for _, block := range body.Blocks {
			switch block.Type {
			case "data", "resource":
				if !v.validateBody(block.Body) {
					isValid = false
				}
			default:
			}
		}
	}

	return isValid
}

func (v Validator) validateBody(body *hclsyntax.Body) bool {
	isValid := true
	for _, attribute := range body.Attributes {
		expression := attribute.Expr
		switch expression.(type) {
		case *hclsyntax.FunctionCallExpr:
			// Function call
		case *hclsyntax.ObjectConsExpr:
			// TypeMap
		case *hclsyntax.ConditionalExpr:
			if !v.validateExpression(expression.(*hclsyntax.ConditionalExpr).TrueResult, attribute.Name) {
				isValid = false
			}
			if !v.validateExpression(expression.(*hclsyntax.ConditionalExpr).FalseResult, attribute.Name) {
				isValid = false
			}
		case *hclsyntax.ScopeTraversalExpr:
			if !v.validateExpression(expression, attribute.Name) {
				isValid = false
			}
		default:
			v.writeLog(fmt.Sprintf("[DEBUG] Skipping Property:%s", attribute.Name), expression.Range())
		}
	}

	for _, block := range body.Blocks {
		if _, ok := v.builtInProperties[block.Type]; !ok {
			if !v.validateBody(block.Body) {
				isValid = false
			}
		}
	}

	return isValid
}

func (v Validator) validateExpression(expression hclsyntax.Expression, propertyName string) bool {
	if _, ok := v.builtInProperties[propertyName]; !ok {
		variables := expression.Variables()
		if len(variables) == 1 {
			variable := variables[0]
			if variable.RootName() == "var" {
				variableName := ""
				// Find last named variable in statement like a[0].b[1]
				for i := len(variable) - 1; i >= 0 && variableName == ""; i-- {
					switch variable[i].(type) {
					case hcl.TraverseAttr:
						variableName = variable[i].(hcl.TraverseAttr).Name
					default:
					}
				}

				if !v.validate(variableName, propertyName, variable[0].SourceRange()) {
					return false
				}
			}
		}
	}
	return true
}

func (v Validator) validate(variableName string, propertyName string, location hcl.Range) bool {
	if variableName == "" || propertyName == "" {
		v.writeLog("[ERROR] Variable or Property name is empty", location)
	}

	if index := strings.LastIndex(variableName, propertyName); index == -1 || index != len(variableName)-len(propertyName) {
		if _, ok := v.ignoredVariables[variableName]; ok {
			v.writeLog(fmt.Sprintf("[INFO] Variable:%s is ignored", variableName), location)
			return true
		} else {
			v.writeLog(fmt.Sprintf("[ERROR] Property:%s Variable:%s is invalid", propertyName, variableName), location)
			return false
		}
	}

	v.writeLog(fmt.Sprintf("[DEBUG] Property:%s Variable:%s is valid", propertyName, variableName), location)
	return true
}

func (v Validator) writeLog(message string, location hcl.Range) {
	log.Printf("%s Line:%d Column:%d %s", location.Filename, location.Start.Line, location.Start.Column, message)
}
