package validators

import (
	"g-case-study/dto/record"
	"time"
)

type RecordsValidator struct {
}

func (t *RecordsValidator) ValidateGetRequestDto(dto *record.GetRecordsRequestDto) ValidationResult {
	if dto == nil {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "startDate field can not be empty",
			InvalidFields:          &[]string{"startDate"},
		}
	}

	if dto.StartDate == nil || *dto.StartDate == "" {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "startDate field can not be empty, should be an ISO date",
			InvalidFields:          &[]string{"startDate"},
		}
	}

	startDate, err := time.Parse("2006-01-02", *dto.StartDate)
	if err != nil {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "startDate field can not be empty, should be an ISO date",
			InvalidFields:          &[]string{"startDate"},
		}
	}
	dto.StartDateDt = startDate

	if dto.EndDate == nil || *dto.EndDate == "" {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "endDate field can not be empty, should be an ISO date",
			InvalidFields:          &[]string{"endDate"},
		}
	}

	endDate, err := time.Parse("2006-01-02", *dto.EndDate)
	if err != nil {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "endDate field can not be empty, should be an ISO date",
			InvalidFields:          &[]string{"endDate"},
		}
	}
	dto.EndDateDt = endDate

	if dto.MinCount == nil || *dto.MinCount < 0 {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "minCount field should be an non-negative integer and should not be empty",
			InvalidFields:          &[]string{"minCount"},
		}
	}

	if dto.MaxCount == nil || *dto.MaxCount < 0 {
		return ValidationResult{
			IsValid:                false,
			ValidationErrorMessage: "maxCount field should be an non-negative integer and should not be empty",
			InvalidFields:          &[]string{"maxCount"},
		}
	}

	return ValidationResult{IsValid: true}
}
