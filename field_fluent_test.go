package form

import (
	"strings"
	"testing"

	"github.com/dracory/hb"
)

func TestFluentFieldWithID(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithID("myID")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `id="myID"`) {
		t.Fatal("Expected id attribute, got:", html)
	}
}

func TestFluentFieldWithName(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING}).WithName("username")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `name="username"`) {
		t.Fatal("Expected name attribute, got:", html)
	}
}

func TestFluentFieldWithLabel(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithLabel("My Label")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `My Label`) {
		t.Fatal("Expected label text, got:", html)
	}
}

func TestFluentFieldWithValue(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithValue("hello")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `value="hello"`) {
		t.Fatal("Expected value attribute, got:", html)
	}
}

func TestFluentFieldWithHelp(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithHelp("Help text here")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `Help text here`) {
		t.Fatal("Expected help text, got:", html)
	}
}

func TestFluentFieldWithPlaceholder(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithPlaceholder("Enter value")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `placeholder="Enter value"`) {
		t.Fatal("Expected placeholder, got:", html)
	}
}

func TestFluentFieldWithRequired(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithRequired()
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `<sup class="text-danger ms-1">*</sup>`) {
		t.Fatal("Expected required marker, got:", html)
	}
}

func TestFluentFieldWithReadonly(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithReadonly()
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `readonly="readonly"`) {
		t.Fatal("Expected readonly attribute, got:", html)
	}
}

func TestFluentFieldWithDisabled(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithDisabled()
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatal("Expected disabled attribute, got:", html)
	}
}

func TestFluentFieldWithInvisible(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).WithInvisible()
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `display:none;`) {
		t.Fatal("Expected invisible style, got:", html)
	}
}

func TestFluentFieldWithMultiple(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_SELECT, Name: "test"}).WithMultiple().
		WithOptions(FieldOption{Key: "a", Value: "A"})
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `multiple="multiple"`) {
		t.Fatal("Expected multiple attribute, got:", html)
	}
}

func TestFluentFieldWithOptions(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_SELECT, Name: "test"}).
		WithOptions(
			FieldOption{Key: "a", Value: "Option A"},
			FieldOption{Key: "b", Value: "Option B"},
		)
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `Option A`) || !strings.Contains(html, `Option B`) {
		t.Fatal("Expected options, got:", html)
	}
}

func TestFluentFieldWithOptionsF(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_SELECT, Name: "test"}).
		WithOptionsF(func() []FieldOption {
			return []FieldOption{{Key: "dyn", Value: "Dynamic"}}
		})
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `Dynamic`) {
		t.Fatal("Expected dynamic option, got:", html)
	}
}

func TestFluentFieldWithAttr(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).
		WithAttr("data-x", "1").
		WithAttr("data-y", "2")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `data-x="1"`) || !strings.Contains(html, `data-y="2"`) {
		t.Fatal("Expected custom attrs, got:", html)
	}
}

func TestFluentFieldWithAttrs(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).
		WithAttrs(map[string]string{"aria-label": "Test"})
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `aria-label="Test"`) {
		t.Fatal("Expected attrs, got:", html)
	}
}

func TestFluentFieldWithCustomInput(t *testing.T) {
	custom := hb.NewDiv().Class("custom-input").HTML("custom content")
	f := (&Field{Type: FORM_FIELD_TYPE_BLOCKEDITOR, Name: "test"}).WithCustomInput(custom)
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `custom content`) {
		t.Fatal("Expected custom input, got:", html)
	}
}

func TestFluentFieldWithValidators(t *testing.T) {
	f := (&Field{Type: FORM_FIELD_TYPE_STRING, Name: "test"}).
		WithValidators(ValidatorMinLength(5))

	if len(f.Validators) != 1 {
		t.Fatal("Expected 1 validator, got:", len(f.Validators))
	}
}

func TestFluentFieldFullChain(t *testing.T) {
	f := (&Field{}).
		WithType(FORM_FIELD_TYPE_EMAIL).
		WithID("emailField").
		WithName("email").
		WithLabel("Email Address").
		WithValue("user@example.com").
		WithPlaceholder("you@example.com").
		WithHelp("We will never share your email.").
		WithRequired().
		WithAttr("autocomplete", "email")

	html := f.BuildFormGroup("").ToHTML()

	expecteds := []string{
		`id="emailField"`,
		`name="email"`,
		`type="email"`,
		`value="user@example.com"`,
		`placeholder="you@example.com"`,
		`Email Address`,
		`We will never share your email.`,
		`<sup class="text-danger ms-1">*</sup>`,
		`autocomplete="email"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal("Expected:", expected, "got:", html)
		}
	}
}
