package form

import (
	"strings"
	"testing"
)

func TestNewFluentForm(t *testing.T) {
	f := New()
	html := f.Build().ToHTML()

	expected := `<form method="POST"></form>`
	if html != expected {
		t.Fatal("Expected:", expected, "but was:", html)
	}
}

func TestFluentFormWithID(t *testing.T) {
	f := New().WithID("myForm")
	html := f.Build().ToHTML()

	if !strings.Contains(html, `id="myForm"`) {
		t.Fatal("Expected id attribute, got:", html)
	}
}

func TestFluentFormWithClass(t *testing.T) {
	f := New().WithClass("custom-form")
	html := f.Build().ToHTML()

	if !strings.Contains(html, `class="custom-form"`) {
		t.Fatal("Expected class attribute, got:", html)
	}
}

func TestFluentFormWithMethod(t *testing.T) {
	f := New().WithMethod("GET")
	html := f.Build().ToHTML()

	if !strings.Contains(html, `method="GET"`) {
		t.Fatal("Expected GET method, got:", html)
	}
}

func TestFluentFormWithAction(t *testing.T) {
	f := New().WithAction("/submit")
	html := f.Build().ToHTML()

	if !strings.Contains(html, `action="/submit"`) {
		t.Fatal("Expected action attribute, got:", html)
	}
}

func TestFluentFormWithFields(t *testing.T) {
	f := New().WithFields(
		&Field{Name: "name", Type: FORM_FIELD_TYPE_STRING},
		&Field{Name: "email", Type: FORM_FIELD_TYPE_EMAIL},
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, `name="name"`) {
		t.Fatal("Expected name field, got:", html)
	}
	if !strings.Contains(html, `name="email"`) {
		t.Fatal("Expected email field, got:", html)
	}
}

func TestFluentFormWithFileManager(t *testing.T) {
	f := New().WithFileManager("/files")

	if f.GetFileManagerURL() != "/files" {
		t.Fatal("Expected /files, got:", f.GetFileManagerURL())
	}
}

func TestFluentFormWithHtmx(t *testing.T) {
	f := New().
		WithHxPost("/api/submit").
		WithHxTarget("#result").
		WithHxSwap("innerHTML")

	html := f.Build().ToHTML()

	expecteds := []string{
		`hx-post="/api/submit"`,
		`hx-target="#result"`,
		`hx-swap="innerHTML"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal("Expected:", expected, "got:", html)
		}
	}
}

func TestFluentFormFullChain(t *testing.T) {
	f := New().
		WithID("signupForm").
		WithClass("needs-validation").
		WithAction("/signup").
		WithMethod("POST").
		WithFileManager("/files").
		WithHxPost("/api/signup").
		WithHxTarget("#response").
		WithHxSwap("outerHTML").
		WithFields(
			&Field{Name: "username", Type: FORM_FIELD_TYPE_STRING},
		)

	html := f.Build().ToHTML()

	expecteds := []string{
		`id="signupForm"`,
		`class="needs-validation"`,
		`action="/signup"`,
		`method="POST"`,
		`hx-post="/api/signup"`,
		`hx-target="#response"`,
		`hx-swap="outerHTML"`,
		`name="username"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal("Expected:", expected, "got:", html)
		}
	}

	if f.GetFileManagerURL() != "/files" {
		t.Fatal("Expected /files, got:", f.GetFileManagerURL())
	}
}
