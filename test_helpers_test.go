package form

import (
	"testing"
)

func TestAssertFormContains(t *testing.T) {
	f := New().WithID("myForm").WithFields(
		NewStringField("name", "Name"),
	)
	AssertFormContains(t, f, `id="myForm"`)
	AssertFormContains(t, f, `name="name"`)
}

func TestAssertFormNotContains(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name"),
	)
	AssertFormNotContains(t, f, `hx-post`)
	AssertFormNotContains(t, f, `is-invalid`)
}

func TestAssertFieldContains(t *testing.T) {
	field := NewEmailField("email", "Email").WithPlaceholder("you@example.com")
	AssertFieldContains(t, field, `type="email"`)
	AssertFieldContains(t, field, `placeholder="you@example.com"`)
}

func TestAssertFieldNotContains(t *testing.T) {
	field := NewStringField("name", "Name")
	AssertFieldNotContains(t, field, `is-invalid`)
	AssertFieldNotContains(t, field, `readonly`)
}

func TestAssertValidationPasses(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
		NewEmailField("email", "Email").WithValidators(ValidatorEmail()),
	)
	AssertValidationPasses(t, f, map[string]string{
		"name":  "John",
		"email": "john@example.com",
	})
}

func TestAssertValidationFails(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
	)
	errs := AssertValidationFails(t, f, map[string]string{
		"name": "",
	})
	if errs[0].Field != "name" {
		t.Fatal("Expected error on name field")
	}
}

func TestAssertValidationFailsOn(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
		NewEmailField("email", "Email").WithRequired(),
	)
	AssertValidationFailsOn(t, f, map[string]string{
		"name":  "John",
		"email": "",
	}, "email")
}

func TestAssertValidationErrorCount(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
		NewEmailField("email", "Email").WithRequired(),
		NewStringField("phone", "Phone").WithRequired(),
	)
	AssertValidationErrorCount(t, f, map[string]string{
		"name":  "",
		"email": "",
		"phone": "123",
	}, 2)
}

func TestHelpersEndToEnd(t *testing.T) {
	f := New().
		WithID("contactForm").
		WithHTMX(HTMXConfig{Post: "/contact", Target: "#result"}).
		WithFields(
			NewFieldRow(
				NewStringField("first", "First").WithRequired(),
				NewStringField("last", "Last").WithRequired(),
			),
			NewEmailField("email", "Email").WithRequired().WithValidators(ValidatorEmail()),
		)

	// Test form structure
	AssertFormContains(t, f, `id="contactForm"`)
	AssertFormContains(t, f, `hx-post="/contact"`)
	AssertFormContains(t, f, `class="row"`)

	// Test validation passes
	AssertValidationPasses(t, f, map[string]string{
		"first": "John",
		"last":  "Doe",
		"email": "john@example.com",
	})

	// Test validation fails (only top-level fields are validated; row children are nested)
	AssertValidationErrorCount(t, f, map[string]string{
		"email": "invalid",
	}, 1)

	// Test inline errors render
	f.Validate(map[string]string{
		"email": "bad",
	})
	AssertFormContains(t, f, `is-invalid`)
	AssertFormContains(t, f, `invalid-feedback`)
}
