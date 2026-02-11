# HTMX Integration

## Simple HTMX Attributes

```golang
f := form.New().
    WithHxPost("/submit").
    WithHxTarget("#result").
    WithHxSwap("innerHTML").
    WithFields(...)
```

## HTMX Config Struct

For advanced HTMX usage, use the `HTMXConfig` struct:

```golang
f := form.New().
    WithHTMX(form.HTMXConfig{
        Post:        "/submit",
        Target:      "#result",
        Swap:        "innerHTML",
        Trigger:     "submit",
        Indicator:   "#spinner",
        Confirm:     "Are you sure?",
        Sync:        "closest form:abort",
        Validate:    true,
        DisabledElt: "this",
        Encoding:    "multipart/form-data",
        PushURL:     "/success",
    }).
    WithFields(...)
```

## Config Fields

| Field | HTMX Attribute | Description |
|---|---|---|
| `Post` | `hx-post` | POST request URL |
| `Get` | `hx-get` | GET request URL |
| `Target` | `hx-target` | Target element selector |
| `Swap` | `hx-swap` | Swap method |
| `Trigger` | `hx-trigger` | Trigger event |
| `Indicator` | `hx-indicator` | Loading indicator selector |
| `Confirm` | `hx-confirm` | Confirmation dialog message |
| `Sync` | `hx-sync` | Sync strategy |
| `Validate` | `hx-validate` | Enable HTML5 validation |
| `DisabledElt` | `hx-disabled-elt` | Element to disable during request |
| `Encoding` | `hx-encoding` | Request encoding type |
| `PushURL` | `hx-push-url` | URL to push to browser history |
