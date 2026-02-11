package form

import (
	"strings"
	"testing"
)

// AssertFormContains checks that the rendered form HTML contains the expected string.
// It calls t.Helper() so the test failure points to the caller.
func AssertFormContains(t *testing.T, form *Form, expected string) {
	t.Helper()
	html := form.Build().ToHTML()
	if !strings.Contains(html, expected) {
		t.Fatalf("Expected form HTML to contain %q, got:\n%s", expected, html)
	}
}

// AssertFormNotContains checks that the rendered form HTML does NOT contain the given string.
func AssertFormNotContains(t *testing.T, form *Form, unexpected string) {
	t.Helper()
	html := form.Build().ToHTML()
	if strings.Contains(html, unexpected) {
		t.Fatalf("Expected form HTML to NOT contain %q, got:\n%s", unexpected, html)
	}
}

// AssertFieldContains checks that the rendered field HTML contains the expected string.
func AssertFieldContains(t *testing.T, field FieldInterface, expected string) {
	t.Helper()
	html := field.BuildFormGroup("").ToHTML()
	if !strings.Contains(html, expected) {
		t.Fatalf("Expected field HTML to contain %q, got:\n%s", expected, html)
	}
}

// AssertFieldNotContains checks that the rendered field HTML does NOT contain the given string.
func AssertFieldNotContains(t *testing.T, field FieldInterface, unexpected string) {
	t.Helper()
	html := field.BuildFormGroup("").ToHTML()
	if strings.Contains(html, unexpected) {
		t.Fatalf("Expected field HTML to NOT contain %q, got:\n%s", unexpected, html)
	}
}

// AssertValidationPasses checks that the form validation passes for the given values.
func AssertValidationPasses(t *testing.T, form *Form, values map[string]string) {
	t.Helper()
	errs := form.Validate(values)
	if len(errs) != 0 {
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = e.Field + ": " + e.Message
		}
		t.Fatalf("Expected validation to pass, got %d error(s):\n%s", len(errs), strings.Join(msgs, "\n"))
	}
}

// AssertValidationFails checks that the form validation fails for the given values.
// It returns the validation errors for further inspection.
func AssertValidationFails(t *testing.T, form *Form, values map[string]string) []ValidationError {
	t.Helper()
	errs := form.Validate(values)
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail, but it passed")
	}
	return errs
}

// AssertValidationFailsOn checks that validation fails specifically on the named field.
func AssertValidationFailsOn(t *testing.T, form *Form, values map[string]string, fieldName string) {
	t.Helper()
	errs := form.Validate(values)
	for _, e := range errs {
		if e.Field == fieldName {
			return
		}
	}
	t.Fatalf("Expected validation to fail on field %q, but it did not", fieldName)
}

// AssertValidationErrorCount checks that validation produces exactly n errors.
func AssertValidationErrorCount(t *testing.T, form *Form, values map[string]string, expectedCount int) {
	t.Helper()
	errs := form.Validate(values)
	if len(errs) != expectedCount {
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = e.Field + ": " + e.Message
		}
		t.Fatalf("Expected %d validation error(s), got %d:\n%s", expectedCount, len(errs), strings.Join(msgs, "\n"))
	}
}
