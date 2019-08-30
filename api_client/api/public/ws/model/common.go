package model

import "api_client/api/common/configuration"

// ResCommon ...
type ResCommon struct {
	Channel configuration.WebSocketChannel `json:"channel"`
}
