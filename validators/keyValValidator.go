package validators

import "g-case-study/dto"

type KeyValValidator struct {
}

func (t *KeyValValidator) ValidateKeyValDto(dto *dto.KeyValDto) ValidationResult {
	keyValidation := t.ValidateKey(dto)
	if !keyValidation.IsValid {
		return keyValidation
	}

	if dto.Value == "" {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "Please add value field",
			InvalidFields:          &[]string{"value"},
		}
	}

	return ValidationResult{
		IsValid: true,
	}
}
func (t *KeyValValidator) ValidateKey(dto *dto.KeyValDto) ValidationResult {
	if dto == nil {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "Please add key field",
			InvalidFields:          &[]string{"key"},
		}
	}

	if dto.Key == "" {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "Key field can not be empty",
			InvalidFields:          &[]string{"key"},
		}
	}

	return ValidationResult{IsValid: true}
}
