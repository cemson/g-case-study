package response

type ApiResponse struct {
	AdditionalData map[string]interface{} `json:"additionalData"`
	Records        *interface{}           `json:"records"`
	*MessageResponseBase
}
