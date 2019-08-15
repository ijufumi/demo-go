package model

import (
	"time"
)

type ResponseCommon struct {
	Messages     []map[string]string `json:"messages"`
	Status       int                 `json:"status"`
	ResponseTime time.Time           `json:"responsetime"`
}
