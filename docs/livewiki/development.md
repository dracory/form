---
path: development.md
page-type: tutorial
summary: Development workflow, testing, and contributing guide for the Dracory Form library.
tags: [development, testing, contributing, workflow]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Development

## Prerequisites

- **Go 1.25+**
- **[Task](https://taskfile.dev/)** (optional, for task runner commands)

## Running Tests

```bash
# Using Go directly
go test ./...

# Using Task runner
task test
```

## Test Coverage

```bash
task cover
```

This generates a `coverage.out` file and opens an HTML coverage report.

## Static Analysis Tools

The project uses several static analysis tools, all configured in `Taskfile.yml`:

| Tool | Install | Run | Purpose |
|------|---------|-----|---------|
| errcheck | `task errcheck:install` | `task errcheck` | Detect unchecked errors |
| nilaway | `task nilaway:install` | `task nilaway` | Detect nil pointer issues |
| gocritic | `task gocritic:install` | `task gocritic` | Code style and correctness |
| golangci-lint | `task golangci-lint:install` | `task golangci-lint` | Meta-linter |
| gosec | `task gosec:install` | `task gosec` | Security scanner |

## CI/CD

GitHub Actions runs on push/PR to `main`:

1. Checkout code
2. Set up Go 1.24
3. `go build -v ./...`
4. `go test -v ./...`

See `.github/workflows/tests.yml`.

## Project Structure

```
form/
├── consts.go                  # Field type constants (FORM_FIELD_TYPE_*)
├── field.go                   # Field struct, rendering, BuildFormGroup
├── field_constructors.go      # NewStringField, NewEmailField, etc.
├── field_constructors_test.go
├── field_fluent.go            # Field With* builder methods
├── field_fluent_test.go
├── field_interface.go         # FieldInterface definition
├── field_option.go            # FieldOption struct
├── field_repeater.go          # Repeater field implementation
├── field_repeater_test.go
├── field_row.go               # FieldRow for grid layouts
├── field_row_test.go
├── field_test.go              # Field rendering tests
├── form.go                    # Form struct, Build(), internal interfaces
├── form_errors_test.go
├── form_fluent.go             # Form With* builder methods + New()
├── form_fluent_test.go
├── form_test.go
├── htmx.go                    # HTMXConfig struct
├── htmx_test.go
├── new_field.go               # NewField(FieldOptions) constructor
├── new_form.go                # NewForm(FormOptions) constructor
├── new_repeater.go            # NewRepeater(RepeaterOptions) constructor
├── test_helpers.go            # AssertFormContains, AssertValidationPasses, etc.
├── test_helpers_test.go
├── theme.go                   # Theme struct, Bootstrap5, Tailwind presets
├── theme_test.go
├── validation.go              # Validators and Form.Validate()
├── validation_rules_test.go
├── validation_test.go
├── Taskfile.yml               # Task runner config
├── go.mod / go.sum
└── .github/workflows/tests.yml
```

## Writing Tests

The library provides test helpers in `test_helpers.go`:

```go
func TestMyForm(t *testing.T) {
    f := form.New().
        WithFields(
            form.NewStringField("name", "Name").WithRequired(),
        )

    // Assert rendered HTML
    form.AssertFormContains(t, f, `type="text"`)
    form.AssertFormNotContains(t, f, `type="email"`)

    // Assert validation
    form.AssertValidationPasses(t, f, map[string]string{"name": "John"})
    form.AssertValidationFails(t, f, map[string]string{"name": ""})
    form.AssertValidationFailsOn(t, f, map[string]string{"name": ""}, "name")
    form.AssertValidationErrorCount(t, f, map[string]string{"name": ""}, 1)
}
```

## Adding a New Field Type

1. Add a constant in `consts.go`
2. Add a constructor in `field_constructors.go`
3. Add rendering logic in the `fieldInput()` switch in `field.go`
4. Add tests in `field_test.go` and `field_constructors_test.go`

## Adding a New Validator

1. Add the validator function in `validation.go`
2. Add tests in `validation_rules_test.go` or `validation_test.go`

## Adding a New Theme

1. Add a constructor function in `theme.go` (e.g., `ThemeMyFramework() *Theme`)
2. Add tests in `theme_test.go`

## Gitpod

The project supports Gitpod for cloud development. Click the "Open in Gitpod" badge in the README.

## See Also

- [Architecture](architecture.md)
- [Conventions](conventions.md)
- [Modules: Test Helpers](modules/test_helpers.md)
- [Troubleshooting](troubleshooting.md)
