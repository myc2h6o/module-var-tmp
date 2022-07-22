package main

import (
	"flag"
	"github.com/myc2h6o/module-var-tmp/validation"
	"os"
)

func main() {
	modulePath := flag.String("path", "invalid", "Module Path")
	flag.Parse()

	ignoredVariables := []string {
		"nb_data_disk",
		"name_prefix",
	}

	validator := validation.NewValidator(*modulePath, ignoredVariables)
	isValid := validator.Validate()

	if !isValid {
		os.Exit(1)
	}
}
