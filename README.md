# tf-module-var-validation
The Validation SDK validates Terraform Module Variables. This SDK considers a valid module variable to be same as the property name or property name with prefix. Valid and invalid examples could be found at [test_modules](test_modules).

## Package Tests
```
go test ./validation -v -count 1
```

## Usage in PowerShell
```
go build
$env:TFLINT_LOG="ERROR"
.\tf-module-var-validation.exe -path .\test_modules\invalid_module_resource
$env:TFLINT_LOG="DEBUG"
.\tf-module-var-validation.exe -path .\test_modules\valid_module
```

## Example output
```
$env:TFLINT_LOG="ERROR"
.\tf-module-var-validation.exe -path .\test_modules\invalid_module_resource

2022/08/22 15:38:23 main.tf Line:3 Column:10 [ERROR] Property:name Variable:name_1 is invalid
2022/08/22 15:38:23 main.tf Line:18 Column:47 [ERROR] Property:condition_prop Variable:condition_prop_true is invalid
2022/08/22 15:38:23 main.tf Line:18 Column:73 [ERROR] Property:condition_prop Variable:condition_prop_false is invalid
2022/08/22 15:38:23 main.tf Line:14 Column:14 [ERROR] Property:street Variable:street_test is invalid
```
```
$env:TFLINT_LOG="DEBUG"
.\tf-module-var-validation.exe -path .\test_modules\valid_module

2022/08/22 15:38:26 data_source.tf Line:2 Column:10 [DEBUG] Property:name Variable:resource_group_name is valid
2022/08/22 15:38:26 main.tf Line:4 Column:14 [DEBUG] Property:location Variable:location is valid
2022/08/22 15:38:26 main.tf Line:5 Column:10 [DEBUG] Property:tags Variable:test_tags is valid
2022/08/22 15:38:26 main.tf Line:3 Column:10 [DEBUG] Property:name Variable:resource_group_name is valid
2022/08/22 15:38:26 main.tf Line:10 Column:10 [DEBUG] Property:name Variable:test_name is valid
2022/08/22 15:38:26 main.tf Line:16 Column:10 [DEBUG] Property:list Variable:test_list is valid
2022/08/22 15:38:26 main.tf Line:27 Column:18 [DEBUG] Skipping Property:ip_addresses
2022/08/22 15:38:26 main.tf Line:35 Column:21 [DEBUG] Skipping Property:string_template
2022/08/22 15:38:26 main.tf Line:13 Column:14 [DEBUG] Property:location Variable:location is valid
2022/08/22 15:38:26 main.tf Line:19 Column:12 [DEBUG] Property:region Variable:region is valid
2022/08/22 15:38:26 main.tf Line:41 Column:47 [DEBUG] Property:condition_prop Variable:true_condition_prop is valid
2022/08/22 15:38:26 main.tf Line:23 Column:14 [DEBUG] Property:street Variable:test_street is valid
```
