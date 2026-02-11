---
path: modules/field_row.md
page-type: module
summary: Documentation for the FieldRow layout component that renders multiple fields in a grid row.
tags: [module, field-row, layout, grid, columns]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Module: Field Row

## Purpose

The `fieldRow` is a layout pseudo-field that renders multiple fields side-by-side in a CSS grid row. It implements `FieldInterface` so it can be used anywhere a field is expected, but it is not an input field itself — it is a container for other fields.

## Key Types

### FieldRowColumn

```go
type FieldRowColumn struct {
    Field    FieldInterface
    ColClass string // e.g. "col-md-6", "col-4". Empty = "col" (auto equal width)
}
```

### fieldRow (unexported)

```go
type fieldRow struct {
    columns        []FieldRowColumn
    rowClass       string           // default: "row"
    theme          *Theme
    errors         map[string]string
    fileManagerURL string
}
```

## Constructors

### Equal-Width Columns

```go
row := form.NewFieldRow(
    form.NewStringField("first_name", "First Name"),
    form.NewStringField("last_name", "Last Name"),
)
```

Renders each field in an equal-width `"col"` div inside a `"row"` wrapper.

### Custom Column Widths

```go
row := form.NewFieldRowWithColumns(
    form.FieldRowColumn{
        Field:    form.NewStringField("city", "City"),
        ColClass: "col-md-8",
    },
    form.FieldRowColumn{
        Field:    form.NewStringField("zip", "ZIP Code"),
        ColClass: "col-md-4",
    },
)
```

### Custom Row Class

```go
row := form.NewFieldRow(field1, field2).
    WithRowClass("row g-3") // adds gutter
```

## Rendered HTML Structure

```html
<div class="row">
  <div class="col">
    <!-- First field's BuildFormGroup output -->
  </div>
  <div class="col">
    <!-- Second field's BuildFormGroup output -->
  </div>
</div>
```

With custom columns:

```html
<div class="row">
  <div class="col-md-8">
    <!-- City field -->
  </div>
  <div class="col-md-4">
    <!-- ZIP field -->
  </div>
</div>
```

## Theme and Error Propagation

The `fieldRow` propagates theme and errors to its child fields:

1. **Theme** — If the row receives a theme (via `setTheme()`), it passes it to each child field that implements `themeable`
2. **Errors** — If the row receives an error map (via `setErrors()`), it checks each child field's name and injects the matching error via `errorAware.setError()`

This ensures that fields inside rows receive the same theme and error treatment as top-level fields.

## Interface Implementations

```go
var _ FieldInterface = (*fieldRow)(nil)
var _ themeable = (*fieldRow)(nil)
var _ rowErrorAware = (*fieldRow)(nil)
```

The `FieldInterface` methods are stub implementations (return zero values) since the row is a layout container, not a real field.

## Usage in Forms

```go
f := form.New().
    WithFields(
        form.NewFieldRow(
            form.NewStringField("first_name", "First Name").WithRequired(),
            form.NewStringField("last_name", "Last Name").WithRequired(),
        ),
        form.NewFieldRowWithColumns(
            form.FieldRowColumn{
                Field:    form.NewEmailField("email", "Email"),
                ColClass: "col-md-8",
            },
            form.FieldRowColumn{
                Field:    form.NewTelField("phone", "Phone"),
                ColClass: "col-md-4",
            },
        ),
        form.NewTextAreaField("notes", "Notes"),
    )
```

## Files

| File | Contents |
|------|----------|
| `field_row.go` | `fieldRow` struct, `FieldRowColumn`, `NewFieldRow()`, `NewFieldRowWithColumns()`, `WithRowClass()`, stub `FieldInterface` methods |

## See Also

- [Module: Field](field.md)
- [Module: Form](form.md)
- [Module: Theme](theme.md)
- [API Reference](../api_reference.md)
- [Cheatsheet](../cheatsheet.md)
