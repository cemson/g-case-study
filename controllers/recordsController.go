package controllers

import (
	"encoding/json"
	"g-case-study/dto/record"
	"g-case-study/services"
	"g-case-study/utilities"
	"g-case-study/validators"
	"net/http"
)

type RecordsController struct {
	recordsService   services.RecordsService
	recordsValidator validators.RecordsValidator
}

func CreateRecordsController() RecordsController {
	return RecordsController{recordsService: services.CreateRecordsService(),
		recordsValidator: validators.RecordsValidator{}}
}

func (t *RecordsController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t.handlePost(w, r)
	} else {
		utilities.MethodNotAllowed(w)
	}
}
func (t *RecordsController) handlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestDto record.GetRecordsRequestDto
	decoder.Decode(&requestDto)

	CallServiceAndRespond(w, func() validators.ValidationResult {
		return t.recordsValidator.ValidateGetRequestDto(&requestDto)
	}, func() interface{} {
		return t.recordsService.GetRecords(&requestDto)
	}, nil)
}
