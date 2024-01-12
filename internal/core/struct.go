package core

import (
	"fmt"
	"reflect"
)

// validatorMap : Defines which validator to use based on the tag.
// If no validator is found, DefaultValidator is used.
// If you create a new Validator, add it to this map with the tag name as the key.
var validatorMap = map[string]Validator{
	"required": RequiredValidator{},
}

// DefaultValidator : does not perform any validations.
type DefaultValidator struct{}

// Validate : returns true and nil error.
func (v DefaultValidator) Validate(_ reflect.Value) (bool, error) {
	return true, nil
}

// RequiredValidator : Validates that the value is not empty.
type RequiredValidator struct{}

// Validate : returns true if the value is not empty, false otherwise.
func (v RequiredValidator) Validate(val reflect.Value) (bool, error) {
	if !val.IsValid() {
		return false, fmt.Errorf("field value for %s is required but missing", AppName)
	}

	// nolint:exhaustive
	switch val.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		if val.IsNil() {
			return false, fmt.Errorf("field value for %s is required but missing", AppName)
		}
	default:
		if val.IsZero() {
			return false, fmt.Errorf("field value for %s is required but missing", AppName)
		}
	}

	return true, nil
}

// Validator : Generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(reflect.Value) (bool, error)
}

// Validate : Validates the struct fields based on the tag.
func Validate(data interface{}) []error {
	errs := make([]error, 0)

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		// Fetch the tag value from the struct field
		tag := v.Type().Field(i).Tag.Get(AppName)

		// Skip if tag is empty or "-"
		if tag == "" || tag == "-" {
			continue
		}
		// Get the validator based on the tag
		val, ok := validatorMap[tag]
		if !ok {
			val = DefaultValidator{}
		}
		// Validate the field
		if valid, err := val.Validate(v.Field(i)); !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %w", v.Type().Field(i).Name, err))
		}
	}
	return errs
}
