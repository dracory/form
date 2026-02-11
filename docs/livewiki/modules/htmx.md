---
path: modules/htmx.md
page-type: module
summary: Documentation for HTMX integration including HTMXConfig and simple hx-* attribute helpers.
tags: [module, htmx, ajax, integration]
created: 2025-02-11
updated: 2025-02-11
version: 1.0.0
---

# Module: HTMX

## Purpose

The HTMX module provides two ways to add HTMX attributes to forms: simple `WithHx*` helper methods for the three most common attributes, and a structured `HTMXConfig` for full control over all supported HTMX attributes.

## Key Types

### HTMXConfig

```go
type HTMXConfig struct {
    Post        string // hx-post URL
    Get         string // hx-get URL
    Target      string // hx-target selector
    Swap        string // hx-swap method
    Trigger     string // hx-trigger event
    Indicator   string // hx-indicator selector
    Confirm     string // hx-confirm message
    Sync        string // hx-sync strategy
    Validate    bool   // hx-validate="true"
    DisabledElt string // hx-disabled-elt selector
    Encoding    string // hx-encoding (e.g. "multipart/form-data")
    PushURL     string // hx-push-url
}
```

## Two Approaches

### Simple Attributes

For the most common case (POST to a URL, target a container, swap content):

```go
f := form.New().
    WithHxPost("/api/submit").
    WithHxTarget("#result").
    WithHxSwap("innerHTML")
```

Renders:
```html
<form method="POST" hx-post="/api/submit" hx-target="#result" hx-swap="innerHTML">
```

### Full HTMXConfig

For advanced HTMX usage with additional attributes:

```go
f := form.New().WithHTMX(form.HTMXConfig{
    Post:        "/api/submit",
    Target:      "#result",
    Swap:        "innerHTML",
    Trigger:     "submit",
    Indicator:   "#loading-spinner",
    Confirm:     "Are you sure you want to submit?",
    Sync:        "closest form:abort",
    Validate:    true,
    DisabledElt: "find button",
    Encoding:    "multipart/form-data",
    PushURL:     "/submitted",
})
```

### Precedence

When both simple attributes and `HTMXConfig` are set, `HTMXConfig` values **override** the simple ones (for non-empty fields):

```go
f := form.New().
    WithHxPost("/old-url").          // will be overridden
    WithHTMX(form.HTMXConfig{
        Post: "/new-url",            // this wins
    })
```

## Attribute Mapping

| HTMXConfig Field | HTML Attribute | Example Value |
|-----------------|----------------|---------------|
| `Post` | `hx-post` | `"/api/submit"` |
| `Get` | `hx-get` | `"/api/search"` |
| `Target` | `hx-target` | `"#result"` |
| `Swap` | `hx-swap` | `"innerHTML"`, `"outerHTML"`, `"beforeend"` |
| `Trigger` | `hx-trigger` | `"submit"`, `"click"`, `"keyup changed delay:500ms"` |
| `Indicator` | `hx-indicator` | `"#spinner"` |
| `Confirm` | `hx-confirm` | `"Are you sure?"` |
| `Sync` | `hx-sync` | `"closest form:abort"` |
| `Validate` | `hx-validate` | `"true"` (only when `true`) |
| `DisabledElt` | `hx-disabled-elt` | `"this"`, `"find button"` |
| `Encoding` | `hx-encoding` | `"multipart/form-data"` |
| `PushURL` | `hx-push-url` | `"/new-url"`, `"true"` |

## HTMX in Repeater Fields

The repeater field uses HTMX internally for add/remove/reorder operations:

```go
form.NewRepeater(form.RepeaterOptions{
    RepeaterAddUrl:      "/repeater/add",      // hx-post on Add button
    RepeaterRemoveUrl:   "/repeater/remove",   // hx-post on Remove button
    RepeaterMoveUpUrl:   "/repeater/up",       // hx-post on Move Up button
    RepeaterMoveDownUrl: "/repeater/down",     // hx-post on Move Down button
})
```

These buttons use `hx-include` to send the form data and `hx-target` to replace the form content.

## Files

| File | Contents |
|------|----------|
| `htmx.go` | `HTMXConfig` struct |
| `form_fluent.go` | `WithHxPost()`, `WithHxTarget()`, `WithHxSwap()`, `WithHTMX()` |
| `form.go` | HTMX attribute application in `Build()` |

## See Also

- [Configuration](../configuration.md)
- [Module: Form](form.md)
- [Module: Repeater](repeater.md)
- [API Reference](../api_reference.md)
