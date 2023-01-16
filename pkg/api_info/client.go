package api_info

import (
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/api_info"
)

func ApiClient(apiKey string, clientId string) *api_info.Client {
	return api_info.New(api.NewClient(apiKey, clientId))
}
