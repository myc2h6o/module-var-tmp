package validation

import (
	"fmt"
)

type logger struct{
}

func (logger) EmitError(error string) {
	fmt.Printf("[ERROR] %s\n", error)
}