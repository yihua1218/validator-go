package validator

import "testing"

var v Validator
var e error

// TestCreate Create Validator
func TestCreate(t *testing.T) {
	v, e = New("samples/validator.json")

	if e != nil {
		t.Error("Cannot create validator:", e)
	}

	v.Test()
}
