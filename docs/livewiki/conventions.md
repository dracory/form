---
path: conventions.md
page-type: reference
summary: Coding and documentation conventions used in the Dracory Form library.
tags: [conventions, style, coding-standards, patterns]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Conventions

## Package Conventions

- **Single package** — All code lives in `package form`. No sub-packages.
- **Module path** — `github.com/dracory/form`
- **Go version** — 1.25+ (specified in `go.mod`)

## Naming Conventions

### Constants

Field type constants use the `FORM_FIELD_TYPE_` prefix in SCREAMING_SNAKE_CASE:

```go
const FORM_FIELD_TYPE_STRING = "string"
const FORM_FIELD_TYPE_EMAIL = "email"
```

Exception: the internal repeater type uses lowercase:

```go
const formFieldTypeRepeater = "repeater"
```

### Types

- **Exported structs** — PascalCase: `Form`, `Field`, `Theme`, `HTMXConfig`, `FieldOption`, `FieldRowColumn`, `TableColumn`, `TableOptions`, `ValidationError`
- **Unexported structs** — camelCase: `fieldRepeater`, `fieldRow`
- **Interfaces** — PascalCase for public (`FieldInterface`), camelCase for internal (`formAware`, `themeable`, `errorAware`, `rowErrorAware`)

### Constructors

Two patterns coexist:

1. **Options-based** — `NewForm(FormOptions{...})`, `NewField(FieldOptions{...})`, `NewRepeater(RepeaterOptions{...})`
2. **Fluent** — `New()` for forms, `NewStringField(name, label)` for fields

### Fluent Methods

All fluent methods follow the `With*` prefix pattern and return the receiver pointer:

```go
func (field *Field) WithRequired() *Field {
    field.Required = true
    return field
}
```

### Getter/Setter Methods

Traditional getter/setter pairs use `Get*`/`Set*` prefix:

```go
func (field *Field) GetName() string
func (field *Field) SetName(fieldName string)
```

## File Organization

| File | Purpose |
|------|---------|
| `consts.go` | Constants only |
| `*_interface.go` | Interface definitions |
| `*_fluent.go` | Fluent builder methods |
| `new_*.go` | Options-based constructors |
| `*_constructors.go` | Type-safe factory functions |
| `*_test.go` | Tests (same package, not `_test` package) |

## Testing Conventions

- Tests are in the **same package** (`package form`, not `package form_test`)
- Test function names: `TestFieldTypeName`, `TestFormWithFeature`, `TestValidateRuleName`
- Expected HTML is compared using `strings.Contains()` for partial matches
- Test helpers in `test_helpers.go` use `t.Helper()` for clean stack traces
- Multiple expectations use a `[]string` slice iterated with a loop:

```go
expecteds := []string{
    `<div class="form-group mb-3">`,
    `type="text"`,
}
for _, expected := range expecteds {
    if !strings.Contains(html, expected) {
        t.Fatal(`Expected: `, expected, ` but was: `, html)
    }
}
```

## HTML Rendering Conventions

- All HTML is generated via `dracory/hb` — no Go templates
- `BuildFormGroup()` returns a complete form group (wrapper div + label + input + help + error)
- `Build()` returns the `<form>` tag with all children
- `.ToHTML()` converts the `hb.Tag` tree to a string
- Auto-generated IDs use `"id_" + uid.HumanUid()`

## Interface Implementation Verification

Compile-time interface checks use blank identifier assignments:

```go
var _ FieldInterface = (*Field)(nil)
var _ themeable = (*Field)(nil)
var _ errorAware = (*Field)(nil)
```

## Error Handling

- Validation errors are value types (`ValidationError`), not Go `error` interface
- Validators return `*ValidationError` (nil = valid)
- Form stores errors as `map[string]string` (field name → first error message)

## Dependencies

- **Minimal dependencies** — Only 4 direct dependencies
- **No standard library wrappers** — Uses `net/http`, `regexp`, `strconv`, `strings` directly
- **Functional utilities** — `samber/lo` for `lo.If`, `lo.Map`, `lo.Find`, `lo.ValueOr`

## See Also

- [Architecture](architecture.md)
- [Development](development.md)
- [API Reference](api_reference.md)
