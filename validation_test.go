package form

import (
	"testing"
)

func TestValidateRequired(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:     "email",
				Type:     FORM_FIELD_TYPE_STRING,
				Required: true,
			},
		},
	})

	errors := form.Validate(map[string]string{
		"email": "",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}

	if errors[0].Field != "email" {
		t.Fatal("Expected field 'email', got:", errors[0].Field)
	}
}

func TestValidateRequiredPasses(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:     "email",
				Type:     FORM_FIELD_TYPE_STRING,
				Required: true,
			},
		},
	})

	errors := form.Validate(map[string]string{
		"email": "test@example.com",
	})

	if len(errors) != 0 {
		t.Fatal("Expected 0 errors, got:", len(errors))
	}
}

func TestValidateMinLength(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:       "password",
				Type:       FORM_FIELD_TYPE_PASSWORD,
				Validators: []Validator{ValidatorMinLength(8)},
			},
		},
	})

	errors := form.Validate(map[string]string{
		"password": "short",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}
}

func TestValidateMaxLength(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:       "username",
				Type:       FORM_FIELD_TYPE_STRING,
				Validators: []Validator{ValidatorMaxLength(5)},
			},
		},
	})

	errors := form.Validate(map[string]string{
		"username": "toolongname",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}
}

func TestValidateMin(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:       "age",
				Type:       FORM_FIELD_TYPE_NUMBER,
				Validators: []Validator{ValidatorMin(18)},
			},
		},
	})

	errors := form.Validate(map[string]string{
		"age": "10",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}
}

func TestValidateMax(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:       "quantity",
				Type:       FORM_FIELD_TYPE_NUMBER,
				Validators: []Validator{ValidatorMax(100)},
			},
		},
	})

	errors := form.Validate(map[string]string{
		"quantity": "200",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}
}

func TestValidatePattern(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:       "code",
				Type:       FORM_FIELD_TYPE_STRING,
				Validators: []Validator{ValidatorPattern(`^[A-Z]{3}$`, "must be 3 uppercase letters")},
			},
		},
	})

	errors := form.Validate(map[string]string{
		"code": "abc",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}

	errors = form.Validate(map[string]string{
		"code": "ABC",
	})

	if len(errors) != 0 {
		t.Fatal("Expected 0 errors, got:", len(errors))
	}
}

func TestValidateEmail(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:       "email",
				Type:       FORM_FIELD_TYPE_EMAIL,
				Validators: []Validator{ValidatorEmail()},
			},
		},
	})

	errors := form.Validate(map[string]string{
		"email": "not-an-email",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error, got:", len(errors))
	}

	errors = form.Validate(map[string]string{
		"email": "user@example.com",
	})

	if len(errors) != 0 {
		t.Fatal("Expected 0 errors, got:", len(errors))
	}
}

func TestValidateMultipleValidators(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				Name:     "password",
				Type:     FORM_FIELD_TYPE_PASSWORD,
				Required: true,
				Validators: []Validator{
					ValidatorMinLength(8),
					ValidatorMaxLength(64),
				},
			},
		},
	})

	// Empty value should trigger required error only
	errors := form.Validate(map[string]string{
		"password": "",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error for empty required field, got:", len(errors))
	}

	// Short value should trigger minLength
	errors = form.Validate(map[string]string{
		"password": "short",
	})

	if len(errors) != 1 {
		t.Fatal("Expected 1 error for short password, got:", len(errors))
	}

	// Valid value should pass
	errors = form.Validate(map[string]string{
		"password": "validpassword123",
	})

	if len(errors) != 0 {
		t.Fatal("Expected 0 errors, got:", len(errors))
	}
}

func TestValidateNoFields(t *testing.T) {
	form := NewForm(FormOptions{})

	errors := form.Validate(map[string]string{
		"anything": "value",
	})

	if len(errors) != 0 {
		t.Fatal("Expected 0 errors for empty form, got:", len(errors))
	}
}
