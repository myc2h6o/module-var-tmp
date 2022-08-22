package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
	"github.com/myc2h6o/module-var-tmp/validation"
)

func main() {
	// Set log level
	logLevel := os.Getenv("TFLINT_LOG")
	if logLevel == "" {
		logLevel = "ERROR"
	}
	log.SetOutput(&logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(strings.ToUpper(logLevel)),
		Writer:   os.Stderr,
	})

	// Get root module path from argument
	modulePath := flag.String("path", "invalid", "Module Path")
	flag.Parse()

	// Get hcl files
	moduleReader := validation.NewReader()
	hclFiles, err := moduleReader.Read(*modulePath)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		os.Exit(1)
	}

	// [TODO] Read ignored variables from tflint annotations in variables.tf
	ignoredVariables := []string{
		"ignored_variable_0",
		"ignored_variable_1",
	}

	validator := validation.NewValidator(hclFiles, ignoredVariables)
	if !validator.Validate() {
		os.Exit(1)
	}
}
