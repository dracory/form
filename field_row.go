package form

import "github.com/dracory/hb"

// FieldRowColumn represents a single column in a field row with its field and optional CSS class.
type FieldRowColumn struct {
	Field    FieldInterface
	ColClass string // e.g. "col-md-6", "col-4". Empty = auto equal width ("col")
}

// fieldRow is a pseudo-field that renders multiple fields in a single row.
type fieldRow struct {
	columns        []FieldRowColumn
	rowClass       string
	theme          *Theme
	errors         map[string]string
	fileManagerURL string
}

// NewFieldRow creates a row of fields. Each field gets an equal-width column by default.
func NewFieldRow(fields ...FieldInterface) *fieldRow {
	columns := make([]FieldRowColumn, len(fields))
	for i, f := range fields {
		columns[i] = FieldRowColumn{Field: f}
	}
	return &fieldRow{columns: columns, rowClass: "row"}
}

// NewFieldRowWithColumns creates a row with explicit column configurations.
func NewFieldRowWithColumns(columns ...FieldRowColumn) *fieldRow {
	return &fieldRow{columns: columns, rowClass: "row"}
}

// WithRowClass sets a custom CSS class for the row wrapper (default: "row").
func (r *fieldRow) WithRowClass(class string) *fieldRow {
	r.rowClass = class
	return r
}

// == INTERFACE ===============================================================

var _ FieldInterface = (*fieldRow)(nil)
var _ themeable = (*fieldRow)(nil)
var _ rowErrorAware = (*fieldRow)(nil)

func (r *fieldRow) setTheme(theme *Theme) {
	r.theme = theme
}

// setErrors sets the error map so the row can distribute errors to child fields.
func (r *fieldRow) setErrors(errors map[string]string) {
	r.errors = errors
}

// == IMPLEMENTATION OF FieldInterface ========================================

func (r *fieldRow) clone() FieldInterface {
	rowCopy := *r
	rowCopy.columns = make([]FieldRowColumn, len(r.columns))
	copy(rowCopy.columns, r.columns)
	return &rowCopy
}

func (r *fieldRow) BuildFormGroup(fileManagerURL string) *hb.Tag {
	row := hb.NewDiv().Class(r.rowClass)

	for _, col := range r.columns {
		colClass := col.ColClass
		if colClass == "" {
			colClass = "col"
		}

		// Pass theme to child field
		if r.theme != nil {
			if th, ok := col.Field.(themeable); ok {
				th.setTheme(r.theme)
			}
		}

		// Pass errors to child field
		if r.errors != nil {
			if ea, ok := col.Field.(errorAware); ok {
				if msg, exists := r.errors[col.Field.GetName()]; exists {
					ea.setError(msg)
				}
			}
		}

		colDiv := hb.NewDiv().Class(colClass).
			Child(col.Field.BuildFormGroup(fileManagerURL))
		row.Child(colDiv)
	}

	return row
}

// Stub implementations for FieldInterface (fieldRow is a layout, not a real field)

func (r *fieldRow) GetID() string                         { return "" }
func (r *fieldRow) SetID(fieldID string)                  {}
func (r *fieldRow) GetLabel() string                      { return "" }
func (r *fieldRow) SetLabel(fieldLabel string)            {}
func (r *fieldRow) GetName() string                       { return "" }
func (r *fieldRow) SetName(fieldName string)              {}
func (r *fieldRow) GetHelp() string                       { return "" }
func (r *fieldRow) SetHelp(fieldHelp string)              {}
func (r *fieldRow) GetOptions() []FieldOption             { return nil }
func (r *fieldRow) SetOptions(fieldOptions []FieldOption) {}
func (r *fieldRow) GetOptionsF() func() []FieldOption     { return nil }
func (r *fieldRow) SetOptionsF(f func() []FieldOption)    {}
func (r *fieldRow) GetRequired() bool                     { return false }
func (r *fieldRow) SetRequired(fieldRequired bool)        {}
func (r *fieldRow) GetType() string                       { return "row" }
func (r *fieldRow) SetType(fieldType string)              {}
func (r *fieldRow) GetValue() string                      { return "" }
func (r *fieldRow) SetValue(fieldValue string)            {}
