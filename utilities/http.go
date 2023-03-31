package utilities

import (
	"encoding/json"
	"fmt"
	"g-case-study/logging"
	"net/http"
)

func MethodNotAllowed(responseWriter http.ResponseWriter) {
	WriteResponse(responseWriter, http.StatusMethodNotAllowed, nil)
}
func BadRequest(responseWriter http.ResponseWriter, errorMessage string) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	writeMessageResponse(responseWriter, errorMessage)
}
func Success(responseWriter http.ResponseWriter, responseBody interface{}) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	err := json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		logging.Log.Error("Json error: %v", err)
		return
	}
}
func SuccessWithMessage(responseWriter http.ResponseWriter, message string) {
	writeMessageResponseWithStatus(responseWriter, message, http.StatusOK)
}
func Forbidden(responseWriter http.ResponseWriter, message string) {
	writeMessageResponseWithStatus(responseWriter, message, http.StatusForbidden)
}
func InternalServerError(responseWriter http.ResponseWriter, errorMessage string) {
	writeMessageResponseWithStatus(responseWriter, errorMessage, http.StatusInternalServerError)
}

func writeMessageResponseWithStatus(responseWriter http.ResponseWriter, message string, status int) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	writeMessageResponse(responseWriter, message)
}
func writeMessageResponse(w http.ResponseWriter, message string) {
	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("%s.", message)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		logging.Log.Error("Json error: %v", err)
		return
	}
}
func WriteResponse(responseWriter http.ResponseWriter, statusCode int, responseBody interface{}) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)

	if responseBody != nil {
		err := json.NewEncoder(responseWriter).Encode(responseBody)
		if err != nil {
			logging.Log.Errorf("Json error: %v", err)
			return
		}
	}
}
