package structs

import "time"

type ById struct {
	Id  string `json:"id" binding:"required"`
	Ref string `json:"ref" binding:"required"`
}

type CreateUrl struct {
	Url       string     `json:"url" binding:"required"`
	ExpiresAt *time.Time `json:"expires_at"`
	Ref       string     `json:"ref"`
}
