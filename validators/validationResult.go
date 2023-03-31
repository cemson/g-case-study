package validators

type ValidationResult struct {
	IsValid                bool
	ValidationErrorMessage string
	InvalidFields          *[]string
	AdditionalInfo         map[string]interface{}
}
