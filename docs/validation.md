# Validation

## Built-in Validators

```golang
field := form.NewStringField("username", "Username").
    WithRequired().
    WithValidators(
        form.ValidatorMinLength(3),
        form.ValidatorMaxLength(50),
        form.ValidatorAlphaNumeric(),
    )
```

| Validator | Description |
|---|---|
| `ValidatorRequired()` | Value must not be empty |
| `ValidatorMinLength(n)` | Minimum character length |
| `ValidatorMaxLength(n)` | Maximum character length |
| `ValidatorMin(n)` | Minimum numeric value |
| `ValidatorMax(n)` | Maximum numeric value |
| `ValidatorPattern(regex, msg)` | Must match regex pattern |
| `ValidatorEmail()` | Must be a valid email address |
| `ValidatorURL()` | Must be a valid HTTP/HTTPS URL |
| `ValidatorIP()` | Must be a valid IPv4 address |
| `ValidatorUUID()` | Must be a valid UUID |
| `ValidatorAlphaNumeric()` | Must contain only letters and numbers |
| `ValidatorOneOf(values...)` | Must be one of the allowed values |
| `ValidatorCustom(fn)` | Custom validation function |

## Using Validation

```golang
f := form.New().WithFields(
    form.NewStringField("name", "Name").WithRequired(),
    form.NewEmailField("email", "Email").WithValidators(form.ValidatorEmail()),
)

// Validate returns errors and stores them for inline display
errs := f.Validate(map[string]string{
    "name":  "",
    "email": "invalid",
})

if len(errs) > 0 {
    // Errors are automatically shown inline when Build() is called
    html := f.Build().ToHTML()
}
```

## Manual Error Display

You can also set errors manually without using `Validate()`:

```golang
f := form.New().
    WithFields(form.NewStringField("name", "Name")).
    WithErrors(map[string]string{
        "name": "This name is already taken",
    })

html := f.Build().ToHTML()
// The "name" field will have the error class and message rendered inline
```

## Custom Validator

```golang
form.ValidatorCustom(func(value string) string {
    if value == "" || value[0] == '#' {
        return ""
    }
    return "must start with #"
})
```

The function receives the field value and returns an error message string.
Return an empty string if the value is valid.
