package form

import (
	"github.com/dracory/hb"
)

// Form represents an HTML form with fields, styling, and optional HTMX integration.
type Form struct {
	id             string
	className      string
	fields         []FieldInterface
	fileManagerURL string
	method         string
	actionUrl      string

	// HTMX helpers
	hxPost   string
	hxTarget string
	hxSwap   string

	theme *Theme
}

// AddField appends a field to the form.
func (form *Form) AddField(field FieldInterface) {
	form.fields = append(form.fields, field)
}

// GetFields returns all fields in the form.
func (form *Form) GetFields() []FieldInterface {
	return form.fields
}

// SetFields replaces all fields in the form.
func (form *Form) SetFields(fields []FieldInterface) {
	form.fields = fields
}

// GetFileManagerURL returns the file manager URL used for image fields.
func (form *Form) GetFileManagerURL() string {
	return form.fileManagerURL
}

// SetFileManagerURL sets the file manager URL used for image fields.
func (form *Form) SetFileManagerURL(url string) {
	form.fileManagerURL = url
}

// formAware is an optional interface for fields that need a reference to their parent form.
type formAware interface {
	setForm(form *Form)
}

// themeable is an optional interface for fields that support theming.
type themeable interface {
	setTheme(theme *Theme)
}

// Build renders the form and all its fields into an hb.Tag HTML element.
func (form *Form) Build() *hb.Tag {
	tags := []hb.TagInterface{}

	theme := form.theme
	if theme == nil {
		theme = defaultTheme
	}

	for _, field := range form.fields {
		if fa, ok := field.(formAware); ok {
			fa.setForm(form)
		}
		if th, ok := field.(themeable); ok {
			th.setTheme(theme)
		}
		tags = append(tags, field.BuildFormGroup(form.fileManagerURL))
	}

	hbForm := hb.Form()
	hbForm.Children(tags)
	hbForm.Method(form.method)

	if form.actionUrl != "" {
		hbForm.Action(form.actionUrl)
	}

	if form.id != "" {
		hbForm.ID(form.id)
	}

	if form.className != "" {
		hbForm.Class(form.className)
	}

	if form.hxPost != "" {
		hbForm.HxPost(form.hxPost)
	}

	if form.hxTarget != "" {
		hbForm.HxTarget(form.hxTarget)
	}

	if form.hxSwap != "" {
		hbForm.HxSwap(hb.SwapMethod(form.hxSwap))
	}

	return hbForm
}
