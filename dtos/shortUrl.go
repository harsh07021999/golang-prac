package dtos

type ShortUrlRequest struct {
	CustomHash  string `json:"custom_hash"`
	OriginalUrl string `json:"original_url"`
}
