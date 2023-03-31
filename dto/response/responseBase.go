package response

type ResponseBase struct {
	IsCompletedSuccessfully bool      `json:"isCompletedSuccessfully"`
	IsValidationError       bool      `json:"isValidationError"`
	NotValidFields          *[]string `json:"notValidFields"`
	IsUnexpectedError       bool      `json:"isUnexpectedError"`
	Code                    int       `json:"code"`
}
