package model

import (
	"time"
)

type ResponseCommon struct {
	Messages     []map[string]string `json:"messages,omitempty"`
	Status       int                 `json:"status"`
	ResponseTime time.Time           `json:"responsetime"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Count       int `json:"count"`
}
