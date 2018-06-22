package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	govalid "github.com/gima/govalid/v1"
	"github.com/tidwall/gjson"
)

// Validator is a struct for validation
type Validator struct {
	File   string
	JSON   map[string]interface{}
	Schema govalid.Validator
	root   gjson.Result
}

var (
	validator = Validator{
		File:   "",
		Schema: nil,
	}
)

// New is for creating a new validator
func New(filename string) (Validator, error) {
	validator.File = filename
	file, e := ioutil.ReadFile(validator.File)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
	} else {
		validator.root = gjson.Parse(string(file))

		validator.root.ForEach(func(key, value gjson.Result) bool {
			fmt.Printf("key: %s\n", key)
			return true
		})

		json.Unmarshal(file, &validator.JSON)
	}
	return validator, e
}

// Test a public method for testing
func Test() {
	fmt.Printf("Public Test():\n")
}
