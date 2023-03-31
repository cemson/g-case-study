package controllers

import (
	"g-case-study/consts"
	"g-case-study/dto/response"
	"g-case-study/exceptions"
	"g-case-study/utilities"
	"g-case-study/validators"
	"net/http"
)

func CallServiceAndRespond(responseWriter http.ResponseWriter, validationFunction func() validators.ValidationResult,
	serviceFunction func() interface{}, serviceFunctionReturnsNothing func()) {
	defer UnhandledException(responseWriter)

	if validationFunction != nil {
		validationResult := validationFunction()
		if !validationResult.IsValid {
			panic(exceptions.ValidationException{
				AdditionalInfo:         validationResult.AdditionalInfo,
				InvalidFields:          validationResult.InvalidFields,
				ValidationErrorMessage: &validationResult.ValidationErrorMessage,
			})
		}
	}

	if serviceFunction == nil && serviceFunctionReturnsNothing == nil {
		panic("service function is empty")
	}

	if serviceFunction != nil {
		serviceResult := serviceFunction()
		utilities.WriteResponse(responseWriter, http.StatusOK, response.ApiResponse{
			AdditionalData: nil,
			Records:        &serviceResult,
			MessageResponseBase: &response.MessageResponseBase{
				Message: "Request handled successfully",
				ResponseBase: response.ResponseBase{
					IsCompletedSuccessfully: true,
					IsValidationError:       false,
					NotValidFields:          nil,
					IsUnexpectedError:       false,
					Code:                    consts.Success,
				},
			},
		})
	} else {
		serviceFunctionReturnsNothing()
		utilities.WriteResponse(responseWriter, http.StatusOK, response.ApiResponse{
			AdditionalData: nil,
			Records:        nil,
			MessageResponseBase: &response.MessageResponseBase{
				Message: "Request handled successfully",
				ResponseBase: response.ResponseBase{
					IsCompletedSuccessfully: true,
					IsValidationError:       false,
					NotValidFields:          nil,
					IsUnexpectedError:       false,
					Code:                    consts.Success,
				},
			},
		})
	}
}
