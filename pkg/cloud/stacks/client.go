package stacks

import (
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
)

func StackClient(apiKey string, clientId string) *stack.Client {
	return stack.New(api.NewClient(apiKey, clientId))
}
