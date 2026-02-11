---
path: modules/test_helpers.md
page-type: module
summary: Documentation for the test assertion helpers provided by the Dracory Form library.
tags: [module, testing, assertions, helpers]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Module: Test Helpers

## Purpose

The test helpers module provides assertion functions that simplify testing forms, fields, and validation. All helpers call `t.Helper()` so test failure messages point to the caller, not the helper function.

## Functions

### HTML Assertions

#### AssertFormContains

```go
func AssertFormContains(t *testing.T, form *Form, expected string)
```

Renders the form to HTML and checks that the output contains the expected string.

```go
f := form.New().WithFields(form.NewStringField("name", "Name"))
form.AssertFormContains(t, f, `type="text"`)
form.AssertFormContains(t, f, `name="name"`)
```

#### AssertFormNotContains

```go
func AssertFormNotContains(t *testing.T, form *Form, unexpected string)
```

Renders the form to HTML and checks that the output does NOT contain the given string.

```go
form.AssertFormNotContains(t, f, `type="email"`)
```

#### AssertFieldContains

```go
func AssertFieldContains(t *testing.T, field FieldInterface, expected string)
```

Renders a single field's form group (with empty `fileManagerURL`) and checks the output.

```go
field := form.NewEmailField("email", "Email")
form.AssertFieldContains(t, field, `type="email"`)
```

#### AssertFieldNotContains

```go
func AssertFieldNotContains(t *testing.T, field FieldInterface, unexpected string)
```

Renders a single field and checks the output does NOT contain the given string.

```go
form.AssertFieldNotContains(t, field, `required`)
```

### Validation Assertions

#### AssertValidationPasses

```go
func AssertValidationPasses(t *testing.T, form *Form, values map[string]string)
```

Validates the given values and fails if any validation errors occur.

```go
f := form.New().WithFields(
    form.NewStringField("name", "Name").WithRequired(),
)
form.AssertValidationPasses(t, f, map[string]string{"name": "John"})
```

#### AssertValidationFails

```go
func AssertValidationFails(t *testing.T, form *Form, values map[string]string) []ValidationError
```

Validates the given values and fails if validation passes. Returns the errors for further inspection.

```go
errs := form.AssertValidationFails(t, f, map[string]string{"name": ""})
// errs[0].Field == "name"
// errs[0].Message == "name is required"
```

#### AssertValidationFailsOn

```go
func AssertValidationFailsOn(t *testing.T, form *Form, values map[string]string, fieldName string)
```

Validates and checks that a specific field has a validation error.

```go
form.AssertValidationFailsOn(t, f, map[string]string{"name": ""}, "name")
```

#### AssertValidationErrorCount

```go
func AssertValidationErrorCount(t *testing.T, form *Form, values map[string]string, expectedCount int)
```

Validates and checks that exactly `expectedCount` errors are returned.

```go
form.AssertValidationErrorCount(t, f, map[string]string{
    "name":  "",
    "email": "bad",
}, 2)
```

## Complete Test Example

```go
func TestContactForm(t *testing.T) {
    f := form.New().
        WithID("contact").
        WithFields(
            form.NewStringField("name", "Name").
                WithRequired().
                WithValidators(form.ValidatorMinLength(2)),
            form.NewEmailField("email", "Email").
                WithRequired().
                WithValidators(form.ValidatorEmail()),
        )

    // Test rendering
    form.AssertFormContains(t, f, `id="contact"`)
    form.AssertFormContains(t, f, `name="name"`)
    form.AssertFormContains(t, f, `name="email"`)

    // Test valid submission
    form.AssertValidationPasses(t, f, map[string]string{
        "name":  "John Doe",
        "email": "john@example.com",
    })

    // Test empty submission
    form.AssertValidationErrorCount(t, f, map[string]string{
        "name":  "",
        "email": "",
    }, 2)

    // Test specific field failure
    form.AssertValidationFailsOn(t, f, map[string]string{
        "name":  "J",
        "email": "john@example.com",
    }, "name")
}
```

## Files

| File | Contents |
|------|----------|
| `test_helpers.go` | All 6 assertion helper functions |

## See Also

- [Development](../development.md)
- [Module: Validation](validation.md)
- [Module: Form](form.md)
- [API Reference](../api_reference.md)
