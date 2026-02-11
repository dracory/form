---
path: modules/field_constructors.md
page-type: module
summary: Documentation for the 18 type-safe field constructors that simplify field creation.
tags: [module, field, constructors, factory]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Module: Field Constructors

## Purpose

Type-safe factory functions that create `*Field` instances with the correct type constant pre-set. These eliminate the need to reference `FORM_FIELD_TYPE_*` constants directly, reducing errors and improving readability.

## Constructors

### Text Input Fields

```go
// Standard text input
form.NewStringField("name", "Full Name")

// Email input with browser validation
form.NewEmailField("email", "Email Address")

// Numeric input
form.NewNumberField("age", "Age")

// Password input (masked)
form.NewPasswordField("password", "Password")

// Telephone input
form.NewTelField("phone", "Phone Number")

// URL input
form.NewURLField("website", "Website")

// Color picker
form.NewColorField("color", "Favorite Color")
```

### Date/Time Fields

```go
// Date picker
form.NewDateField("dob", "Date of Birth")

// DateTime picker (datetime-local)
form.NewDateTimeField("event_time", "Event Time")
```

### Choice Fields

```go
opts := []form.FieldOption{
    {Key: "us", Value: "United States"},
    {Key: "uk", Value: "United Kingdom"},
}

// Dropdown select
form.NewSelectField("country", "Country", opts)

// Radio button group
form.NewRadioField("gender", "Gender", opts)

// Checkbox (single toggle)
form.NewCheckboxField("agree", "I Agree")
```

### Text Area Fields

```go
// Plain textarea
form.NewTextAreaField("bio", "Biography")

// WYSIWYG editor (Trumbowyg)
form.NewHtmlAreaField("content", "Content")
```

### File/Media Fields

```go
// File upload
form.NewFileField("document", "Upload Document")

// Image with preview and file manager
form.NewImageField("avatar", "Profile Image")
```

### Special Fields

```go
// Hidden input (note: second param is VALUE, not label)
form.NewHiddenField("csrf_token", "abc123")

// Raw HTML (rendered as-is, no wrapper)
form.NewRawField("<hr class='my-3' />")
```

## Constructor Signatures

| Constructor | Parameters | Notes |
|-------------|-----------|-------|
| `NewStringField(name, label)` | `string, string` | |
| `NewEmailField(name, label)` | `string, string` | |
| `NewNumberField(name, label)` | `string, string` | |
| `NewPasswordField(name, label)` | `string, string` | |
| `NewTelField(name, label)` | `string, string` | |
| `NewURLField(name, label)` | `string, string` | |
| `NewColorField(name, label)` | `string, string` | |
| `NewDateField(name, label)` | `string, string` | |
| `NewDateTimeField(name, label)` | `string, string` | |
| `NewSelectField(name, label, options)` | `string, string, []FieldOption` | |
| `NewRadioField(name, label, options)` | `string, string, []FieldOption` | |
| `NewCheckboxField(name, label)` | `string, string` | |
| `NewTextAreaField(name, label)` | `string, string` | |
| `NewHtmlAreaField(name, label)` | `string, string` | |
| `NewFileField(name, label)` | `string, string` | |
| `NewImageField(name, label)` | `string, string` | |
| `NewHiddenField(name, value)` | `string, string` | **value**, not label |
| `NewRawField(value)` | `string` | Raw HTML content |

## Chaining with Fluent Methods

All constructors return `*Field`, so you can chain `With*` methods:

```go
form.NewStringField("username", "Username").
    WithRequired().
    WithPlaceholder("Enter username").
    WithValidators(
        form.ValidatorMinLength(3),
        form.ValidatorAlphaNumeric(),
    ).
    WithHelp("3-20 alphanumeric characters")
```

## Generic Constructor

For cases where you need full control, use `NewField()`:

```go
form.NewField(form.FieldOptions{
    Type:        form.FORM_FIELD_TYPE_STRING,
    Name:        "custom",
    Label:       "Custom Field",
    Required:    true,
    Placeholder: "...",
    Attrs:       map[string]string{"data-x": "y"},
})
```

## Files

| File | Contents |
|------|----------|
| `field_constructors.go` | All 18 type-safe constructors |
| `new_field.go` | `NewField(FieldOptions)` generic constructor |

## See Also

- [Module: Field](field.md)
- [API Reference](../api_reference.md)
- [Cheatsheet](../cheatsheet.md)
