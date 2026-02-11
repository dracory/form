# Field Types

The library supports the following field types, each with a type-safe constructor:

| Constructor | Type | HTML Element |
|---|---|---|
| `NewStringField(name, label)` | string | `<input type="text">` |
| `NewEmailField(name, label)` | email | `<input type="email">` |
| `NewNumberField(name, label)` | number | `<input type="number">` |
| `NewPasswordField(name, label)` | password | `<input type="password">` |
| `NewHiddenField(name, value)` | hidden | `<input type="hidden">` |
| `NewDateField(name, label)` | date | `<input type="date">` |
| `NewDateTimeField(name, label)` | datetime | `<input type="datetime-local">` |
| `NewSelectField(name, label, options)` | select | `<select>` |
| `NewTextAreaField(name, label)` | textarea | `<textarea>` |
| `NewCheckboxField(name, label)` | checkbox | `<input type="checkbox">` |
| `NewRadioField(name, label, options)` | radio | `<input type="radio">` |
| `NewFileField(name, label)` | file | `<input type="file">` |
| `NewImageField(name, label)` | image | Image preview + URL input |
| `NewColorField(name, label)` | color | `<input type="color">` |
| `NewTelField(name, label)` | tel | `<input type="tel">` |
| `NewURLField(name, label)` | url | `<input type="url">` |
| `NewHtmlAreaField(name, label)` | htmlarea | Trumbowyg WYSIWYG editor |
| `NewRawField(value)` | raw | Raw HTML output |

All constructors return `*Field`, which supports chaining with `With*` methods.
