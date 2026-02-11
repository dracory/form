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

	theme  *Theme
	errors map[string]string // field name -> error message

	htmxConfig *HTMXConfig
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

// SetErrors sets the validation error messages to display inline next to fields.
// The map keys are field names, values are error messages.
func (form *Form) SetErrors(errors map[string]string) {
	form.errors = errors
}

// GetErrors returns the current validation error messages.
func (form *Form) GetErrors() map[string]string {
	return form.errors
}

// formAware is an optional interface for fields that need a reference to their parent form.
type formAware interface {
	setForm(form *Form)
}

// themeable is an optional interface for fields that support theming.
type themeable interface {
	setTheme(theme *Theme)
}

// errorAware is an optional interface for fields that support inline error display.
type errorAware interface {
	setError(message string)
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
		if ea, ok := field.(errorAware); ok && form.errors != nil {
			if msg, exists := form.errors[field.GetName()]; exists {
				ea.setError(msg)
			}
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

	// Apply HTMXConfig if set (overrides individual hx* fields where non-empty)
	if c := form.htmxConfig; c != nil {
		if c.Post != "" {
			hbForm.HxPost(c.Post)
		}
		if c.Get != "" {
			hbForm.Attr("hx-get", c.Get)
		}
		if c.Target != "" {
			hbForm.HxTarget(c.Target)
		}
		if c.Swap != "" {
			hbForm.HxSwap(hb.SwapMethod(c.Swap))
		}
		if c.Trigger != "" {
			hbForm.Attr("hx-trigger", c.Trigger)
		}
		if c.Indicator != "" {
			hbForm.Attr("hx-indicator", c.Indicator)
		}
		if c.Confirm != "" {
			hbForm.Attr("hx-confirm", c.Confirm)
		}
		if c.Sync != "" {
			hbForm.Attr("hx-sync", c.Sync)
		}
		if c.Validate {
			hbForm.Attr("hx-validate", "true")
		}
		if c.DisabledElt != "" {
			hbForm.Attr("hx-disabled-elt", c.DisabledElt)
		}
		if c.Encoding != "" {
			hbForm.Attr("hx-encoding", c.Encoding)
		}
		if c.PushURL != "" {
			hbForm.Attr("hx-push-url", c.PushURL)
		}
	}

	return hbForm
}
