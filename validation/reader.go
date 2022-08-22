package validation

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Reader struct {
}

func NewReader() Reader {
	return Reader{}
}

// Read gets all the .tf files under the modulePath (excluding sub-directories) and transform them into hclFile
func (Reader) Read(modulePath string) (map[string]*hcl.File, error) {
	files, err := ioutil.ReadDir(modulePath)
	if err != nil {
		return nil, err
	}

	hclFiles := make(map[string]*hcl.File)
	for _, file := range files {
		fileName := file.Name()
		if !file.IsDir() && strings.HasSuffix(fileName, ".tf") {
			bytes, err := ioutil.ReadFile(filepath.Join(modulePath, fileName))
			if err != nil {
				return nil, err
			}

			hclFile, diags := hclsyntax.ParseConfig(bytes, fileName, hcl.InitialPos)
			if diags.HasErrors() {
				return nil, fmt.Errorf(diags.Error())
			}

			hclFiles[fileName] = hclFile
		}
	}

	return hclFiles, nil
}
