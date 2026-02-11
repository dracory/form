---
path: overview.md
page-type: overview
summary: High-level overview of the Dracory Form library for building HTML forms in Go.
tags: [overview, introduction, architecture]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Dracory Form Library

A Go library for building HTML forms with a fluent API, validation, theming, HTMX integration, and grid layouts.

## What It Does

The `github.com/dracory/form` package lets you programmatically construct HTML forms in Go. Instead of writing raw HTML templates, you describe forms using Go structs and builder methods, then render them to HTML via the [dracory/hb](https://github.com/dracory/hb) HTML builder.

## Key Features

- **18 Field Types** — String, email, number, password, date, datetime, select, textarea, checkbox, radio, file, image, color, tel, URL, HTML area (WYSIWYG), block editor, raw HTML, and table
- **Fluent Builder API** — Chain `With*` methods for readable, concise form construction
- **Type-Safe Constructors** — `NewStringField()`, `NewEmailField()`, etc. prevent type constant mistakes
- **13 Built-in Validators** — Required, min/max length, min/max value, pattern, email, URL, IP, UUID, alphanumeric, one-of, and custom
- **Inline Error Display** — Validation errors render next to their fields automatically
- **Theming** — Built-in Bootstrap 5 and Tailwind CSS themes, plus custom theme support
- **HTMX Integration** — Simple `hx-post`/`hx-target`/`hx-swap` helpers and a full `HTMXConfig` struct
- **Grid Layouts** — `NewFieldRow()` for multi-column field rows
- **Repeater Fields** — Dynamic add/remove field groups with HTMX-powered actions
- **Test Helpers** — `AssertFormContains`, `AssertValidationPasses`, etc.

## Quick Example

```go
import "github.com/dracory/form"

f := form.New().
    WithID("contactForm").
    WithFields(
        form.NewStringField("name", "Full Name").WithRequired(),
        form.NewEmailField("email", "Email").WithPlaceholder("you@example.com"),
    )

html := f.Build().ToHTML()
```

## Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/dracory/hb` | HTML builder for generating HTML tags |
| `github.com/dracory/uid` | Unique ID generation for field IDs |
| `github.com/samber/lo` | Functional utilities (lo.If, lo.Map, etc.) |
| `github.com/spf13/cast` | Type casting (used in repeater) |

## License

AGPL-3.0. Commercial licenses available via [lesichkov.co.uk/contact](https://lesichkov.co.uk/contact).

## See Also

- [Getting Started](getting_started.md)
- [Architecture](architecture.md)
- [API Reference](api_reference.md)
- [Cheatsheet](cheatsheet.md)
