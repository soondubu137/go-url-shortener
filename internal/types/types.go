// Code generated by goctl. DO NOT EDIT.
package types

type RestoreRequest struct {
	ShortURL string `json:"short_url"`
}

type RestoreResponse struct {
	OriginalURL string `json:"original_url"`
}

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
