package record

type RecordSummaryDto struct {
	Key        string `bson:"key" json:"key"`
	CreatedAt  string `bson:"createdAt" json:"createdAt"`
	TotalCount int    `bson:"totalCount" json:"totalCount"`
}
