package validation

import (
	"github.com/hashicorp/hcl/v2"
	"path"
	"runtime"
	"testing"
)

func TestValidate_ValidModule(t *testing.T) {
	hclFiles := getHclFiles("valid_module", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if !isValid {
		t.Error("valid_module should be a valid module")
	}
}

func TestValidate_InvalidModuleResource(t *testing.T) {
	hclFiles := getHclFiles("invalid_module_resource", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("invalid_module should be an invalid module")
	}
}


func TestValidate_InvalidModuleDataSource(t *testing.T) {
	hclFiles := getHclFiles("invalid_module_data_source", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("invalid_module should be an invalid module")
	}
}

func TestValidate_IgnoredVariables(t *testing.T) {
	hclFiles := getHclFiles("invalid_module_resource", t)
	ignoredVariables := []string {
		"name_1",
		"address",
		"tag",
		"street_test",
		"condition_prop_true",
		"condition_prop_false",
	}

	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if !isValid {
		t.Error("variables are not ignored correctly")
	}
}

func getHclFiles(moduleName string, t *testing.T) map[string]*hcl.File {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "..", "test_modules", moduleName)
	moduleReader := NewReader()
	hclFiles, err := moduleReader.Read(modulePath)
	if err != nil {
		t.Error(err)
	}
	return hclFiles
}