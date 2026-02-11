package form

import (
	"strings"
	"testing"
)

func TestWithHTMXPost(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:   "/submit",
		Target: "#result",
		Swap:   "innerHTML",
	})
	html := f.Build().ToHTML()

	expecteds := []string{
		`hx-post="/submit"`,
		`hx-target="#result"`,
		`hx-swap="innerHTML"`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}

func TestWithHTMXGet(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Get:    "/search",
		Target: "#results",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-get="/search"`) {
		t.Fatal("Expected hx-get, got:", html)
	}
}

func TestWithHTMXTrigger(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:    "/submit",
		Trigger: "submit",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-trigger="submit"`) {
		t.Fatal("Expected hx-trigger, got:", html)
	}
}

func TestWithHTMXIndicator(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:      "/submit",
		Indicator: "#spinner",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-indicator="#spinner"`) {
		t.Fatal("Expected hx-indicator, got:", html)
	}
}

func TestWithHTMXConfirm(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:    "/delete",
		Confirm: "Are you sure?",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-confirm="Are you sure?"`) {
		t.Fatal("Expected hx-confirm, got:", html)
	}
}

func TestWithHTMXSync(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post: "/submit",
		Sync: "closest form:abort",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-sync="closest form:abort"`) {
		t.Fatal("Expected hx-sync, got:", html)
	}
}

func TestWithHTMXValidate(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:     "/submit",
		Validate: true,
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-validate="true"`) {
		t.Fatal("Expected hx-validate, got:", html)
	}
}

func TestWithHTMXDisabledElt(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:        "/submit",
		DisabledElt: "this",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-disabled-elt="this"`) {
		t.Fatal("Expected hx-disabled-elt, got:", html)
	}
}

func TestWithHTMXEncoding(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:     "/upload",
		Encoding: "multipart/form-data",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-encoding="multipart/form-data"`) {
		t.Fatal("Expected hx-encoding, got:", html)
	}
}

func TestWithHTMXPushURL(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
		Post:    "/submit",
		PushURL: "/success",
	})
	html := f.Build().ToHTML()

	if !strings.Contains(html, `hx-push-url="/success"`) {
		t.Fatal("Expected hx-push-url, got:", html)
	}
}

func TestWithHTMXFullConfig(t *testing.T) {
	f := New().WithHTMX(HTMXConfig{
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
		PushURL:     "/done",
	})
	html := f.Build().ToHTML()

	expecteds := []string{
		`hx-post="/submit"`,
		`hx-target="#result"`,
		`hx-swap="innerHTML"`,
		`hx-trigger="submit"`,
		`hx-indicator="#spinner"`,
		`hx-confirm="Are you sure?"`,
		`hx-sync="closest form:abort"`,
		`hx-validate="true"`,
		`hx-disabled-elt="this"`,
		`hx-encoding="multipart/form-data"`,
		`hx-push-url="/done"`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected:", exp, "got:", html)
		}
	}
}

func TestWithHTMXDoesNotBreakLegacy(t *testing.T) {
	f := New().
		WithHxPost("/legacy").
		WithHxTarget("#old").
		WithHxSwap("outerHTML")

	html := f.Build().ToHTML()

	expecteds := []string{
		`hx-post="/legacy"`,
		`hx-target="#old"`,
		`hx-swap="outerHTML"`,
	}
	for _, exp := range expecteds {
		if !strings.Contains(html, exp) {
			t.Fatal("Expected legacy attr:", exp, "got:", html)
		}
	}
}
