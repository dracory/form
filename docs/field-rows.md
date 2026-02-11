# Field Rows (Grid Layouts)

Group fields into rows for multi-column layouts.

## Equal-Width Columns

```golang
f := form.New().WithFields(
    form.NewFieldRow(
        form.NewStringField("first", "First Name"),
        form.NewStringField("last", "Last Name"),
    ),
)
```

Renders a `<div class="row">` with equal-width `<div class="col">` columns.

## Custom Column Widths

```golang
f := form.New().WithFields(
    form.NewFieldRowWithColumns(
        form.FieldRowColumn{Field: form.NewStringField("city", "City"), ColClass: "col-md-8"},
        form.FieldRowColumn{Field: form.NewStringField("zip", "ZIP"), ColClass: "col-md-4"},
    ),
)
```

## Custom Row Class

```golang
form.NewFieldRow(
    form.NewStringField("a", "A"),
    form.NewStringField("b", "B"),
).WithRowClass("row g-3")
```

## Mixing Rows and Regular Fields

```golang
f := form.New().WithFields(
    form.NewStringField("title", "Title"),
    form.NewFieldRow(
        form.NewStringField("first", "First"),
        form.NewStringField("last", "Last"),
    ),
    form.NewEmailField("email", "Email"),
)
```

Theme and error propagation work automatically inside rows.
