package exceptions

type ValidationException struct {
	AdditionalInfo         map[string]interface{}
	InvalidFields          *[]string
	ValidationErrorMessage *string
	ExceptionBase
}
