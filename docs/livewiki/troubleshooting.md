---
path: troubleshooting.md
page-type: reference
summary: Common issues and solutions when using the Dracory Form library.
tags: [troubleshooting, errors, debugging, faq]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Troubleshooting

## Common Issues

### Form renders empty HTML

**Symptom:** `form.Build().ToHTML()` returns `<form method="POST"></form>`

**Cause:** No fields were added to the form.

**Solution:** Ensure you pass fields via `WithFields()` or `AddField()`:

```go
f := form.New().WithFields(
    form.NewStringField("name", "Name"),
)
```

### Field ID is auto-generated on every render

**Symptom:** Field IDs change between renders (e.g., `id_abc123`, `id_def456`).

**Cause:** If no `ID` is set on a field, `BuildFormGroup()` generates one using `uid.HumanUid()`.

**Solution:** Set an explicit ID:

```go
form.NewStringField("name", "Name").WithID("field_name")
```

### Select field readonly doesn't work

**Symptom:** A readonly select field can still be changed.

**Cause:** HTML `<select>` elements don't support the `readonly` attribute natively.

**Solution:** The library handles this automatically — readonly selects are rendered as `disabled` with a hidden input to preserve the value. The select's `name` is changed to `NAME_Readonly` and a hidden field with the original name carries the value.

### Validation errors not showing inline

**Symptom:** `Validate()` returns errors but they don't appear next to fields.

**Cause:** You must call `Build()` **after** `Validate()`, or the theme's `ErrorClass`/`ErrorInputClass` are empty.

**Solution:**

```go
errors := f.Validate(values)
// errors are now stored on the form
html := f.Build().ToHTML() // errors render inline
```

Or set errors manually:

```go
f.WithErrors(map[string]string{
    "email": "Invalid email address",
})
html := f.Build().ToHTML()
```

### Repeater shows "Form Error" messages

**Symptom:** Repeater renders an alert div with an error message.

**Possible causes and solutions:**

| Error Message | Cause | Solution |
|---------------|-------|----------|
| `Repeater has no name` | `Name` not set in `RepeaterOptions` | Set `Name` field |
| `Repeater X has no repeaterAddUrl` | `RepeaterAddUrl` not set | Set `RepeaterAddUrl` |
| `Repeater X has no repeaterRemoveUrl` | `RepeaterRemoveUrl` not set | Set `RepeaterRemoveUrl` |

### Block editor shows "CustomInput is nil"

**Symptom:** Block editor field renders an alert and a fallback textarea.

**Cause:** The `blockeditor` field type requires a `CustomInput` to be set.

**Solution:**

```go
form.NewField(form.FieldOptions{
    Type:        form.FORM_FIELD_TYPE_BLOCKEDITOR,
    Name:        "content",
    Label:       "Content",
    CustomInput: myCustomEditorTag, // must be an hb.TagInterface
})
```

### Checkbox value not preserved

**Symptom:** Checkbox always submits `"1"` regardless of the set value.

**Cause:** If `Value` is empty, the checkbox defaults to `"1"`.

**Solution:** Set an explicit value if needed:

```go
form.NewCheckboxField("agree", "I Agree").WithValue("yes")
```

Checked state is determined by value being `"1"`, `"true"`, `"on"`, or `"yes"`.

### Trumbowyg WYSIWYG not initializing

**Symptom:** HTML area field shows a plain textarea.

**Cause:** Trumbowyg JS/CSS are not loaded on the page, or jQuery is missing.

**Solution:** Ensure your page includes:
1. jQuery
2. Trumbowyg CSS
3. Trumbowyg JS

The library generates the initialization script automatically via `TrumbowygScript()`.

### Theme not applied to all fields

**Symptom:** Some fields use Bootstrap classes while others use your custom theme.

**Cause:** Theme is injected during `Build()`. Fields rendered outside of `Build()` (e.g., calling `BuildFormGroup()` directly) use the default theme.

**Solution:** Either:
- Always render via `form.Build()`, or
- Set the theme on individual fields (not directly supported — use the form-level theme)

### OptionsF not being called

**Symptom:** Select field with `OptionsF` shows no dynamic options.

**Cause:** `OptionsF` is called during rendering. If you're inspecting the field before rendering, options won't be populated.

**Solution:** `OptionsF` is evaluated lazily during `fieldSelect()`. Ensure the function is set before `Build()` is called.

## Debugging Tips

1. **Inspect rendered HTML:** Call `.ToHTML()` and print the output to see exactly what's generated.
2. **Check field types:** Use `field.IsString()`, `field.IsSelect()`, etc. to verify field types.
3. **Validate separately:** Call `Validate()` and inspect the returned `[]ValidationError` slice.
4. **Use test helpers:** `AssertFormContains()` and `AssertFieldContains()` help pinpoint rendering issues.

## See Also

- [Getting Started](getting_started.md)
- [Configuration](configuration.md)
- [Development](development.md)
- [API Reference](api_reference.md)
