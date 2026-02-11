---
path: llm-context.md
page-type: overview
summary: Complete codebase summary optimized for LLM consumption.
tags: [llm, context, summary]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# LLM Context: Dracory Form

## Project Summary

`github.com/dracory/form` is a Go library for programmatically building HTML forms. It generates HTML via the `dracory/hb` HTML builder library (not Go templates). The library provides 18 field types, a fluent builder API, 13 built-in validators with inline error display, theming (Bootstrap 5 default, Tailwind CSS), HTMX integration, grid layouts via field rows, and dynamic repeater fields. Everything lives in a single `package form` with no sub-packages.

The library supports two construction patterns: an options-struct API (`NewForm(FormOptions{...})`) for backward compatibility and a fluent builder API (`New().WithID(...).WithFields(...)`) for ergonomic chaining. Forms render to `*hb.Tag` via `Build()`, which is then converted to an HTML string via `.ToHTML()`.

## Key Technologies

- **Go 1.25+**
- **github.com/dracory/hb** — HTML builder (tag construction, not templates)
- **github.com/dracory/uid** — Unique ID generation
- **github.com/samber/lo** — Functional utilities (lo.If, lo.Map, lo.Find)
- **github.com/spf13/cast** — Type casting (repeater index conversion)

## Directory Structure

```
form/                          # Single package, flat layout
├── consts.go                  # 20 field type constants (FORM_FIELD_TYPE_*)
├── field.go                   # Field struct (681 lines), rendering logic, BuildFormGroup
├── field_constructors.go      # 18 type-safe constructors (NewStringField, etc.)
├── field_fluent.go            # 19 fluent With* methods for Field
├── field_interface.go         # FieldInterface (17 methods)
├── field_option.go            # FieldOption{Key, Value} struct
├── field_repeater.go          # fieldRepeater - dynamic add/remove groups via HTMX
├── field_row.go               # fieldRow - multi-column grid layouts
├── form.go                    # Form struct, Build(), internal interfaces
├── form_fluent.go             # 12 fluent With* methods for Form + New()
├── htmx.go                    # HTMXConfig struct (12 fields)
├── new_field.go               # NewField(FieldOptions) constructor
├── new_form.go                # NewForm(FormOptions) constructor
├── new_repeater.go            # NewRepeater(RepeaterOptions) constructor
├── test_helpers.go            # 6 test assertion helpers
├── theme.go                   # Theme struct (17 CSS class fields), Bootstrap5, Tailwind
├── validation.go              # 13 validators + Form.Validate()
├── Taskfile.yml               # Task runner (test, cover, lint, security)
└── *_test.go                  # Comprehensive test files
```

## Core Concepts

1. **Form** — Container struct holding fields, theme, HTMX config, errors. `Build()` renders all fields into an `<form>` tag.

2. **Field** — Exported struct implementing `FieldInterface`. Holds type, name, label, value, validators, options, and rendering config. `BuildFormGroup()` renders label + input + help + error.

3. **FieldInterface** — Public interface (17 methods) that all field types implement: `Field`, `fieldRepeater`, `fieldRow`.

4. **Internal interfaces** — `formAware` (repeater needs parent form), `themeable` (accepts theme), `errorAware` (accepts error message), `rowErrorAware` (distributes errors to children). Checked via type assertion in `Build()`.

5. **Theme** — Strategy struct with 17 CSS class fields. `ThemeBootstrap5()` is the default. `ThemeTailwind()` is an alternative. Custom themes are plain struct instances.

6. **Validator** — Function type `func(fieldName, value string) *ValidationError`. 13 built-in validators. `Form.Validate()` runs validators and stores errors for inline display.

7. **HTMXConfig** — Struct with 12 fields mapping to `hx-*` attributes. Applied during `Build()`.

8. **FieldRow** — Layout pseudo-field rendering children in a Bootstrap/CSS grid row. Not a real input field.

9. **Repeater** — Dynamic field group with HTMX-powered add/remove/reorder. Clones child fields for each value set.

## Common Patterns

- **Builder pattern** — All `With*` methods return the receiver pointer for chaining
- **Interface satisfaction checks** — `var _ FieldInterface = (*Field)(nil)` at package level
- **Switch-based rendering** — `fieldInput()` uses a type switch to create the appropriate HTML input
- **Lazy options** — `OptionsF func() []FieldOption` evaluated during rendering
- **Clone pattern** — `clone() FieldInterface` for repeater field duplication
- **Error injection** — `Build()` injects errors into fields via internal interfaces before rendering

## Important Files

| File | Key Contents |
|------|-------------|
| `field.go` | `Field` struct, `BuildFormGroup()`, `fieldInput()` switch, type check methods, `TrumbowygScript()` |
| `form.go` | `Form` struct, `Build()` rendering pipeline, internal interfaces (`formAware`, `themeable`, `errorAware`, `rowErrorAware`) |
| `form_fluent.go` | `New()` constructor, all form `With*` methods |
| `field_fluent.go` | All field `With*` methods (19 methods) |
| `field_constructors.go` | 18 type-safe constructors (`NewStringField`, `NewEmailField`, etc.) |
| `validation.go` | `Validator` type, 13 built-in validators, `Form.Validate()` |
| `theme.go` | `Theme` struct, `ThemeBootstrap5()`, `ThemeTailwind()`, `defaultTheme` |
| `field_interface.go` | `FieldInterface` definition (17 methods) |
| `field_repeater.go` | `fieldRepeater` struct, HTMX-powered add/remove/reorder |
| `field_row.go` | `fieldRow` struct, `NewFieldRow()`, `NewFieldRowWithColumns()` |
| `htmx.go` | `HTMXConfig` struct (12 HTMX attribute fields) |
| `test_helpers.go` | 6 assertion helpers for testing forms and validation |
| `consts.go` | 20 field type constants |
