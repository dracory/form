# Fluent API

## Form Builder Methods

All methods return `*Form` for chaining:

```golang
f := form.New().
    WithID("myForm").
    WithClass("my-class").
    WithMethod("POST").
    WithAction("/submit").
    WithFields(field1, field2, field3).
    WithFileManager("/file-manager").
    WithTheme(form.ThemeTailwind()).
    WithErrors(map[string]string{"name": "Name is required"}).
    WithHTMX(form.HTMXConfig{Post: "/submit", Target: "#result"})
```

| Method | Description |
|---|---|
| `New()` | Creates a new form with POST method |
| `WithID(id)` | Sets the form HTML id |
| `WithClass(class)` | Sets the form CSS class |
| `WithMethod(method)` | Sets the HTTP method |
| `WithAction(url)` | Sets the form action URL |
| `WithFields(fields...)` | Sets the form fields |
| `WithFileManager(url)` | Sets the file manager URL for image fields |
| `WithTheme(theme)` | Sets the CSS theme |
| `WithErrors(errors)` | Sets inline validation error messages |
| `WithHTMX(config)` | Sets HTMX attributes via config struct |
| `WithHxPost(url)` | Sets hx-post attribute |
| `WithHxTarget(target)` | Sets hx-target attribute |
| `WithHxSwap(swap)` | Sets hx-swap attribute |

## Field Builder Methods

All methods return `*Field` for chaining:

```golang
field := form.NewStringField("name", "Full Name").
    WithID("nameField").
    WithValue("John").
    WithPlaceholder("Enter your name").
    WithHelp("Your full legal name").
    WithRequired().
    WithReadonly().
    WithDisabled().
    WithAttr("autocomplete", "name").
    WithValidators(form.ValidatorMinLength(2))
```

| Method | Description |
|---|---|
| `WithID(id)` | Sets the field HTML id |
| `WithName(name)` | Sets the field name attribute |
| `WithLabel(label)` | Sets the field label text |
| `WithValue(value)` | Sets the field value |
| `WithType(fieldType)` | Sets the field type |
| `WithHelp(help)` | Sets help text below the input |
| `WithPlaceholder(text)` | Sets placeholder text |
| `WithRequired()` | Marks the field as required |
| `WithReadonly()` | Marks the field as readonly |
| `WithDisabled()` | Marks the field as disabled |
| `WithInvisible()` | Hides the field via CSS |
| `WithMultiple()` | Enables multi-select |
| `WithOptions(options...)` | Sets static options (select, radio) |
| `WithOptionsF(fn)` | Sets a dynamic options provider function |
| `WithCustomInput(tag)` | Sets a custom input element (blockeditor) |
| `WithAttr(key, value)` | Sets a single custom HTML attribute |
| `WithAttrs(attrs)` | Sets multiple custom HTML attributes |
| `WithValidators(validators...)` | Sets validators |
| `WithTableOptions(opts)` | Sets table options for table fields |
