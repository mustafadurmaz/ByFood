package models

type ProcessURLRequest struct {
	URL       string `json:"url" binding:"required"`
	Operation string `json:"operation" binding:"required,oneof=canonical redirection all"`
}

type ProcessURLResponse struct {
	ProcessedURL string `json:"processed_url"`
}