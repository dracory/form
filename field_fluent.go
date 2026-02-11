package form

import "github.com/dracory/hb"

// WithID sets the field's HTML id attribute.
func (field *Field) WithID(id string) *Field {
	field.ID = id
	return field
}

// WithName sets the field's name attribute.
func (field *Field) WithName(name string) *Field {
	field.Name = name
	return field
}

// WithLabel sets the field's label text.
func (field *Field) WithLabel(label string) *Field {
	field.Label = label
	return field
}

// WithValue sets the field's value.
func (field *Field) WithValue(value string) *Field {
	field.Value = value
	return field
}

// WithType sets the field's type.
func (field *Field) WithType(fieldType string) *Field {
	field.Type = fieldType
	return field
}

// WithHelp sets the field's help text displayed below the input.
func (field *Field) WithHelp(help string) *Field {
	field.Help = help
	return field
}

// WithPlaceholder sets the field's placeholder text.
func (field *Field) WithPlaceholder(placeholder string) *Field {
	field.Placeholder = placeholder
	return field
}

// WithRequired marks the field as required.
func (field *Field) WithRequired() *Field {
	field.Required = true
	return field
}

// WithReadonly marks the field as readonly.
func (field *Field) WithReadonly() *Field {
	field.Readonly = true
	return field
}

// WithDisabled marks the field as disabled.
func (field *Field) WithDisabled() *Field {
	field.Disabled = true
	return field
}

// WithInvisible marks the field as invisible (hidden via CSS).
func (field *Field) WithInvisible() *Field {
	field.Invisible = true
	return field
}

// WithMultiple enables multi-select on select fields.
func (field *Field) WithMultiple() *Field {
	field.Multiple = true
	return field
}

// WithOptions sets the field's static options (for select, radio, etc.).
func (field *Field) WithOptions(options ...FieldOption) *Field {
	field.Options = options
	return field
}

// WithOptionsF sets a dynamic options provider function.
func (field *Field) WithOptionsF(optionsF func() []FieldOption) *Field {
	field.OptionsF = optionsF
	return field
}

// WithCustomInput sets a custom hb.Tag to use as the input element.
func (field *Field) WithCustomInput(input hb.TagInterface) *Field {
	field.CustomInput = input
	return field
}

// WithAttr sets a single custom HTML attribute on the field's input element.
func (field *Field) WithAttr(key, value string) *Field {
	if field.Attrs == nil {
		field.Attrs = map[string]string{}
	}
	field.Attrs[key] = value
	return field
}

// WithAttrs sets multiple custom HTML attributes on the field's input element.
func (field *Field) WithAttrs(attrs map[string]string) *Field {
	field.Attrs = attrs
	return field
}

// WithValidators sets the field's validators.
func (field *Field) WithValidators(validators ...Validator) *Field {
	field.Validators = validators
	return field
}

// WithTableOptions sets the table options for table-type fields.
func (field *Field) WithTableOptions(opts TableOptions) *Field {
	field.TableOptions = opts
	return field
}
