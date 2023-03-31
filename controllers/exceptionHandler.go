package controllers

import (
	"g-case-study/consts"
	"g-case-study/dto/response"
	"g-case-study/exceptions"
	"g-case-study/logging"
	"g-case-study/utilities"
	"net/http"
)

func UnhandledException(responseWriter http.ResponseWriter) {
	if err := recover(); err != nil {
		switch err.(type) {
		default:
			logging.Log.Errorf("Unhandled exception during requests handling.\n%v", err)
			utilities.WriteResponse(responseWriter, http.StatusInternalServerError, response.ApiResponse{
				AdditionalData: nil,
				Records:        nil,
				MessageResponseBase: &response.MessageResponseBase{
					Message: "Unhandled exception during requests handling",
					ResponseBase: response.ResponseBase{
						IsCompletedSuccessfully: false,
						IsValidationError:       false,
						NotValidFields:          nil,
						IsUnexpectedError:       true,
						Code:                    consts.UnexpectedError,
					},
				},
			})
		case exceptions.ValidationException:
			validationException := err.(exceptions.ValidationException)
			logging.Log.Warnf("Validation exception during requests handling.\n%v", utilities.JsonSerialize(validationException))
			utilities.WriteResponse(responseWriter, http.StatusBadRequest, response.ApiResponse{
				AdditionalData: validationException.AdditionalInfo,
				Records:        nil,
				MessageResponseBase: &response.MessageResponseBase{
					Message: *validationException.ValidationErrorMessage,
					ResponseBase: response.ResponseBase{
						IsCompletedSuccessfully: false,
						IsValidationError:       true,
						NotValidFields:          validationException.InvalidFields,
						IsUnexpectedError:       false,
						Code:                    consts.ValidationError,
					},
				},
			})
		}
	}
}
