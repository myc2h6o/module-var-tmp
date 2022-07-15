package main

import (
	"flag"
	"github.com/myc2h6o/module-var-tmp/validation"
	"os"
)

func main() {
	modulePath := flag.String("path", "invalid", "Module Path")
	flag.Parse()

	validator := validation.NewValidator(*modulePath)
	isValid := validator.Validate()

	if !isValid {
		os.Exit(1)
	}
}
