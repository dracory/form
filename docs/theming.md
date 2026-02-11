# Theming

The library ships with Bootstrap 5 (default) and Tailwind CSS themes. You can also create custom themes.

## Built-in Themes

```golang
// Bootstrap 5 (default, no configuration needed)
f := form.New().WithFields(...)

// Tailwind CSS
f := form.New().
    WithTheme(form.ThemeTailwind()).
    WithFields(...)
```

## Custom Theme

```golang
myTheme := &form.Theme{
    FormGroupClass:     "my-group",
    LabelClass:         "my-label",
    InputClass:         "my-input",
    SelectClass:        "my-select",
    TextAreaClass:      "my-textarea",
    CheckboxWrapClass:  "my-check",
    CheckboxInputClass: "my-check-input",
    RadioWrapClass:     "my-radio",
    RadioInputClass:    "my-radio-input",
    RadioLabelClass:    "my-radio-label",
    FileInputClass:     "my-file",
    HelpClass:          "my-help",
    RequiredClass:      "my-required",
    RequiredMarker:     "*",
    TableClass:         "my-table",
    ErrorClass:         "my-error",
    ErrorInputClass:    "my-input-error",
}

f := form.New().WithTheme(myTheme).WithFields(...)
```

## Theme Fields

| Field | Bootstrap 5 Default | Description |
|---|---|---|
| `FormGroupClass` | `form-group mb-3` | Wrapper div for each field |
| `LabelClass` | `form-label` | Label element |
| `InputClass` | `form-control` | Text, email, number, etc. inputs |
| `SelectClass` | `form-select` | Select dropdowns |
| `TextAreaClass` | `form-control` | Textarea elements |
| `CheckboxWrapClass` | `form-check` | Checkbox wrapper div |
| `CheckboxInputClass` | `form-check-input` | Checkbox input |
| `RadioWrapClass` | `form-check` | Radio button wrapper div |
| `RadioInputClass` | `form-check-input` | Radio button input |
| `RadioLabelClass` | `form-check-label` | Radio button label |
| `FileInputClass` | `form-control` | File input |
| `HelpClass` | `text-info` | Help text paragraph |
| `RequiredClass` | `text-danger ms-1` | Required marker (asterisk) |
| `RequiredMarker` | `*` | Required indicator text |
| `TableClass` | `table table-striped table-hover mb-0` | Table element |
| `ErrorClass` | `invalid-feedback` | Error message div |
| `ErrorInputClass` | `is-invalid` | Added to invalid inputs |
