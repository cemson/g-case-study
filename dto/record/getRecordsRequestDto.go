package record

import "time"

type GetRecordsRequestDto struct {
	StartDate   *string `json:"startDate"`
	EndDate     *string `json:"endDate"`
	MinCount    *int    `json:"minCount"`
	MaxCount    *int    `json:"maxCount"`
	StartDateDt time.Time
	EndDateDt   time.Time
}
