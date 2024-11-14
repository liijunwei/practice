package validator

import (
	"regexp"
)

var EmailRegexp = regexp.MustCompile(`^[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$`)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, ok := v.Errors[key]; !ok {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) In(value string, list ...string) bool {
	for _, val := range list {
		if value == val {
			return true
		}
	}

	return false
}

func Match(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique(values []string) bool {
	uniqueVals := make(map[string]bool)

	for _, val := range values {
		uniqueVals[val] = true
	}

	return len(values) == len(uniqueVals)
}
