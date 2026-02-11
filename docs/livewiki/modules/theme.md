---
path: modules/theme.md
page-type: module
summary: Documentation for the Theme struct, built-in themes, and custom theme creation.
tags: [module, theme, css, bootstrap, tailwind, styling]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Module: Theme

## Purpose

The Theme module decouples form rendering from any specific CSS framework. A `Theme` struct holds CSS class names for every form element type. The library ships with Bootstrap 5 (default) and Tailwind CSS themes, and supports fully custom themes.

## Key Types

### Theme

```go
type Theme struct {
    FormGroupClass     string // Wrapper div for each field
    LabelClass         string // Label element
    InputClass         string // Text/number/date/etc. inputs
    SelectClass        string // Select dropdowns
    TextAreaClass      string // Textarea elements
    CheckboxWrapClass  string // Checkbox wrapper div
    CheckboxInputClass string // Checkbox input
    RadioWrapClass     string // Radio button wrapper
    RadioInputClass    string // Radio button input
    RadioLabelClass    string // Radio button label
    FileInputClass     string // File input
    HelpClass          string // Help text paragraph
    RequiredClass      string // Required marker (sup element)
    RequiredMarker     string // Required indicator text (e.g., "*")
    TableClass         string // Table element
    ErrorClass         string // Error message div
    ErrorInputClass    string // Class added to invalid inputs
}
```

## Built-in Themes

### Bootstrap 5 (Default)

```go
theme := form.ThemeBootstrap5()
```

| Field | Value |
|-------|-------|
| `FormGroupClass` | `"form-group mb-3"` |
| `LabelClass` | `"form-label"` |
| `InputClass` | `"form-control"` |
| `SelectClass` | `"form-select"` |
| `TextAreaClass` | `"form-control"` |
| `CheckboxWrapClass` | `"form-check"` |
| `CheckboxInputClass` | `"form-check-input"` |
| `RadioWrapClass` | `"form-check"` |
| `RadioInputClass` | `"form-check-input"` |
| `RadioLabelClass` | `"form-check-label"` |
| `FileInputClass` | `"form-control"` |
| `HelpClass` | `"text-info"` |
| `RequiredClass` | `"text-danger ms-1"` |
| `RequiredMarker` | `"*"` |
| `TableClass` | `"table table-striped table-hover mb-0"` |
| `ErrorClass` | `"invalid-feedback"` |
| `ErrorInputClass` | `"is-invalid"` |

### Tailwind CSS

```go
theme := form.ThemeTailwind()
```

| Field | Value |
|-------|-------|
| `FormGroupClass` | `"mb-4"` |
| `LabelClass` | `"block text-sm font-medium text-gray-700 mb-1"` |
| `InputClass` | `"block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"` |
| `SelectClass` | Same as InputClass |
| `TextAreaClass` | Same as InputClass |
| `CheckboxWrapClass` | `"flex items-center"` |
| `CheckboxInputClass` | `"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"` |
| `RadioWrapClass` | `"flex items-center"` |
| `RadioInputClass` | `"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-500"` |
| `RadioLabelClass` | `"ml-2 block text-sm text-gray-900"` |
| `FileInputClass` | Tailwind file input utilities |
| `HelpClass` | `"mt-1 text-sm text-gray-500"` |
| `RequiredClass` | `"text-red-500 ml-1"` |
| `ErrorClass` | `"mt-1 text-sm text-red-600"` |
| `ErrorInputClass` | `"border-red-500"` |

## Theme Resolution

```
field.getTheme() → field.theme (if injected by form) → defaultTheme (Bootstrap 5)
```

- During `Form.Build()`, the form's theme is injected into each `themeable` field
- If no theme is set on the form, `defaultTheme` (Bootstrap 5) is used
- Fields rendered outside of `Build()` (e.g., `BuildFormGroup()` directly) use `defaultTheme`

## Custom Theme Example

```go
myTheme := &form.Theme{
    FormGroupClass:     "space-y-2",
    LabelClass:         "font-bold text-lg",
    InputClass:         "input input-bordered w-full",
    SelectClass:        "select select-bordered w-full",
    TextAreaClass:      "textarea textarea-bordered w-full",
    CheckboxWrapClass:  "form-control",
    CheckboxInputClass: "checkbox",
    RadioWrapClass:     "form-control",
    RadioInputClass:    "radio",
    RadioLabelClass:    "label",
    FileInputClass:     "file-input file-input-bordered w-full",
    HelpClass:          "text-sm opacity-70",
    RequiredClass:      "text-error ml-1",
    RequiredMarker:     "*",
    TableClass:         "table table-zebra",
    ErrorClass:         "text-sm text-error mt-1",
    ErrorInputClass:    "input-error",
}

f := form.New().WithTheme(myTheme).WithFields(...)
```

## Files

| File | Contents |
|------|----------|
| `theme.go` | `Theme` struct, `ThemeBootstrap5()`, `ThemeTailwind()`, `defaultTheme` |

## See Also

- [Configuration](../configuration.md)
- [Module: Form](form.md)
- [Module: Field](field.md)
- [API Reference](../api_reference.md)
