package form

import (
	"strings"
	"testing"
)

func TestFieldBlockEditor(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_BLOCKEDITOR,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label>`,
		`<textarea class="form-control" id="ID" name="NAME">VALUE</textarea>`,
		`<div class="alert alert-danger">CustomInput is nil</div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldDate(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_DATE,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="date" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldDateTime(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_DATETIME,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="datetime-local" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldHidden(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_HIDDEN,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><input class="form-control" id="ID" name="NAME" type="hidden" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldHtmlArea(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_HTMLAREA,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`function initWysiwyg(textareaID, config) {`,
		`<textarea class="form-control" id="ID" name="NAME">VALUE</textarea>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldImage(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_IMAGE,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><div class="row g-3" style="border: 1px solid silver;border-radius: 10px; margin-top: 0px; margin-left: 0px;margin-right: 0px;"><div class="col-md-2"><img class="img-fluid rounded-start" src="VALUE" style="margin-bottom: 15px;width:100%;max-height:100px;" /></div><div class="col-md-10"><textarea  class="form-control" id="ID" name="NAME" style="height:70px;" type="text">VALUE</textarea><span>The URL can be base64 encoded image URL</span></div></div></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldNumber(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_NUMBER,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="number" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldPassword(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_PASSWORD,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="password" value="VALUE" /></div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldRaw(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE VALUE1 <br /> VALUE2 VALUE3",
		Type:  FORM_FIELD_TYPE_RAW,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`VALUE`,
		`VALUE1`,
		`<br />`,
		`VALUE2`,
		`VALUE3`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldSelect(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_SELECT,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><select class="form-select" id="ID" name="NAME"></select></div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldSelectWithOptions(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_SELECT,
		Options: []FieldOption{
			{
				Key:   "key1",
				Value: "value1",
			},
			{
				Key:   "key2",
				Value: "value2",
			},
			{
				Key:   "VALUE",
				Value: "value3",
			},
		},
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><select class="form-select" id="ID" name="NAME">`,
		`<option value="key1">value1</option>`,
		`<option value="key2">value2</option>`,
		`<option selected="selected" value="VALUE">value3</option>`,
		`</select></div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldString(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_STRING,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="text" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldTextArea(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_TEXTAREA,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><textarea class="form-control" id="ID" name="NAME">VALUE</textarea></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldReadonly(t *testing.T) {
	field := Field{
		ID:       "ID",
		Name:     "NAME",
		Value:    "VALUE",
		Type:     FORM_FIELD_TYPE_STRING,
		Readonly: true,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`readonly="readonly"`,
		`background: #efefef;`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldReadonlySelect(t *testing.T) {
	field := Field{
		ID:       "ID",
		Name:     "NAME",
		Value:    "VALUE",
		Type:     FORM_FIELD_TYPE_SELECT,
		Readonly: true,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`disabled="disabled"`,
		`name="NAME_Readonly"`,
		`type="hidden"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldDisabled(t *testing.T) {
	field := Field{
		ID:       "ID",
		Name:     "NAME",
		Value:    "VALUE",
		Type:     FORM_FIELD_TYPE_STRING,
		Disabled: true,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`disabled="disabled"`,
		`background: #efefef;`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldWithHelp(t *testing.T) {
	field := Field{
		ID:   "ID",
		Name: "NAME",
		Type: FORM_FIELD_TYPE_STRING,
		Help: "This is help text",
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`class="text-info"`,
		`This is help text`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldWithPlaceholder(t *testing.T) {
	field := Field{
		ID:          "ID",
		Name:        "NAME",
		Type:        FORM_FIELD_TYPE_STRING,
		Placeholder: "Enter value",
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `placeholder="Enter value"`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldInvisible(t *testing.T) {
	field := Field{
		ID:        "ID",
		Name:      "NAME",
		Type:      FORM_FIELD_TYPE_STRING,
		Invisible: true,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `display:none;`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldRequired(t *testing.T) {
	field := Field{
		ID:       "ID",
		Name:     "NAME",
		Type:     FORM_FIELD_TYPE_STRING,
		Required: true,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`<sup class="text-danger ms-1">*</sup>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldEmail(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "test@example.com",
		Type:  FORM_FIELD_TYPE_EMAIL,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `type="email"`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldTel(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "+1234567890",
		Type:  FORM_FIELD_TYPE_TEL,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `type="tel"`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldUrl(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "https://example.com",
		Type:  FORM_FIELD_TYPE_URL,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `type="url"`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldColor(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "#ff0000",
		Type:  FORM_FIELD_TYPE_COLOR,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `type="color"`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldCheckbox(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "1",
		Type:  FORM_FIELD_TYPE_CHECKBOX,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`type="checkbox"`,
		`class="form-check-input"`,
		`checked="checked"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldCheckboxUnchecked(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "0",
		Type:  FORM_FIELD_TYPE_CHECKBOX,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	if strings.Contains(html, `checked="checked"`) {
		t.Fatal(`Expected checkbox to not be checked, but was: `, html)
	}
}

func TestFieldRadio(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "opt2",
		Type:  FORM_FIELD_TYPE_RADIO,
		Options: []FieldOption{
			{Key: "opt1", Value: "Option 1"},
			{Key: "opt2", Value: "Option 2"},
		},
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`type="radio"`,
		`class="form-check-input"`,
		`class="form-check-label"`,
		`Option 1`,
		`Option 2`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldFile(t *testing.T) {
	field := Field{
		ID:   "ID",
		Name: "NAME",
		Type: FORM_FIELD_TYPE_FILE,
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`type="file"`,
		`class="form-control"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldWithAttrs(t *testing.T) {
	field := Field{
		ID:   "ID",
		Name: "NAME",
		Type: FORM_FIELD_TYPE_STRING,
		Attrs: map[string]string{
			"data-custom": "value",
			"aria-label":  "Custom label",
		},
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`data-custom="value"`,
		`aria-label="Custom label"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldSelectMultiple(t *testing.T) {
	field := Field{
		ID:       "ID",
		Name:     "NAME",
		Type:     FORM_FIELD_TYPE_SELECT,
		Multiple: true,
		Options: []FieldOption{
			{Key: "a", Value: "A"},
			{Key: "b", Value: "B"},
		},
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expected := `multiple="multiple"`
	if !strings.Contains(html, expected) {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldSelectWithOptionsF(t *testing.T) {
	field := Field{
		ID:   "ID",
		Name: "NAME",
		Type: FORM_FIELD_TYPE_SELECT,
		OptionsF: func() []FieldOption {
			return []FieldOption{
				{Key: "dynamic1", Value: "Dynamic 1"},
				{Key: "dynamic2", Value: "Dynamic 2"},
			}
		},
	}

	formGroup := field.BuildFormGroup("")
	html := formGroup.ToHTML()

	expecteds := []string{
		`<option value="dynamic1">Dynamic 1</option>`,
		`<option value="dynamic2">Dynamic 2</option>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldImageWithFileManagerURL(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "https://example.com/image.png",
		Type:  FORM_FIELD_TYPE_IMAGE,
	}

	formGroup := field.BuildFormGroup("https://example.com/filemanager")
	html := formGroup.ToHTML()

	expecteds := []string{
		`href="https://example.com/filemanager"`,
		`target="_blank"`,
		`Browse`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}
