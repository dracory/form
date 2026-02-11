---
path: cheatsheet.md
page-type: reference
summary: Quick reference for common operations with the Dracory Form library.
tags: [cheatsheet, quick-reference, patterns, examples]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Cheatsheet

## Create a Form

```go
// Fluent API
f := form.New().
    WithID("myForm").
    WithMethod("POST").
    WithAction("/submit").
    WithFields(...)

// Options API
f := form.NewForm(form.FormOptions{
    ID:     "myForm",
    Method: "POST",
    Fields: []form.FieldInterface{...},
})
```

## Common Field Types

```go
form.NewStringField("name", "Name")
form.NewEmailField("email", "Email")
form.NewPasswordField("pass", "Password")
form.NewNumberField("age", "Age")
form.NewDateField("dob", "Date of Birth")
form.NewDateTimeField("ts", "Timestamp")
form.NewTextAreaField("bio", "Biography")
form.NewHiddenField("token", "abc123")       // name, VALUE (not label)
form.NewCheckboxField("agree", "I Agree")
form.NewColorField("color", "Pick Color")
form.NewTelField("phone", "Phone")
form.NewURLField("website", "Website")
form.NewFileField("doc", "Upload Document")
form.NewImageField("avatar", "Avatar")
form.NewHtmlAreaField("content", "Content")   // WYSIWYG
form.NewRawField("<hr/>")                     // Raw HTML
```

## Select & Radio with Options

```go
opts := []form.FieldOption{
    {Key: "us", Value: "United States"},
    {Key: "uk", Value: "United Kingdom"},
}

form.NewSelectField("country", "Country", opts)
form.NewRadioField("country", "Country", opts)
```

## Field Modifiers

```go
field.WithRequired()
field.WithReadonly()
field.WithDisabled()
field.WithInvisible()
field.WithMultiple()                          // multi-select
field.WithPlaceholder("Enter value...")
field.WithHelp("Help text below input")
field.WithValue("default value")
field.WithID("custom_id")
field.WithAttr("data-custom", "value")
field.WithAttrs(map[string]string{"k": "v"})
```

## Validation

```go
// Add validators to fields
field.WithValidators(
    form.ValidatorRequired(),
    form.ValidatorMinLength(3),
    form.ValidatorMaxLength(100),
    form.ValidatorEmail(),
    form.ValidatorURL(),
    form.ValidatorIP(),
    form.ValidatorUUID(),
    form.ValidatorAlphaNumeric(),
    form.ValidatorMin(0),
    form.ValidatorMax(999),
    form.ValidatorPattern(`^\d{3}$`, "must be 3 digits"),
    form.ValidatorOneOf("a", "b", "c"),
    form.ValidatorCustom(func(v string) string {
        if v == "bad" { return "value is bad" }
        return ""
    }),
)

// Validate submitted values
errors := f.Validate(map[string]string{
    "name":  "John",
    "email": "john@example.com",
})

// Render with inline errors
html := f.Build().ToHTML()
```

## Themes

```go
form.New().WithTheme(form.ThemeBootstrap5())  // default
form.New().WithTheme(form.ThemeTailwind())
form.New().WithTheme(&form.Theme{             // custom
    InputClass: "my-input",
    LabelClass: "my-label",
    // ...
})
```

## HTMX

```go
// Simple
form.New().
    WithHxPost("/submit").
    WithHxTarget("#result").
    WithHxSwap("innerHTML")

// Full config
form.New().WithHTMX(form.HTMXConfig{
    Post:      "/submit",
    Target:    "#result",
    Swap:      "innerHTML",
    Indicator: "#spinner",
    Confirm:   "Sure?",
    Validate:  true,
})
```

## Grid Layouts

```go
// Equal-width columns
form.NewFieldRow(
    form.NewStringField("first", "First Name"),
    form.NewStringField("last", "Last Name"),
)

// Custom column widths
form.NewFieldRowWithColumns(
    form.FieldRowColumn{Field: form.NewStringField("city", "City"), ColClass: "col-md-8"},
    form.FieldRowColumn{Field: form.NewStringField("zip", "ZIP"), ColClass: "col-md-4"},
)
```

## Repeater

```go
form.NewRepeater(form.RepeaterOptions{
    Name:                "items",
    Label:               "Items",
    RepeaterAddUrl:      "/repeater/add",
    RepeaterRemoveUrl:   "/repeater/remove",
    RepeaterMoveUpUrl:   "/repeater/up",
    RepeaterMoveDownUrl: "/repeater/down",
    Fields: []form.FieldInterface{
        form.NewStringField("item_name", "Item Name"),
        form.NewNumberField("qty", "Quantity"),
    },
    Values: []map[string]string{
        {"item_name": "Widget", "qty": "5"},
    },
})
```

## Render to HTML

```go
html := f.Build().ToHTML()
```

## Test Helpers

```go
form.AssertFormContains(t, f, `type="text"`)
form.AssertFormNotContains(t, f, `type="email"`)
form.AssertFieldContains(t, field, `class="form-control"`)
form.AssertValidationPasses(t, f, values)
form.AssertValidationFails(t, f, values)
form.AssertValidationFailsOn(t, f, values, "fieldName")
form.AssertValidationErrorCount(t, f, values, 2)
```

## Task Runner Commands

```bash
task test              # Run tests
task cover             # Coverage report
task errcheck          # Check unchecked errors
task nilaway           # Check nil pointers
task gocritic          # Code style checks
task golangci-lint     # Meta-linter
task gosec             # Security scan
```

## See Also

- [API Reference](api_reference.md)
- [Getting Started](getting_started.md)
- [Configuration](configuration.md)
