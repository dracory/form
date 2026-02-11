package form

import (
	"strings"
	"testing"
)

func TestNewStringField(t *testing.T) {
	f := NewStringField("username", "Username")
	html := f.BuildFormGroup("").ToHTML()

	expecteds := []string{`type="text"`, `name="username"`, `Username`}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}

func TestNewEmailField(t *testing.T) {
	f := NewEmailField("email", "Email")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="email"`) {
		t.Fatal("Expected email type, got:", html)
	}
}

func TestNewNumberField(t *testing.T) {
	f := NewNumberField("age", "Age")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="number"`) {
		t.Fatal("Expected number type, got:", html)
	}
}

func TestNewPasswordField(t *testing.T) {
	f := NewPasswordField("pass", "Password")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="password"`) {
		t.Fatal("Expected password type, got:", html)
	}
}

func TestNewHiddenField(t *testing.T) {
	f := NewHiddenField("token", "abc123")
	html := f.BuildFormGroup("").ToHTML()

	expecteds := []string{`type="hidden"`, `name="token"`, `value="abc123"`}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}

func TestNewDateField(t *testing.T) {
	f := NewDateField("dob", "Date of Birth")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="date"`) {
		t.Fatal("Expected date type, got:", html)
	}
}

func TestNewDateTimeField(t *testing.T) {
	f := NewDateTimeField("event", "Event Time")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `name="event"`) {
		t.Fatal("Expected name, got:", html)
	}
}

func TestNewSelectField(t *testing.T) {
	opts := []FieldOption{{Key: "a", Value: "A"}, {Key: "b", Value: "B"}}
	f := NewSelectField("choice", "Choice", opts)
	html := f.BuildFormGroup("").ToHTML()

	expecteds := []string{`<select`, `name="choice"`, `Option A`, `Option B`}
	// select renders options with their Value text
	if !strings.Contains(html, `<option value="a">A</option>`) {
		t.Fatal("Expected option A, got:", html)
	}
	if !strings.Contains(html, `<option value="b">B</option>`) {
		t.Fatal("Expected option B, got:", html)
	}
	_ = expecteds
}

func TestNewTextAreaField(t *testing.T) {
	f := NewTextAreaField("bio", "Biography")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `<textarea`) {
		t.Fatal("Expected textarea, got:", html)
	}
}

func TestNewCheckboxField(t *testing.T) {
	f := NewCheckboxField("agree", "I Agree")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="checkbox"`) {
		t.Fatal("Expected checkbox type, got:", html)
	}
}

func TestNewRadioField(t *testing.T) {
	opts := []FieldOption{{Key: "m", Value: "Male"}, {Key: "f", Value: "Female"}}
	f := NewRadioField("gender", "Gender", opts)
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="radio"`) {
		t.Fatal("Expected radio type, got:", html)
	}
	if !strings.Contains(html, `Male`) || !strings.Contains(html, `Female`) {
		t.Fatal("Expected radio options, got:", html)
	}
}

func TestNewFileField(t *testing.T) {
	f := NewFileField("doc", "Document")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="file"`) {
		t.Fatal("Expected file type, got:", html)
	}
}

func TestNewImageField(t *testing.T) {
	f := NewImageField("avatar", "Avatar")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `name="avatar"`) {
		t.Fatal("Expected name, got:", html)
	}
}

func TestNewColorField(t *testing.T) {
	f := NewColorField("theme", "Theme Color")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="color"`) {
		t.Fatal("Expected color type, got:", html)
	}
}

func TestNewTelField(t *testing.T) {
	f := NewTelField("phone", "Phone")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="tel"`) {
		t.Fatal("Expected tel type, got:", html)
	}
}

func TestNewURLField(t *testing.T) {
	f := NewURLField("website", "Website")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `type="url"`) {
		t.Fatal("Expected url type, got:", html)
	}
}

func TestNewHtmlAreaField(t *testing.T) {
	f := NewHtmlAreaField("content", "Content")
	html := f.BuildFormGroup("").ToHTML()

	if !strings.Contains(html, `name="content"`) {
		t.Fatal("Expected name, got:", html)
	}
}

func TestNewRawField(t *testing.T) {
	f := NewRawField(`<hr />`)
	html := f.BuildFormGroup("").ToHTML()

	if html != `<hr />` {
		t.Fatal("Expected raw HTML, got:", html)
	}
}

func TestConstructorWithChaining(t *testing.T) {
	f := NewEmailField("email", "Email Address").
		WithRequired().
		WithPlaceholder("user@example.com").
		WithValidators(ValidatorEmail())

	html := f.BuildFormGroup("").ToHTML()

	expecteds := []string{
		`type="email"`,
		`placeholder="user@example.com"`,
		`<sup class="text-danger ms-1">*</sup>`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}

	if len(f.Validators) != 1 {
		t.Fatal("Expected 1 validator, got:", len(f.Validators))
	}
}

func TestConstructorInFluentForm(t *testing.T) {
	f := New().
		WithID("myForm").
		WithFields(
			NewStringField("name", "Name").WithRequired(),
			NewEmailField("email", "Email").WithPlaceholder("you@example.com"),
			NewHiddenField("csrf", "token123"),
		)

	html := f.Build().ToHTML()

	expecteds := []string{
		`id="myForm"`,
		`name="name"`,
		`name="email"`,
		`value="token123"`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}
