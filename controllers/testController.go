package controllers

import (
	"g-case-study/utilities"
	"net/http"
)

type TestController struct {
}

func CreateTestController() TestController {
	return TestController{}
}

func (t TestController) HandleHealthCheck(responseWriter http.ResponseWriter, request *http.Request) {
	utilities.Success(responseWriter, map[string]string{"message": "g-case-study rest api is well alive..."})
}
