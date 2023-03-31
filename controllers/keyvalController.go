package controllers

import (
	"encoding/json"
	"g-case-study/dto"
	"g-case-study/services"
	"g-case-study/utilities"
	"g-case-study/validators"
	"net/http"
)

type KeyValController struct {
	keyValService   services.KeyValService
	keyValValidator validators.KeyValValidator
}

func CreateKeyValValidator() KeyValController {
	testService := services.CreateKeyValService()
	return KeyValController{
		keyValService:   testService,
		keyValValidator: validators.KeyValValidator{},
	}
}

func (t *KeyValController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t.handleSet(w, r)
	} else if r.Method == http.MethodGet {
		t.handleGet(w, r)
	} else {
		utilities.MethodNotAllowed(w)
	}
}

func (t *KeyValController) handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	keyValDto := dto.KeyValDto{
		Key: key,
	}

	CallServiceAndRespond(w, func() validators.ValidationResult {
		return t.keyValValidator.ValidateKey(&keyValDto)
	}, func() interface{} {
		return t.keyValService.Get(keyValDto.Key)
	}, nil)
}
func (t *KeyValController) handleSet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var keyValDto dto.KeyValDto
	decoder.Decode(&keyValDto)

	CallServiceAndRespond(w, func() validators.ValidationResult {
		return t.keyValValidator.ValidateKeyValDto(&keyValDto)
	}, func() interface{} {
		return t.keyValService.Set(&keyValDto)
	}, nil)
}
