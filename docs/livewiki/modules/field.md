---
path: modules/field.md
page-type: module
summary: Documentation for the Field struct, rendering logic, and BuildFormGroup pipeline.
tags: [module, field, rendering, input-types]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Module: Field

## Purpose

The `Field` struct represents a single form field. It holds all configuration (type, name, label, value, validators, options, attributes) and implements `FieldInterface`. Its `BuildFormGroup()` method renders the complete form group: wrapper div, label, input element, error message, and help text.

## Key Types

### Field

```go
type Field struct {
    ID           string
    Type         string
    Name         string
    Label        string
    Help         string
    Options      []FieldOption
    OptionsF     func() []FieldOption
    Value        string
    Required     bool
    Readonly     bool
    Disabled     bool
    TableOptions TableOptions
    Placeholder  string
    Invisible    bool
    CustomInput  hb.TagInterface
    Attrs        map[string]string
    Multiple     bool
    Validators   []Validator
    theme        *Theme          // unexported, set by form
    errorMessage string          // unexported, set by form
}
```

### FieldOption

```go
type FieldOption struct {
    Key   string
    Value string
}
```

### TableColumn / TableOptions

```go
type TableColumn struct {
    Label string
    Width int
}

type TableOptions struct {
    Header          []TableColumn
    Rows            [][]Field
    RowAddButton    *hb.Tag
    RowDeleteButton *hb.Tag
}
```

## Rendering Pipeline

### BuildFormGroup(fileManagerURL string) *hb.Tag

1. **Raw fields** — Returns `hb.NewHTML(field.Value)` directly (no wrapper)
2. **Auto-ID** — If `ID` is empty, generates `"id_" + uid.HumanUid()`
3. **Create input** — Calls `fieldInput()` which switches on field type
4. **Readonly handling** — Selects get `disabled` + hidden input; others get `readonly`
5. **Disabled handling** — Adds `disabled` attribute + dimmed background
6. **Custom attributes** — Applies `field.Attrs` map
7. **Error display** — Adds `ErrorInputClass` to input + error div with `ErrorClass`
8. **Label** — Created with theme's `LabelClass`, required marker if needed
9. **Help text** — Paragraph with theme's `HelpClass`
10. **Invisible** — Adds `display:none` style

### fieldInput() Type Switch

| Field Type | Rendering |
|-----------|-----------|
| string, email, number, password, date, hidden, tel, url, color | `<input>` with appropriate `type` |
| datetime | `<input type="datetime-local">` |
| select | `<select>` with `<option>` children from Options + OptionsF |
| textarea | `<textarea>` |
| checkbox | `<input type="checkbox">` in wrapper div |
| radio | Multiple `<input type="radio">` from Options |
| file | `<input type="file">` |
| image | Image preview + textarea + file manager link |
| htmlarea | Textarea + Trumbowyg initialization script |
| blockeditor | CustomInput or fallback textarea with error |
| table | `<table>` with header, rows, add/delete buttons |
| raw | `hb.NewHTML(value)` |

## Type Check Methods

The Field provides boolean type-check helpers:

```go
field.IsString()      field.IsEmail()       field.IsNumber()
field.IsPassword()    field.IsDate()        field.IsDateTime()
field.IsSelect()      field.IsTextArea()    field.IsCheckbox()
field.IsRadio()       field.IsFile()        field.IsImage()
field.IsColor()       field.IsTel()         field.IsUrl()
field.IsHtmlArea()    field.IsBlockEditor() field.IsHidden()
field.IsTable()       field.IsRaw()
field.IsReadonly()    field.IsDisabled()    field.IsRequired()
```

## Interface Implementations

```go
var _ FieldInterface = (*Field)(nil)
var _ themeable = (*Field)(nil)
var _ errorAware = (*Field)(nil)
```

## Files

| File | Contents |
|------|----------|
| `field.go` | `Field` struct, `BuildFormGroup()`, `fieldInput()`, type checks, `TrumbowygScript()`, `clone()` |
| `field_fluent.go` | 19 `With*` fluent methods |
| `field_interface.go` | `FieldInterface` definition |
| `field_option.go` | `FieldOption` struct |
| `new_field.go` | `NewField(FieldOptions)` constructor |
| `field_constructors.go` | 18 type-safe constructors |

## See Also

- [Module: Field Constructors](field_constructors.md)
- [Module: Form](form.md)
- [Module: Validation](validation.md)
- [Module: Theme](theme.md)
- [API Reference](../api_reference.md)
