package record

import "time"

type RecordSummaryDto struct {
	Key        string    `bson:"key" json:"key"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	TotalCount int       `bson:"totalCount" json:"totalCount"`
}
