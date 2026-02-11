package form

import (
	"strings"
	"testing"
)

func TestNewFieldRowEqualColumns(t *testing.T) {
	f := New().WithFields(
		NewFieldRow(
			NewStringField("first", "First Name"),
			NewStringField("last", "Last Name"),
		),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, `class="row"`) {
		t.Fatal("Expected row class, got:", html)
	}

	// Two equal "col" divs
	if strings.Count(html, `class="col"`) != 2 {
		t.Fatal("Expected 2 col divs, got:", html)
	}

	if !strings.Contains(html, `name="first"`) || !strings.Contains(html, `name="last"`) {
		t.Fatal("Expected both fields, got:", html)
	}
}

func TestNewFieldRowWithColumns(t *testing.T) {
	f := New().WithFields(
		NewFieldRowWithColumns(
			FieldRowColumn{Field: NewStringField("first", "First"), ColClass: "col-md-4"},
			FieldRowColumn{Field: NewStringField("last", "Last"), ColClass: "col-md-8"},
		),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, `class="col-md-4"`) {
		t.Fatal("Expected col-md-4, got:", html)
	}
	if !strings.Contains(html, `class="col-md-8"`) {
		t.Fatal("Expected col-md-8, got:", html)
	}
}

func TestFieldRowWithCustomRowClass(t *testing.T) {
	f := New().WithFields(
		NewFieldRow(
			NewStringField("a", "A"),
		).WithRowClass("row g-3"),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, `class="row g-3"`) {
		t.Fatal("Expected custom row class, got:", html)
	}
}

func TestFieldRowWithTheme(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewFieldRow(
			NewStringField("name", "Name"),
		),
	)
	html := f.Build().ToHTML()

	// Child field should use Tailwind classes
	if !strings.Contains(html, tw.FormGroupClass) {
		t.Fatal("Expected Tailwind form group class in row child, got:", html)
	}
	if !strings.Contains(html, tw.LabelClass) {
		t.Fatal("Expected Tailwind label class in row child, got:", html)
	}
}

func TestFieldRowWithErrors(t *testing.T) {
	f := New().WithFields(
		NewFieldRow(
			NewStringField("first", "First Name"),
			NewStringField("last", "Last Name"),
		),
	).WithErrors(map[string]string{
		"first": "First name is required",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `First name is required`) {
		t.Fatal("Expected error on first field in row, got:", html)
	}
	if !strings.Contains(html, `is-invalid`) {
		t.Fatal("Expected error input class, got:", html)
	}
}

func TestFieldRowMixedWithRegularFields(t *testing.T) {
	f := New().WithFields(
		NewStringField("title", "Title"),
		NewFieldRow(
			NewStringField("first", "First"),
			NewStringField("last", "Last"),
		),
		NewEmailField("email", "Email"),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, `name="title"`) {
		t.Fatal("Expected title field, got:", html)
	}
	if !strings.Contains(html, `class="row"`) {
		t.Fatal("Expected row, got:", html)
	}
	if !strings.Contains(html, `name="email"`) {
		t.Fatal("Expected email field, got:", html)
	}
}

func TestFieldRowThreeColumns(t *testing.T) {
	f := New().WithFields(
		NewFieldRow(
			NewStringField("a", "A"),
			NewStringField("b", "B"),
			NewStringField("c", "C"),
		),
	)
	html := f.Build().ToHTML()

	if strings.Count(html, `class="col"`) != 3 {
		t.Fatal("Expected 3 col divs, got:", html)
	}
}
