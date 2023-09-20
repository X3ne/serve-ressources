package responses

type CdnResponse struct {
	Url string `json:"url"`
}

func NewCdnResponse(url string) *CdnResponse {
	return &CdnResponse{
		Url: url,
	}
}
