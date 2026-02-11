package form

// NewStringField creates a new text input field with the given name and label.
func NewStringField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_STRING, Name: name, Label: label}
}

// NewEmailField creates a new email input field with the given name and label.
func NewEmailField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_EMAIL, Name: name, Label: label}
}

// NewNumberField creates a new number input field with the given name and label.
func NewNumberField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_NUMBER, Name: name, Label: label}
}

// NewPasswordField creates a new password input field with the given name and label.
func NewPasswordField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_PASSWORD, Name: name, Label: label}
}

// NewHiddenField creates a new hidden input field with the given name and value.
func NewHiddenField(name, value string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_HIDDEN, Name: name, Value: value}
}

// NewDateField creates a new date input field with the given name and label.
func NewDateField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_DATE, Name: name, Label: label}
}

// NewDateTimeField creates a new datetime input field with the given name and label.
func NewDateTimeField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_DATETIME, Name: name, Label: label}
}

// NewSelectField creates a new select field with the given name, label, and options.
func NewSelectField(name, label string, options []FieldOption) *Field {
	return &Field{Type: FORM_FIELD_TYPE_SELECT, Name: name, Label: label, Options: options}
}

// NewTextAreaField creates a new textarea field with the given name and label.
func NewTextAreaField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_TEXTAREA, Name: name, Label: label}
}

// NewCheckboxField creates a new checkbox field with the given name and label.
func NewCheckboxField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_CHECKBOX, Name: name, Label: label}
}

// NewRadioField creates a new radio button group with the given name, label, and options.
func NewRadioField(name, label string, options []FieldOption) *Field {
	return &Field{Type: FORM_FIELD_TYPE_RADIO, Name: name, Label: label, Options: options}
}

// NewFileField creates a new file upload field with the given name and label.
func NewFileField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_FILE, Name: name, Label: label}
}

// NewImageField creates a new image field with the given name and label.
func NewImageField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_IMAGE, Name: name, Label: label}
}

// NewColorField creates a new color picker field with the given name and label.
func NewColorField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_COLOR, Name: name, Label: label}
}

// NewTelField creates a new telephone input field with the given name and label.
func NewTelField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_TEL, Name: name, Label: label}
}

// NewURLField creates a new URL input field with the given name and label.
func NewURLField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_URL, Name: name, Label: label}
}

// NewHtmlAreaField creates a new HTML area (WYSIWYG) field with the given name and label.
func NewHtmlAreaField(name, label string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_HTMLAREA, Name: name, Label: label}
}

// NewRawField creates a new raw HTML field with the given value (rendered as-is).
func NewRawField(value string) *Field {
	return &Field{Type: FORM_FIELD_TYPE_RAW, Value: value}
}
