package validation

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"os"
	"path/filepath"
	"strings"
)

type reader struct {
}

func (reader) read (modulePath string) ([]*hcl.File, error) {
	var filePaths []string
	err := filepath.Walk(modulePath, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), ".tf") && info.Name() != "variables.tf" && info.Name() != "outputs.tf" {
				filePaths = append(filePaths, filePath)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	hclFiles := make([]*hcl.File, len(filePaths))
	for i, filePath := range filePaths {
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		hclFile, diags := hclsyntax.ParseConfig(bytes, filePath, hcl.InitialPos)
		if diags.HasErrors() {
			return nil, fmt.Errorf(diags.Error())
		}

		hclFiles[i] = hclFile
	}

	return hclFiles, nil
}