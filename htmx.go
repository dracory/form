package form

// HTMXConfig provides a structured way to configure HTMX attributes on a form.
type HTMXConfig struct {
	Post        string // hx-post URL
	Get         string // hx-get URL
	Target      string // hx-target selector
	Swap        string // hx-swap method
	Trigger     string // hx-trigger event
	Indicator   string // hx-indicator selector
	Confirm     string // hx-confirm message
	Sync        string // hx-sync strategy
	Validate    bool   // hx-validate
	DisabledElt string // hx-disabled-elt selector
	Encoding    string // hx-encoding (e.g. "multipart/form-data")
	PushURL     string // hx-push-url
}
