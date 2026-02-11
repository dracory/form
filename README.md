# Form

<a href="https://gitpod.io/#https://github.com/dracory/form" target="_blank" style="float:right;"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

A Go library for building HTML forms with a fluent API, validation, theming, HTMX integration, and grid layouts.

## License

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0). You can find a copy of the license at [https://www.gnu.org/licenses/agpl-3.0.en.html](https://www.gnu.org/licenses/agpl-3.0.txt)

For commercial use, please use my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## Installation

```bash
go get github.com/dracory/form
```

## Quick Start

```golang
import "github.com/dracory/form"

f := form.New().
    WithID("contactForm").
    WithFields(
        form.NewStringField("name", "Full Name").WithRequired(),
        form.NewEmailField("email", "Email").WithPlaceholder("you@example.com"),
    )

html := f.Build().ToHTML()
```

## Documentation

- [Field Types](docs/field-types.md) - All 18 supported field types and their constructors
- [Fluent API](docs/fluent-api.md) - Form and Field builder methods for chaining
- [Validation](docs/validation.md) - 13 built-in validators, custom validators, inline error display
- [Theming](docs/theming.md) - Bootstrap 5, Tailwind CSS, and custom themes
- [HTMX Integration](docs/htmx.md) - Simple attributes and structured HTMXConfig
- [Field Rows](docs/field-rows.md) - Grid layouts with multi-column rows
- [Repeater](docs/repeater.md) - Dynamic add/remove field groups
- [Test Helpers](docs/test-helpers.md) - Assertion helpers for testing forms
- [Advanced](docs/advanced.md) - Trumbowyg WYSIWYG config, legacy API