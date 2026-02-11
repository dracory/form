---
path: configuration.md
page-type: reference
summary: Configuration options, themes, and environment setup for the Dracory Form library.
tags: [configuration, themes, options, setup]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Configuration

## Form Configuration

Forms are configured either via `FormOptions` or fluent `With*` methods. There are no environment variables or config files — everything is done in code.

### FormOptions Fields

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `ID` | `string` | `""` | HTML `id` attribute |
| `ClassName` | `string` | `""` | CSS class |
| `Method` | `string` | `"POST"` | HTTP method |
| `ActionURL` | `string` | `""` | Form action URL |
| `Fields` | `[]FieldInterface` | `nil` | Form fields |
| `FileManagerURL` | `string` | `""` | URL for image field file browser |
| `HxPost` | `string` | `""` | HTMX post URL |
| `HxTarget` | `string` | `""` | HTMX target selector |
| `HxSwap` | `string` | `""` | HTMX swap method |

### FieldOptions Fields

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `ID` | `string` | Auto-generated | HTML `id` (auto: `"id_" + uid.HumanUid()`) |
| `Type` | `string` | `""` | Field type constant |
| `Name` | `string` | `""` | Input `name` attribute |
| `Label` | `string` | `""` | Label text (falls back to Name) |
| `Help` | `string` | `""` | Help text below the input |
| `Value` | `string` | `""` | Field value |
| `Placeholder` | `string` | `""` | Placeholder text |
| `Required` | `bool` | `false` | Whether field is required |
| `Readonly` | `bool` | `false` | Whether field is readonly |
| `Disabled` | `bool` | `false` | Whether field is disabled |
| `Invisible` | `bool` | `false` | Whether field is hidden via CSS |
| `Multiple` | `bool` | `false` | Enable multi-select |
| `Options` | `[]FieldOption` | `nil` | Static options for select/radio |
| `OptionsF` | `func() []FieldOption` | `nil` | Dynamic options provider |
| `CustomInput` | `hb.TagInterface` | `nil` | Custom input element |
| `Attrs` | `map[string]string` | `nil` | Custom HTML attributes |
| `Validators` | `[]Validator` | `nil` | Validation functions |
| `TableOptions` | `TableOptions` | zero | Table field configuration |

## Theme Configuration

### Built-in Themes

| Theme | Function | Description |
|-------|----------|-------------|
| Bootstrap 5 | `ThemeBootstrap5()` | Default theme. Uses `form-control`, `form-group`, etc. |
| Tailwind CSS | `ThemeTailwind()` | Uses Tailwind utility classes |

### Custom Theme

Create a custom theme by instantiating `Theme` directly:

```go
myTheme := &form.Theme{
    FormGroupClass:     "my-form-group",
    LabelClass:         "my-label",
    InputClass:         "my-input",
    SelectClass:        "my-select",
    TextAreaClass:      "my-textarea",
    CheckboxWrapClass:  "my-checkbox-wrap",
    CheckboxInputClass: "my-checkbox",
    RadioWrapClass:     "my-radio-wrap",
    RadioInputClass:    "my-radio",
    RadioLabelClass:    "my-radio-label",
    FileInputClass:     "my-file-input",
    HelpClass:          "my-help",
    RequiredClass:      "my-required",
    RequiredMarker:     "*",
    TableClass:         "my-table",
    ErrorClass:         "my-error",
    ErrorInputClass:    "my-error-input",
}

f := form.New().WithTheme(myTheme)
```

### Theme Fields Reference

| Field | Bootstrap 5 Default | Purpose |
|-------|-------------------|---------|
| `FormGroupClass` | `"form-group mb-3"` | Wrapper div for each field |
| `LabelClass` | `"form-label"` | Label element |
| `InputClass` | `"form-control"` | Text/number/date/etc. inputs |
| `SelectClass` | `"form-select"` | Select dropdowns |
| `TextAreaClass` | `"form-control"` | Textarea elements |
| `CheckboxWrapClass` | `"form-check"` | Checkbox wrapper div |
| `CheckboxInputClass` | `"form-check-input"` | Checkbox input |
| `RadioWrapClass` | `"form-check"` | Radio button wrapper |
| `RadioInputClass` | `"form-check-input"` | Radio button input |
| `RadioLabelClass` | `"form-check-label"` | Radio button label |
| `FileInputClass` | `"form-control"` | File input |
| `HelpClass` | `"text-info"` | Help text paragraph |
| `RequiredClass` | `"text-danger ms-1"` | Required marker (`<sup>`) |
| `RequiredMarker` | `"*"` | Required indicator text |
| `TableClass` | `"table table-striped table-hover mb-0"` | Table element |
| `ErrorClass` | `"invalid-feedback"` | Error message div |
| `ErrorInputClass` | `"is-invalid"` | Class added to invalid inputs |

## HTMX Configuration

### Simple Attributes

```go
form.New().
    WithHxPost("/submit").
    WithHxTarget("#result").
    WithHxSwap("innerHTML")
```

### Full HTMXConfig

```go
form.New().WithHTMX(form.HTMXConfig{
    Post:        "/submit",       // hx-post
    Get:         "/search",       // hx-get
    Target:      "#result",       // hx-target
    Swap:        "innerHTML",     // hx-swap
    Trigger:     "submit",        // hx-trigger
    Indicator:   "#spinner",      // hx-indicator
    Confirm:     "Are you sure?", // hx-confirm
    Sync:        "closest form:abort", // hx-sync
    Validate:    true,            // hx-validate="true"
    DisabledElt: "this",          // hx-disabled-elt
    Encoding:    "multipart/form-data", // hx-encoding
    PushURL:     "/new-url",      // hx-push-url
})
```

> **Note:** `HTMXConfig` fields override the simple `WithHxPost`/`WithHxTarget`/`WithHxSwap` when both are set.

## Trumbowyg WYSIWYG Configuration

The `htmlarea` field type uses Trumbowyg. Custom config can be passed via options:

```go
form.NewHtmlAreaField("content", "Content").
    WithOptions(form.FieldOption{
        Key:   "config",
        Value: `{"btns": [["bold", "italic"]], "autogrow": true}`,
    })
```

If no config option is provided, a default config is used with common formatting buttons.

## See Also

- [API Reference](api_reference.md)
- [Modules: Theme](modules/theme.md)
- [Modules: HTMX](modules/htmx.md)
- [Architecture](architecture.md)
