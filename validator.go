package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	govalid "github.com/gima/govalid/v1"
)

// Validator is a struct for validation
type Validator struct {
	File   string
	JSON   map[string]interface{}
	Schema govalid.Validator
}

var (
	validator = Validator{
		File:   "",
		Schema: nil,
	}
)

// New is for creating a new validator
func New(filename string) Validator {
	validator.File = filename
	file, e := ioutil.ReadFile(validator.File)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}
	json.Unmarshal(file, &validator.JSON)
	return validator
}
