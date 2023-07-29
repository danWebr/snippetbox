package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Valid returns true if no errors were recorded
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

// AddFieldError adds a new error to the FieldErrors map
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	// Only add message if no entry exists already
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// AddNonFieldError adds a new error that are not related to a specific form field
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

// CheckField adds an error message to FieldErrors map if validation check is not 'ok'
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank returns true if input string is not empty
func NotBlank(input string) bool {
	return strings.TrimSpace(input) != ""
}

// MinChars returns true if input string is at least n characters long
func MinChars(input string, n int) bool {
	return utf8.RuneCountInString(input) >= n
}

// MaxChars returns true if input string is not longer than n characters
func MaxChars(input string, n int) bool {
	return utf8.RuneCountInString(input) <= n
}

// PermittedInt returns true if value is in a list of permitted values
func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

// Matches returns true if input string matches a provided compiled regular expression pattern
func Matches(input string, rx *regexp.Regexp) bool {
	return rx.MatchString(input)
}
