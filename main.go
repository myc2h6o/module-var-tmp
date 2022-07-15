package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"os"
	"strings"
)

func main() {
	diskFile, err := os.ReadFile("D:\\terraform-test\\aks data references resource GH 13152 15557\\modules\\RoleAssignment\\main.tf")
	if err != nil {
		fmt.Println(err)
		return
	}

	hclFile, _ := hclsyntax.ParseConfig(diskFile, "main.tf", hcl.InitialPos)
	body := hclFile.Body.(*hclsyntax.Body)

	for _, block := range body.Blocks {
		switch block.Type {
		case "data", "resource":
			for _, attribute := range block.Body.Attributes {
				// TODO nested property
				propertyName := attribute.Name
				for _, variable := range attribute.Expr.Variables() {
					if variable.RootName() == "var" {
						variableName := variable[1].(hcl.TraverseAttr).Name

						// Validate
						fmt.Println(variableName, propertyName)
						if strings.LastIndex(variableName, propertyName) == len(variableName) - len(propertyName) {
							fmt.Println("valid")
						} else {
							fmt.Println("invalid")
						}
					}
				}
			}
		default:
		}
	}
}
