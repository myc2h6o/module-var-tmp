package validation

import (
	"fmt"
)

type logger struct{
}

func (logger) EmitError(error string) {
	fmt.Printf("[ERROR] %s\n", error)
}

func (logger) EmitWarning(error string) {
	fmt.Printf("[WARNING] %s\n", error)
}

func (logger) EmitInfo(error string) {
	fmt.Printf("[INFO] %s\n", error)
}