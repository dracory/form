# Test Helpers

The library provides test helper functions for use in your tests.

## Usage

```golang
import "github.com/dracory/form"

func TestMyForm(t *testing.T) {
    f := form.New().WithID("myForm").WithFields(
        form.NewStringField("name", "Name").WithRequired(),
        form.NewEmailField("email", "Email").WithValidators(form.ValidatorEmail()),
    )

    // Assert rendered HTML content
    form.AssertFormContains(t, f, `id="myForm"`)
    form.AssertFormNotContains(t, f, `is-invalid`)

    // Assert field HTML content
    field := form.NewEmailField("email", "Email")
    form.AssertFieldContains(t, field, `type="email"`)
    form.AssertFieldNotContains(t, field, `readonly`)

    // Assert validation outcomes
    form.AssertValidationPasses(t, f, map[string]string{
        "name": "John", "email": "john@example.com",
    })
    form.AssertValidationFails(t, f, map[string]string{
        "name": "", "email": "invalid",
    })
    form.AssertValidationFailsOn(t, f, map[string]string{
        "name": "", "email": "john@example.com",
    }, "name")
    form.AssertValidationErrorCount(t, f, map[string]string{
        "name": "", "email": "invalid",
    }, 2)
}
```

## Available Helpers

| Helper | Description |
|---|---|
| `AssertFormContains(t, form, expected)` | Form HTML contains string |
| `AssertFormNotContains(t, form, unexpected)` | Form HTML does not contain string |
| `AssertFieldContains(t, field, expected)` | Field HTML contains string |
| `AssertFieldNotContains(t, field, unexpected)` | Field HTML does not contain string |
| `AssertValidationPasses(t, form, values)` | Validation produces no errors |
| `AssertValidationFails(t, form, values)` | Validation produces errors |
| `AssertValidationFailsOn(t, form, values, fieldName)` | Validation fails on specific field |
| `AssertValidationErrorCount(t, form, values, count)` | Validation produces exactly N errors |

All helpers call `t.Helper()` so test failures point to the caller's line.
