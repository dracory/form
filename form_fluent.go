package form

import "net/http"

// New creates a new empty Form with POST as the default method.
// Use the With* methods to configure it via chaining.
func New() *Form {
	return &Form{
		method: http.MethodPost,
	}
}

// WithID sets the form's HTML id attribute.
func (form *Form) WithID(id string) *Form {
	form.id = id
	return form
}

// WithClass sets the form's CSS class.
func (form *Form) WithClass(className string) *Form {
	form.className = className
	return form
}

// WithMethod sets the form's HTTP method.
func (form *Form) WithMethod(method string) *Form {
	form.method = method
	return form
}

// WithAction sets the form's action URL.
func (form *Form) WithAction(actionURL string) *Form {
	form.actionUrl = actionURL
	return form
}

// WithFields sets the form's fields.
func (form *Form) WithFields(fields ...FieldInterface) *Form {
	form.fields = fields
	return form
}

// WithFileManager sets the file manager URL used for image fields.
func (form *Form) WithFileManager(url string) *Form {
	form.fileManagerURL = url
	return form
}

// WithHxPost sets the hx-post attribute for HTMX integration.
func (form *Form) WithHxPost(url string) *Form {
	form.hxPost = url
	return form
}

// WithHxTarget sets the hx-target attribute for HTMX integration.
func (form *Form) WithHxTarget(target string) *Form {
	form.hxTarget = target
	return form
}

// WithHxSwap sets the hx-swap attribute for HTMX integration.
func (form *Form) WithHxSwap(swap string) *Form {
	form.hxSwap = swap
	return form
}
