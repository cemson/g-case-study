package tests

import (
	"g-case-study/dto/record"
	"g-case-study/globals"
	"g-case-study/repo"
	"g-case-study/settings"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindRecordSummaries(t *testing.T) {
	appSettings := settings.AppSettings{
		MongoDbName:    "getircase-study",
		MongoDbAddress: "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true",
	}
	globals.SetAppSettings(&appSettings)

	recordsRepo := repo.CreateRecordsRepository()

	startDate := time.Date(2016, 1, 28, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 2, 2, 0, 0, 0, 0, time.UTC)
	countStart := 2800
	countEnd := 3000

	summaries, err := recordsRepo.FindRecordSummaries(startDate, endDate, countStart, countEnd)
	assert.NoError(t, err)
	assert.NotEmpty(t, summaries)

	// Assert that the result contains at least one RecordSummaryDto object
	assert.IsType(t, []record.RecordSummaryDto{}, summaries)
	assert.GreaterOrEqual(t, len(summaries), 1)

	// Assert that the RecordSummaryDto object's totalCount is within the range [countStart, countEnd]
	for _, s := range summaries {
		assert.GreaterOrEqual(t, s.TotalCount, countStart)
		assert.LessOrEqual(t, s.TotalCount, countEnd)
	}
}
