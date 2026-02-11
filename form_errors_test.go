package form

import (
	"strings"
	"testing"
)

func TestInlineErrorDisplay(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
		NewEmailField("email", "Email"),
	).WithErrors(map[string]string{
		"name": "Name is required",
	})

	html := f.Build().ToHTML()

	expecteds := []string{
		`class="invalid-feedback"`,
		`Name is required`,
		`is-invalid`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}

func TestInlineErrorNotShownWhenNoError(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name"),
	)

	html := f.Build().ToHTML()

	if strings.Contains(html, `invalid-feedback`) {
		t.Fatal("Should not contain error class when no errors, got:", html)
	}
	if strings.Contains(html, `is-invalid`) {
		t.Fatal("Should not contain error input class when no errors, got:", html)
	}
}

func TestInlineErrorOnlyOnMatchingField(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name"),
		NewEmailField("email", "Email"),
	).WithErrors(map[string]string{
		"email": "Invalid email",
	})

	html := f.Build().ToHTML()

	if !strings.Contains(html, `Invalid email`) {
		t.Fatal("Expected error on email field, got:", html)
	}
}

func TestInlineErrorWithTailwindTheme(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewStringField("name", "Name"),
	).WithErrors(map[string]string{
		"name": "Name is required",
	})

	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.ErrorClass) {
		t.Fatal("Expected Tailwind error class, got:", html)
	}
	if !strings.Contains(html, tw.ErrorInputClass) {
		t.Fatal("Expected Tailwind error input class, got:", html)
	}
	if strings.Contains(html, `invalid-feedback`) {
		t.Fatal("Should not contain Bootstrap error class with Tailwind theme, got:", html)
	}
}

func TestValidateThenBuildShowsErrors(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
		NewEmailField("email", "Email").WithRequired(),
	)

	errs := f.Validate(map[string]string{
		"name":  "",
		"email": "",
	})

	if len(errs) != 2 {
		t.Fatal("Expected 2 errors, got:", len(errs))
	}

	html := f.Build().ToHTML()

	expecteds := []string{
		`name is required`,
		`email is required`,
		`is-invalid`,
		`invalid-feedback`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}

func TestValidatePassThenBuildNoErrors(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name").WithRequired(),
	)

	errs := f.Validate(map[string]string{
		"name": "John",
	})

	if len(errs) != 0 {
		t.Fatal("Expected 0 errors, got:", len(errs))
	}

	html := f.Build().ToHTML()

	if strings.Contains(html, `invalid-feedback`) {
		t.Fatal("Should not contain error class when validation passes, got:", html)
	}
}

func TestSetErrorsMethod(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name"),
	)

	f.SetErrors(map[string]string{
		"name": "Custom error",
	})

	html := f.Build().ToHTML()

	if !strings.Contains(html, `Custom error`) {
		t.Fatal("Expected custom error, got:", html)
	}

	// Verify GetErrors
	errs := f.GetErrors()
	if errs["name"] != "Custom error" {
		t.Fatal("Expected GetErrors to return set errors")
	}
}
