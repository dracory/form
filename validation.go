package form

import (
	"regexp"
	"strconv"
	"strings"
)

// ValidationError represents a single validation error for a field.
type ValidationError struct {
	Field   string
	Message string
}

// Validator is a function that validates a field value and returns an error message if invalid.
type Validator func(fieldName string, value string) *ValidationError

// ValidatorRequired returns a validator that checks if a value is non-empty.
func ValidatorRequired() Validator {
	return func(fieldName string, value string) *ValidationError {
		if strings.TrimSpace(value) == "" {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " is required",
			}
		}
		return nil
	}
}

// ValidatorMinLength returns a validator that checks if a value has at least minLength characters.
func ValidatorMinLength(minLength int) Validator {
	return func(fieldName string, value string) *ValidationError {
		if len(value) < minLength {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " must be at least " + strconv.Itoa(minLength) + " characters",
			}
		}
		return nil
	}
}

// ValidatorMaxLength returns a validator that checks if a value has at most maxLength characters.
func ValidatorMaxLength(maxLength int) Validator {
	return func(fieldName string, value string) *ValidationError {
		if len(value) > maxLength {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " must be at most " + strconv.Itoa(maxLength) + " characters",
			}
		}
		return nil
	}
}

// ValidatorMin returns a validator that checks if a numeric value is at least min.
func ValidatorMin(min float64) Validator {
	return func(fieldName string, value string) *ValidationError {
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " must be a valid number",
			}
		}
		if v < min {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " must be at least " + strconv.FormatFloat(min, 'f', -1, 64),
			}
		}
		return nil
	}
}

// ValidatorMax returns a validator that checks if a numeric value is at most max.
func ValidatorMax(max float64) Validator {
	return func(fieldName string, value string) *ValidationError {
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " must be a valid number",
			}
		}
		if v > max {
			return &ValidationError{
				Field:   fieldName,
				Message: fieldName + " must be at most " + strconv.FormatFloat(max, 'f', -1, 64),
			}
		}
		return nil
	}
}

// ValidatorPattern returns a validator that checks if a value matches a regex pattern.
func ValidatorPattern(pattern string, message string) Validator {
	re := regexp.MustCompile(pattern)
	return func(fieldName string, value string) *ValidationError {
		if value != "" && !re.MatchString(value) {
			msg := message
			if msg == "" {
				msg = fieldName + " has an invalid format"
			}
			return &ValidationError{
				Field:   fieldName,
				Message: msg,
			}
		}
		return nil
	}
}

// ValidatorEmail returns a validator that checks if a value is a valid email address.
func ValidatorEmail() Validator {
	return ValidatorPattern(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, "must be a valid email address")
}

// ValidatorURL returns a validator that checks if a value is a valid URL.
func ValidatorURL() Validator {
	return ValidatorPattern(`^https?://[^\s/$.?#].[^\s]*$`, "must be a valid URL")
}

// ValidatorIP returns a validator that checks if a value is a valid IPv4 address.
func ValidatorIP() Validator {
	return ValidatorPattern(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`, "must be a valid IP address")
}

// ValidatorUUID returns a validator that checks if a value is a valid UUID.
func ValidatorUUID() Validator {
	return ValidatorPattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`, "must be a valid UUID")
}

// ValidatorAlphaNumeric returns a validator that checks if a value contains only letters and numbers.
func ValidatorAlphaNumeric() Validator {
	return ValidatorPattern(`^[a-zA-Z0-9]+$`, "must contain only letters and numbers")
}

// ValidatorOneOf returns a validator that checks if a value is one of the allowed values.
func ValidatorOneOf(allowed ...string) Validator {
	return func(fieldName string, value string) *ValidationError {
		if value == "" {
			return nil
		}
		for _, a := range allowed {
			if value == a {
				return nil
			}
		}
		return &ValidationError{
			Field:   fieldName,
			Message: fieldName + " must be one of the allowed values",
		}
	}
}

// ValidatorCustom returns a validator that uses a custom function.
// The function receives the value and returns an error message if invalid, or empty string if valid.
func ValidatorCustom(fn func(value string) string) Validator {
	return func(fieldName string, value string) *ValidationError {
		if msg := fn(value); msg != "" {
			return &ValidationError{
				Field:   fieldName,
				Message: msg,
			}
		}
		return nil
	}
}

// Validate validates the given values against the form fields and their validators.
// It returns a slice of ValidationError. An empty slice means validation passed.
// Errors are also stored on the form for inline display when Build() is called.
func (form *Form) Validate(values map[string]string) []ValidationError {
	var errors []ValidationError

	for _, field := range form.fields {
		f, ok := field.(*Field)
		if !ok {
			continue
		}

		value := values[f.Name]

		if f.Required && strings.TrimSpace(value) == "" {
			errors = append(errors, ValidationError{
				Field:   f.Name,
				Message: f.Name + " is required",
			})
			continue
		}

		for _, validator := range f.Validators {
			if err := validator(f.Name, value); err != nil {
				errors = append(errors, *err)
			}
		}
	}

	// Store errors for inline display
	errorMap := make(map[string]string)
	for _, e := range errors {
		if _, exists := errorMap[e.Field]; !exists {
			errorMap[e.Field] = e.Message
		}
	}
	form.errors = errorMap

	return errors
}
