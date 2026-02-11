package form

// Theme defines the CSS classes used when rendering form fields.
// This allows decoupling from any specific CSS framework.
type Theme struct {
	FormGroupClass     string
	LabelClass         string
	InputClass         string
	SelectClass        string
	TextAreaClass      string
	CheckboxWrapClass  string
	CheckboxInputClass string
	RadioWrapClass     string
	RadioInputClass    string
	RadioLabelClass    string
	FileInputClass     string
	HelpClass          string
	RequiredClass      string
	RequiredMarker     string
	TableClass         string
	ErrorClass         string // CSS class for the error message element
	ErrorInputClass    string // CSS class added to invalid inputs
}

// ThemeBootstrap5 returns the default Bootstrap 5 theme.
func ThemeBootstrap5() *Theme {
	return &Theme{
		FormGroupClass:     "form-group mb-3",
		LabelClass:         "form-label",
		InputClass:         "form-control",
		SelectClass:        "form-select",
		TextAreaClass:      "form-control",
		CheckboxWrapClass:  "form-check",
		CheckboxInputClass: "form-check-input",
		RadioWrapClass:     "form-check",
		RadioInputClass:    "form-check-input",
		RadioLabelClass:    "form-check-label",
		FileInputClass:     "form-control",
		HelpClass:          "text-info",
		RequiredClass:      "text-danger ms-1",
		RequiredMarker:     "*",
		TableClass:         "table table-striped table-hover mb-0",
		ErrorClass:         "invalid-feedback",
		ErrorInputClass:    "is-invalid",
	}
}

// ThemeTailwind returns a Tailwind CSS theme.
func ThemeTailwind() *Theme {
	return &Theme{
		FormGroupClass:     "mb-4",
		LabelClass:         "block text-sm font-medium text-gray-700 mb-1",
		InputClass:         "block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm",
		SelectClass:        "block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm",
		TextAreaClass:      "block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm",
		CheckboxWrapClass:  "flex items-center",
		CheckboxInputClass: "h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500",
		RadioWrapClass:     "flex items-center",
		RadioInputClass:    "h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-500",
		RadioLabelClass:    "ml-2 block text-sm text-gray-900",
		FileInputClass:     "block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold file:bg-indigo-50 file:text-indigo-700 hover:file:bg-indigo-100",
		HelpClass:          "mt-1 text-sm text-gray-500",
		RequiredClass:      "text-red-500 ml-1",
		RequiredMarker:     "*",
		TableClass:         "min-w-full divide-y divide-gray-200",
		ErrorClass:         "mt-1 text-sm text-red-600",
		ErrorInputClass:    "border-red-500",
	}
}

// defaultTheme is the package-level default theme (Bootstrap 5).
var defaultTheme = ThemeBootstrap5()
