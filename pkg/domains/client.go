package domains

import (
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/api_info"
	"github.com/sitehostnz/gosh/pkg/api/domain"
)

func DomainClient(apiKey string, clientId string) *domain.Client {
	return domain.New(api.NewClient(apiKey, clientId))
}

func ApiClient(apiKey string, clientId string) *api_info.Client {
	return api_info.New(api.NewClient(apiKey, clientId))
}
