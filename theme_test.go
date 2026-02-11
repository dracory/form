package form

import (
	"strings"
	"testing"
)

func TestThemeBootstrap5Default(t *testing.T) {
	f := New().WithFields(
		NewStringField("name", "Name"),
	)
	html := f.Build().ToHTML()

	expecteds := []string{
		`class="form-group mb-3"`,
		`class="form-label"`,
		`class="form-control"`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected Bootstrap 5 class:", exp, "got:", html)
		}
	}
}

func TestThemeTailwind(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewStringField("name", "Name"),
	)
	html := f.Build().ToHTML()

	expecteds := []string{
		`class="mb-4"`,
		`class="block text-sm font-medium text-gray-700 mb-1"`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected Tailwind class:", exp, "got:", html)
		}
	}

	// Should NOT contain Bootstrap classes
	if strings.Contains(html, `form-group`) {
		t.Fatal("Should not contain Bootstrap classes with Tailwind theme, got:", html)
	}
}

func TestThemeTailwindSelect(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewSelectField("color", "Color", []FieldOption{
			{Key: "r", Value: "Red"},
		}),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.SelectClass) {
		t.Fatal("Expected Tailwind select class, got:", html)
	}
	if strings.Contains(html, `form-select`) {
		t.Fatal("Should not contain Bootstrap select class, got:", html)
	}
}

func TestThemeTailwindCheckbox(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewCheckboxField("agree", "I Agree").WithValue("1"),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.CheckboxWrapClass) {
		t.Fatal("Expected Tailwind checkbox wrap class, got:", html)
	}
	if !strings.Contains(html, tw.CheckboxInputClass) {
		t.Fatal("Expected Tailwind checkbox input class, got:", html)
	}
}

func TestThemeTailwindRadio(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewRadioField("opt", "Option", []FieldOption{
			{Key: "a", Value: "A"},
		}),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.RadioWrapClass) {
		t.Fatal("Expected Tailwind radio wrap class, got:", html)
	}
	if !strings.Contains(html, tw.RadioLabelClass) {
		t.Fatal("Expected Tailwind radio label class, got:", html)
	}
}

func TestThemeTailwindTextArea(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewTextAreaField("bio", "Bio"),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.TextAreaClass) {
		t.Fatal("Expected Tailwind textarea class, got:", html)
	}
}

func TestThemeTailwindFile(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewFileField("doc", "Document"),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.FileInputClass) {
		t.Fatal("Expected Tailwind file class, got:", html)
	}
}

func TestThemeRequired(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewStringField("name", "Name").WithRequired(),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.RequiredClass) {
		t.Fatal("Expected Tailwind required class, got:", html)
	}
}

func TestThemeHelp(t *testing.T) {
	tw := ThemeTailwind()
	f := New().WithTheme(tw).WithFields(
		NewStringField("name", "Name").WithHelp("Enter your name"),
	)
	html := f.Build().ToHTML()

	if !strings.Contains(html, tw.HelpClass) {
		t.Fatal("Expected Tailwind help class, got:", html)
	}
}

func TestCustomTheme(t *testing.T) {
	custom := &Theme{
		FormGroupClass: "my-group",
		LabelClass:     "my-label",
		InputClass:     "my-input",
		SelectClass:    "my-select",
		TextAreaClass:  "my-textarea",
		HelpClass:      "my-help",
		RequiredClass:  "my-required",
		RequiredMarker: "!",
		TableClass:     "my-table",
	}

	f := New().WithTheme(custom).WithFields(
		NewStringField("name", "Name").WithRequired().WithHelp("help"),
	)
	html := f.Build().ToHTML()

	expecteds := []string{
		`class="my-group"`,
		`class="my-label"`,
		`class="my-input"`,
		`class="my-required"`,
		`class="my-help"`,
		`>!</sup>`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected custom class:", exp, "got:", html)
		}
	}
}

func TestFieldWithoutFormUsesDefaultTheme(t *testing.T) {
	f := NewStringField("name", "Name")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `class="form-group mb-3"`) {
		t.Fatal("Expected default Bootstrap 5 theme, got:", html)
	}
}
