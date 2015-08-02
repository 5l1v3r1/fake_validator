package api

import (
	"errors"
	"strconv"
)

// Errors
var (
	ErrorOutOfRange = errors.New("Number out of range")
	ErrorWrongInput = errors.New("Wrong input")
)

// NumberValidator type
type NumberValidator struct {
	number string
}

// NewNumberValidator returns number validator
func NewNumberValidator(number string) Validator {
	return NumberValidator{number}
}

// Validate number
func (n NumberValidator) Validate() (bool, error) {
	num, err := strconv.ParseInt(n.number, 10, 64)
	if err != nil {
		return false, err
	}

	// Number
	if num < 0 || num > 100 {
		return false, ErrorOutOfRange
	}

	return true, nil
}
