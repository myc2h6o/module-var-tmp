# module-var-tmp

## Usage
```
go build
.\module-var-tmp.exe -path .\test
```

## Example output
```
[ERROR] test\main.tf:Line:16 Column:11; Property: tag; Variable: tag_a
[ERROR] test\main.tf:Line:23 Column:8; Property: n_d; Variable: nested_prop_n_d_invalid
[ERROR] test\sub-module\a.tf:Line:16 Column:11; Property: tag; Variable: tag_a
[ERROR] test\sub-module\a.tf:Line:23 Column:8; Property: n_d; Variable: nested_prop_n_d_invalid
```
