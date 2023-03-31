package services

import (
	"g-case-study/dto/record"
	"g-case-study/repo"
)

type RecordsService struct {
	recordsRepository *repo.RecordsRepository
}

func CreateRecordsService() RecordsService {
	return RecordsService{recordsRepository: repo.CreateRecordsRepository()}
}

func (t *RecordsService) GetRecords(requestDto *record.GetRecordsRequestDto) []record.RecordSummaryDto {
	result, err := t.recordsRepository.FindRecordSummaries(requestDto.StartDateDt, requestDto.EndDateDt, *requestDto.MinCount, *requestDto.MaxCount)
	if err != nil {
		panic(err)
	}

	return result
}
