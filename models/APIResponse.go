package models

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"messages"`
}
