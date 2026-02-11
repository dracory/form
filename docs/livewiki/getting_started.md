---
path: getting_started.md
page-type: tutorial
summary: Installation, setup, and quick start guide for the Dracory Form library.
tags: [getting-started, installation, tutorial, quick-start]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Getting Started

## Prerequisites

- **Go 1.25+** (as specified in `go.mod`)

## Installation

```bash
go get github.com/dracory/form
```

## Quick Start

### 1. Create a Simple Form

```go
package main

import (
    "fmt"
    "github.com/dracory/form"
)

func main() {
    f := form.New().
        WithID("contactForm").
        WithFields(
            form.NewStringField("name", "Full Name").WithRequired(),
            form.NewEmailField("email", "Email").WithPlaceholder("you@example.com"),
            form.NewTextAreaField("message", "Message"),
        )

    html := f.Build().ToHTML()
    fmt.Println(html)
}
```

### 2. Add Validation

```go
f := form.New().
    WithFields(
        form.NewStringField("username", "Username").
            WithRequired().
            WithValidators(
                form.ValidatorMinLength(3),
                form.ValidatorMaxLength(20),
                form.ValidatorAlphaNumeric(),
            ),
        form.NewEmailField("email", "Email").
            WithRequired().
            WithValidators(form.ValidatorEmail()),
        form.NewPasswordField("password", "Password").
            WithRequired().
            WithValidators(form.ValidatorMinLength(8)),
    )

// Validate submitted values
errors := f.Validate(map[string]string{
    "username": "jo",
    "email":    "not-an-email",
    "password": "short",
})

// errors will contain validation failures
// The form also stores errors internally for inline display
html := f.Build().ToHTML() // errors rendered next to fields
```

### 3. Use HTMX Integration

```go
f := form.New().
    WithID("myForm").
    WithHxPost("/api/submit").
    WithHxTarget("#result").
    WithHxSwap("innerHTML").
    WithFields(
        form.NewStringField("name", "Name").WithRequired(),
    )
```

Or use the structured `HTMXConfig`:

```go
f := form.New().
    WithID("myForm").
    WithHTMX(form.HTMXConfig{
        Post:      "/api/submit",
        Target:    "#result",
        Swap:      "innerHTML",
        Indicator: "#spinner",
        Validate:  true,
    }).
    WithFields(
        form.NewStringField("name", "Name").WithRequired(),
    )
```

### 4. Apply a Theme

```go
// Use Tailwind CSS instead of the default Bootstrap 5
f := form.New().
    WithTheme(form.ThemeTailwind()).
    WithFields(
        form.NewStringField("name", "Name"),
    )
```

### 5. Multi-Column Layouts

```go
f := form.New().
    WithFields(
        form.NewFieldRow(
            form.NewStringField("first_name", "First Name"),
            form.NewStringField("last_name", "Last Name"),
        ),
        form.NewEmailField("email", "Email"),
    )
```

### 6. Using the Options-Based Constructor

If you prefer the options pattern over the fluent API:

```go
f := form.NewForm(form.FormOptions{
    ID:     "myForm",
    Method: "POST",
    Fields: []form.FieldInterface{
        form.NewField(form.FieldOptions{
            Name:     "email",
            Label:    "Email",
            Type:     form.FORM_FIELD_TYPE_EMAIL,
            Required: true,
        }),
    },
})
```

## See Also

- [Overview](overview.md)
- [API Reference](api_reference.md)
- [Field Constructors](modules/field_constructors.md)
- [Cheatsheet](cheatsheet.md)
