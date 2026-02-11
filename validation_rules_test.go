package form

import (
	"testing"
)

func TestValidatorURL_Valid(t *testing.T) {
	v := ValidatorURL()
	if err := v("url", "https://example.com"); err != nil {
		t.Fatal("Expected valid URL, got error:", err.Message)
	}
	if err := v("url", "http://example.com/path?q=1"); err != nil {
		t.Fatal("Expected valid URL, got error:", err.Message)
	}
}

func TestValidatorURL_Invalid(t *testing.T) {
	v := ValidatorURL()
	if err := v("url", "not-a-url"); err == nil {
		t.Fatal("Expected error for invalid URL")
	}
	if err := v("url", "ftp://example.com"); err == nil {
		t.Fatal("Expected error for ftp URL")
	}
}

func TestValidatorURL_Empty(t *testing.T) {
	v := ValidatorURL()
	if err := v("url", ""); err != nil {
		t.Fatal("Expected no error for empty value, got:", err.Message)
	}
}

func TestValidatorIP_Valid(t *testing.T) {
	v := ValidatorIP()
	if err := v("ip", "192.168.1.1"); err != nil {
		t.Fatal("Expected valid IP, got error:", err.Message)
	}
	if err := v("ip", "0.0.0.0"); err != nil {
		t.Fatal("Expected valid IP, got error:", err.Message)
	}
	if err := v("ip", "255.255.255.255"); err != nil {
		t.Fatal("Expected valid IP, got error:", err.Message)
	}
}

func TestValidatorIP_Invalid(t *testing.T) {
	v := ValidatorIP()
	if err := v("ip", "999.999.999.999"); err == nil {
		t.Fatal("Expected error for invalid IP")
	}
	if err := v("ip", "abc"); err == nil {
		t.Fatal("Expected error for non-IP string")
	}
}

func TestValidatorUUID_Valid(t *testing.T) {
	v := ValidatorUUID()
	if err := v("id", "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatal("Expected valid UUID, got error:", err.Message)
	}
	if err := v("id", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"); err != nil {
		t.Fatal("Expected valid UUID, got error:", err.Message)
	}
}

func TestValidatorUUID_Invalid(t *testing.T) {
	v := ValidatorUUID()
	if err := v("id", "not-a-uuid"); err == nil {
		t.Fatal("Expected error for invalid UUID")
	}
	if err := v("id", "550e8400e29b41d4a716446655440000"); err == nil {
		t.Fatal("Expected error for UUID without dashes")
	}
}

func TestValidatorAlphaNumeric_Valid(t *testing.T) {
	v := ValidatorAlphaNumeric()
	if err := v("code", "abc123"); err != nil {
		t.Fatal("Expected valid alphanumeric, got error:", err.Message)
	}
	if err := v("code", "ABC"); err != nil {
		t.Fatal("Expected valid alphanumeric, got error:", err.Message)
	}
}

func TestValidatorAlphaNumeric_Invalid(t *testing.T) {
	v := ValidatorAlphaNumeric()
	if err := v("code", "abc-123"); err == nil {
		t.Fatal("Expected error for string with dash")
	}
	if err := v("code", "hello world"); err == nil {
		t.Fatal("Expected error for string with space")
	}
	if err := v("code", "test@email"); err == nil {
		t.Fatal("Expected error for string with @")
	}
}

func TestValidatorOneOf_Valid(t *testing.T) {
	v := ValidatorOneOf("red", "green", "blue")
	if err := v("color", "red"); err != nil {
		t.Fatal("Expected valid value, got error:", err.Message)
	}
	if err := v("color", "blue"); err != nil {
		t.Fatal("Expected valid value, got error:", err.Message)
	}
}

func TestValidatorOneOf_Invalid(t *testing.T) {
	v := ValidatorOneOf("red", "green", "blue")
	if err := v("color", "yellow"); err == nil {
		t.Fatal("Expected error for value not in list")
	}
}

func TestValidatorOneOf_Empty(t *testing.T) {
	v := ValidatorOneOf("red", "green", "blue")
	if err := v("color", ""); err != nil {
		t.Fatal("Expected no error for empty value, got:", err.Message)
	}
}

func TestValidatorCustom_Valid(t *testing.T) {
	v := ValidatorCustom(func(value string) string {
		if value == "secret" {
			return ""
		}
		return "must be 'secret'"
	})
	if err := v("code", "secret"); err != nil {
		t.Fatal("Expected valid, got error:", err.Message)
	}
}

func TestValidatorCustom_Invalid(t *testing.T) {
	v := ValidatorCustom(func(value string) string {
		if value == "secret" {
			return ""
		}
		return "must be 'secret'"
	})
	err := v("code", "wrong")
	if err == nil {
		t.Fatal("Expected error for invalid value")
	}
	if err.Message != "must be 'secret'" {
		t.Fatal("Expected custom message, got:", err.Message)
	}
}

func TestValidatorCustom_WithFieldName(t *testing.T) {
	v := ValidatorCustom(func(value string) string {
		if len(value) > 0 && value[0] == '#' {
			return ""
		}
		return "must start with #"
	})
	err := v("color", "red")
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Field != "color" {
		t.Fatal("Expected field name 'color', got:", err.Field)
	}
}

func TestValidatorsInForm(t *testing.T) {
	f := New().WithFields(
		NewURLField("website", "Website").WithValidators(ValidatorURL()),
		NewStringField("code", "Code").WithValidators(ValidatorAlphaNumeric()),
		NewSelectField("color", "Color", []FieldOption{
			{Key: "red", Value: "Red"},
			{Key: "blue", Value: "Blue"},
		}).WithValidators(ValidatorOneOf("red", "blue")),
	)

	errs := f.Validate(map[string]string{
		"website": "not-a-url",
		"code":    "has spaces",
		"color":   "green",
	})

	if len(errs) != 3 {
		t.Fatal("Expected 3 errors, got:", len(errs))
	}
}
