package validation

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"strings"
)

type Validator struct{
	path string
	logger logger
	ignoredVariables map[string]bool
	builtInProperties map[string]bool
}

func NewValidator(path string, ignoredVariables []string) Validator {
	validator := Validator{
		path:             path,
		logger:           logger{},
		ignoredVariables: map[string]bool{},
		builtInProperties: map[string]bool{
			"count":      true,
			"for_each":   true,
			"timeouts":   true,
			"lifecycle":  true,
			"depends_on": true,
		},
	}

	for _, variable := range ignoredVariables {
		validator.ignoredVariables[variable] = true
	}

	return validator
}

// Validate returns true as successful
func (v Validator) Validate() bool {
	moduleReader := reader{
		logger: v.logger,
	}

	hclFiles, err := moduleReader.read(v.path)
	if err != nil {
		v.logger.EmitError(err.Error())
		return false
	}

	isAllValid := true

	for _, hclFile := range hclFiles {
		body := hclFile.Body.(*hclsyntax.Body)
		for _, block := range body.Blocks {
			switch block.Type {
			case "data", "resource":
				isValid := v.validateBody(block.Body)
				if !isValid {
					isAllValid = false
				}
			default:
			}
		}
	}

	return isAllValid
}

func (v Validator) validateBody(body *hclsyntax.Body) bool {
	isAllValid := true
	for _, attribute := range body.Attributes {
		expression := attribute.Expr
		switch expression.(type) {
		// TODO other Expr types
		case *hclsyntax.FunctionCallExpr:
		case *hclsyntax.ConditionalExpr:
			isValid := v.validateExpression(expression.(*hclsyntax.ConditionalExpr).TrueResult, attribute.Name)
			if !isValid {
				isAllValid = false
			}

			isValid = v.validateExpression(expression.(*hclsyntax.ConditionalExpr).FalseResult, attribute.Name)
			if !isValid {
				isAllValid = false
			}
		default:
			isValid := v.validateExpression(expression, attribute.Name)
			if !isValid {
				isAllValid = false
			}
		}
	}

	for _, block := range body.Blocks {
		isValid := v.validateBody(block.Body)
		if !isValid {
			isAllValid = false
		}
	}

	return isAllValid
}

func (v Validator) validateExpression(expression hclsyntax.Expression, propertyName string) bool {
	isAllValid := true
	if _, ok := v.builtInProperties[propertyName]; !ok {
		for _, variable := range expression.Variables() {
			if variable.RootName() == "var" {
				variableName := ""
				for i := len(variable)-1; i >= 0; i-- {
					switch variable[i].(type) {
					case hcl.TraverseAttr:
						variableName = variable[i].(hcl.TraverseAttr).Name
						break
					default:
					}
				}

				isValid := v.validate(variableName, propertyName, variable[0].SourceRange())
				if !isValid {
					isAllValid = false
				}
			}
		}
	}
	return isAllValid
}

func (v Validator) validate(variableName string, propertyName string, sourceRange hcl.Range) bool {
	message := fmt.Sprintf("%s:Line:%d Column:%d; Property: %s; Variable: %s", sourceRange.Filename, sourceRange.Start.Line, sourceRange.Start.Column, propertyName, variableName)
	if variableName == "" || propertyName == "" {
		v.logger.EmitError(message)
	}

	if strings.LastIndex(variableName, propertyName) != len(variableName) - len(propertyName) {
		if _, ok := v.ignoredVariables[variableName]; ok {
			v.logger.EmitWarning(fmt.Sprintf("(Ignored): %s", message))
			return true
		}else {
			v.logger.EmitError(message)
			return false
		}
	}

	v.logger.EmitInfo(message)
	return true
}