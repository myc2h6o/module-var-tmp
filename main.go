package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"os"
)

func main() {
	f, err := os.ReadFile("D:\\terraform-test\\aks data references resource GH 13152 15557\\modules\\RoleAssignment\\main.tf")
	if err != nil {
		fmt.Println(err)
		return
	}



	a, b := hclsyntax.ParseConfig(f, "main.tf", hcl.InitialPos)
fmt.Println(a)
	fmt.Println(b)
}
