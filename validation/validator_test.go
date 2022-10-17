package validation

import (
	"github.com/hashicorp/hcl/v2"
	"path"
	"runtime"
	"testing"
)

func TestValidate_ValidModule(t *testing.T) {
	hclFiles := getHclFiles("valid_combination", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if !isValid {
		t.Error("`valid_combination` should be a valid module")
	}
}

func TestValidate_InvalidModuleResource(t *testing.T) {
	hclFiles := getHclFiles("invalid_resource_combination", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_resource_combination` should be an invalid module")
	}
}

func TestValidate_InvalidModuleDataSource(t *testing.T) {
	hclFiles := getHclFiles("invalid_data_source", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_data_source` should be an invalid module")
	}
}

func TestValidate_IgnoredVariables(t *testing.T) {
	hclFiles := getHclFiles("invalid_resource_combination", t)
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
		t.Error("after ignoring invalid variables, `invalid_resource_combination` should be a valid module")
	}
}

func TestValidate_InvalidConditionFalse(t *testing.T) {
	hclFiles := getHclFiles("invalid_condition_false", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_condition_false` should be an invalid module")
	}
}

func TestValidate_invalidConditionTrue(t *testing.T) {
	hclFiles := getHclFiles("invalid_condition_true", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_condition_true` should be an invalid module")
	}
}

func TestValidate_invalidDifferentName(t *testing.T) {
	hclFiles := getHclFiles("invalid_different_name", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_different_name` should be an invalid module")
	}
}

func TestValidate_invalidList(t *testing.T) {
	hclFiles := getHclFiles("invalid_list", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_list` should be an invalid module")
	}
}

func TestValidate_invalidNameSuffix(t *testing.T) {
	hclFiles := getHclFiles("invalid_name_suffix", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_name_suffix` should be an invalid module")
	}
}

func TestValidateInvalidNested(t *testing.T) {
	hclFiles := getHclFiles("invalid_nested", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_nested` should be an invalid module")
	}
}

func TestValidate_invalidNestedList(t *testing.T) {
	hclFiles := getHclFiles("invalid_nested_list", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_nested_list` should be an invalid module")
	}
}

func TestValidate_invalidNestedProperty(t *testing.T) {
	hclFiles := getHclFiles("invalid_nested_property", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_nested_property` should be an invalid module")
	}
}

func TestValidate_invalidPartialName(t *testing.T) {
	hclFiles := getHclFiles("invalid_partial_name", t)
	var ignoredVariables []string
	validator := NewValidator(hclFiles, ignoredVariables)
	isValid := validator.Validate()
	if isValid {
		t.Error("`invalid_partial_name` should be an invalid module")
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