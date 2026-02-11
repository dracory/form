---
path: api_reference.md
page-type: reference
summary: Complete API reference for all public types, functions, and methods in the Dracory Form library.
tags: [api, reference, functions, methods, types]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# API Reference

## Form

### Constructors

| Function | Description |
|----------|-------------|
| `New() *Form` | Creates an empty form with POST method. Use `With*` methods to configure. |
| `NewForm(opts FormOptions) *Form` | Creates a form from an options struct. Defaults to POST. |

### FormOptions

```go
type FormOptions struct {
    ActionURL      string
    ClassName      string
    ID             string
    Fields         []FieldInterface
    FileManagerURL string
    Method         string
    HxPost         string
    HxTarget       string
    HxSwap         string
}
```

### Fluent Methods (Form)

| Method | Returns | Description |
|--------|---------|-------------|
| `WithID(id string)` | `*Form` | Sets the HTML `id` attribute |
| `WithClass(className string)` | `*Form` | Sets the CSS class |
| `WithMethod(method string)` | `*Form` | Sets the HTTP method |
| `WithAction(actionURL string)` | `*Form` | Sets the form action URL |
| `WithFields(fields ...FieldInterface)` | `*Form` | Sets the form fields |
| `WithFileManager(url string)` | `*Form` | Sets the file manager URL for image fields |
| `WithHxPost(url string)` | `*Form` | Sets `hx-post` attribute |
| `WithHxTarget(target string)` | `*Form` | Sets `hx-target` attribute |
| `WithHxSwap(swap string)` | `*Form` | Sets `hx-swap` attribute |
| `WithTheme(theme *Theme)` | `*Form` | Sets the rendering theme |
| `WithErrors(errors map[string]string)` | `*Form` | Sets inline error messages |
| `WithHTMX(config HTMXConfig)` | `*Form` | Sets HTMX config (extended attributes) |

### Getter/Setter Methods (Form)

| Method | Description |
|--------|-------------|
| `AddField(field FieldInterface)` | Appends a field |
| `GetFields() []FieldInterface` | Returns all fields |
| `SetFields(fields []FieldInterface)` | Replaces all fields |
| `GetFileManagerURL() string` | Returns file manager URL |
| `SetFileManagerURL(url string)` | Sets file manager URL |
| `SetErrors(errors map[string]string)` | Sets error messages |
| `GetErrors() map[string]string` | Returns error messages |

### Core Methods (Form)

| Method | Returns | Description |
|--------|---------|-------------|
| `Build()` | `*hb.Tag` | Renders the form to an HTML tag. Call `.ToHTML()` on the result. |
| `Validate(values map[string]string)` | `[]ValidationError` | Validates values against field validators. Stores errors for inline display. |

---

## Field

### Type-Safe Constructors

| Function | Field Type | Parameters |
|----------|-----------|------------|
| `NewStringField(name, label)` | `string` (text input) | name, label |
| `NewEmailField(name, label)` | `email` | name, label |
| `NewNumberField(name, label)` | `number` | name, label |
| `NewPasswordField(name, label)` | `password` | name, label |
| `NewHiddenField(name, value)` | `hidden` | name, **value** (not label) |
| `NewDateField(name, label)` | `date` | name, label |
| `NewDateTimeField(name, label)` | `datetime` | name, label |
| `NewSelectField(name, label, options)` | `select` | name, label, `[]FieldOption` |
| `NewTextAreaField(name, label)` | `textarea` | name, label |
| `NewCheckboxField(name, label)` | `checkbox` | name, label |
| `NewRadioField(name, label, options)` | `radio` | name, label, `[]FieldOption` |
| `NewFileField(name, label)` | `file` | name, label |
| `NewImageField(name, label)` | `image` | name, label |
| `NewColorField(name, label)` | `color` | name, label |
| `NewTelField(name, label)` | `tel` | name, label |
| `NewURLField(name, label)` | `url` | name, label |
| `NewHtmlAreaField(name, label)` | `htmlarea` (WYSIWYG) | name, label |
| `NewRawField(value)` | `raw` (HTML as-is) | value |

### Generic Constructor

```go
NewField(opts FieldOptions) *Field
```

### Fluent Methods (Field)

| Method | Returns | Description |
|--------|---------|-------------|
| `WithID(id string)` | `*Field` | Sets the HTML `id` |
| `WithName(name string)` | `*Field` | Sets the field name |
| `WithLabel(label string)` | `*Field` | Sets the label text |
| `WithValue(value string)` | `*Field` | Sets the field value |
| `WithType(fieldType string)` | `*Field` | Sets the field type |
| `WithHelp(help string)` | `*Field` | Sets help text below the input |
| `WithPlaceholder(placeholder string)` | `*Field` | Sets placeholder text |
| `WithRequired()` | `*Field` | Marks as required |
| `WithReadonly()` | `*Field` | Marks as readonly |
| `WithDisabled()` | `*Field` | Marks as disabled |
| `WithInvisible()` | `*Field` | Hides via CSS (`display:none`) |
| `WithMultiple()` | `*Field` | Enables multi-select |
| `WithOptions(options ...FieldOption)` | `*Field` | Sets static options |
| `WithOptionsF(fn func() []FieldOption)` | `*Field` | Sets dynamic options provider |
| `WithCustomInput(input hb.TagInterface)` | `*Field` | Sets a custom input element |
| `WithAttr(key, value string)` | `*Field` | Sets a single HTML attribute |
| `WithAttrs(attrs map[string]string)` | `*Field` | Sets multiple HTML attributes |
| `WithValidators(validators ...Validator)` | `*Field` | Sets validators |
| `WithTableOptions(opts TableOptions)` | `*Field` | Sets table options |

### Field Methods

| Method | Returns | Description |
|--------|---------|-------------|
| `BuildFormGroup(fileManagerURL string)` | `*hb.Tag` | Renders the complete form group (label + input + help + error) |
| `IsString()`, `IsEmail()`, `IsSelect()`, etc. | `bool` | Type check helpers |
| `IsReadonly()`, `IsDisabled()`, `IsRequired()` | `bool` | State check helpers |
| `TrumbowygScript()` | `string` | Returns Trumbowyg WYSIWYG init JS |

---

## FieldOption

```go
type FieldOption struct {
    Key   string
    Value string
}
```

Used for select, radio, and other option-based fields.

---

## FieldRow

| Function | Description |
|----------|-------------|
| `NewFieldRow(fields ...FieldInterface) *fieldRow` | Creates a row with equal-width columns |
| `NewFieldRowWithColumns(columns ...FieldRowColumn) *fieldRow` | Creates a row with explicit column configs |

```go
type FieldRowColumn struct {
    Field    FieldInterface
    ColClass string // e.g. "col-md-6". Empty = "col" (auto equal width)
}
```

| Method | Description |
|--------|-------------|
| `WithRowClass(class string) *fieldRow` | Sets the row wrapper CSS class (default: `"row"`) |

---

## Repeater

```go
NewRepeater(opts RepeaterOptions) *fieldRepeater
```

```go
type RepeaterOptions struct {
    Label               string
    Type                string
    Name                string
    Value               string
    Help                string
    Fields              []FieldInterface
    Values              []map[string]string
    RepeaterAddUrl      string
    RepeaterMoveUpUrl   string
    RepeaterMoveDownUrl string
    RepeaterRemoveUrl   string
}
```

---

## Validation

### Validator Type

```go
type Validator func(fieldName string, value string) *ValidationError

type ValidationError struct {
    Field   string
    Message string
}
```

### Built-in Validators

| Validator | Description |
|-----------|-------------|
| `ValidatorRequired()` | Non-empty value |
| `ValidatorMinLength(n int)` | Minimum character count |
| `ValidatorMaxLength(n int)` | Maximum character count |
| `ValidatorMin(n float64)` | Minimum numeric value |
| `ValidatorMax(n float64)` | Maximum numeric value |
| `ValidatorPattern(pattern, message string)` | Regex match |
| `ValidatorEmail()` | Valid email format |
| `ValidatorURL()` | Valid URL format |
| `ValidatorIP()` | Valid IPv4 address |
| `ValidatorUUID()` | Valid UUID format |
| `ValidatorAlphaNumeric()` | Letters and numbers only |
| `ValidatorOneOf(allowed ...string)` | Value in allowed set |
| `ValidatorCustom(fn func(string) string)` | Custom validation function |

---

## Theme

```go
type Theme struct {
    FormGroupClass     string
    LabelClass         string
    InputClass         string
    SelectClass        string
    TextAreaClass      string
    CheckboxWrapClass  string
    CheckboxInputClass string
    RadioWrapClass     string
    RadioInputClass    string
    RadioLabelClass    string
    FileInputClass     string
    HelpClass          string
    RequiredClass      string
    RequiredMarker     string
    TableClass         string
    ErrorClass         string
    ErrorInputClass    string
}
```

| Function | Description |
|----------|-------------|
| `ThemeBootstrap5() *Theme` | Returns Bootstrap 5 theme (default) |
| `ThemeTailwind() *Theme` | Returns Tailwind CSS theme |

---

## HTMXConfig

```go
type HTMXConfig struct {
    Post        string
    Get         string
    Target      string
    Swap        string
    Trigger     string
    Indicator   string
    Confirm     string
    Sync        string
    Validate    bool
    DisabledElt string
    Encoding    string
    PushURL     string
}
```

---

## Test Helpers

| Function | Description |
|----------|-------------|
| `AssertFormContains(t, form, expected)` | Checks rendered HTML contains string |
| `AssertFormNotContains(t, form, unexpected)` | Checks rendered HTML does NOT contain string |
| `AssertFieldContains(t, field, expected)` | Checks field HTML contains string |
| `AssertFieldNotContains(t, field, unexpected)` | Checks field HTML does NOT contain string |
| `AssertValidationPasses(t, form, values)` | Asserts validation passes |
| `AssertValidationFails(t, form, values)` | Asserts validation fails, returns errors |
| `AssertValidationFailsOn(t, form, values, fieldName)` | Asserts validation fails on specific field |
| `AssertValidationErrorCount(t, form, values, count)` | Asserts exact error count |

---

## Field Type Constants

| Constant | Value |
|----------|-------|
| `FORM_FIELD_TYPE_STRING` | `"string"` |
| `FORM_FIELD_TYPE_EMAIL` | `"email"` |
| `FORM_FIELD_TYPE_NUMBER` | `"number"` |
| `FORM_FIELD_TYPE_PASSWORD` | `"password"` |
| `FORM_FIELD_TYPE_HIDDEN` | `"hidden"` |
| `FORM_FIELD_TYPE_DATE` | `"date"` |
| `FORM_FIELD_TYPE_DATETIME` | `"datetime"` |
| `FORM_FIELD_TYPE_SELECT` | `"select"` |
| `FORM_FIELD_TYPE_TEXTAREA` | `"textarea"` |
| `FORM_FIELD_TYPE_CHECKBOX` | `"checkbox"` |
| `FORM_FIELD_TYPE_RADIO` | `"radio"` |
| `FORM_FIELD_TYPE_FILE` | `"file"` |
| `FORM_FIELD_TYPE_IMAGE` | `"image"` |
| `FORM_FIELD_TYPE_COLOR` | `"color"` |
| `FORM_FIELD_TYPE_TEL` | `"tel"` |
| `FORM_FIELD_TYPE_URL` | `"url"` |
| `FORM_FIELD_TYPE_HTMLAREA` | `"htmlarea"` |
| `FORM_FIELD_TYPE_BLOCKEDITOR` | `"blockeditor"` |
| `FORM_FIELD_TYPE_RAW` | `"raw"` |
| `FORM_FIELD_TYPE_TABLE` | `"table"` |

## See Also

- [Overview](overview.md)
- [Architecture](architecture.md)
- [Cheatsheet](cheatsheet.md)
- [Modules: Form](modules/form.md)
- [Modules: Field](modules/field.md)
