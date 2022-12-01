package googleworkspace_client

import (
	"context"
	"errors"
	"github.com/selefra/selefra-provider-googleworkspace/constants"
	"strings"
)

type Client struct {
	Config *Config
}

func NewClients(configs Configs) ([]*Client, error) {
	var clients []*Client
	for i := range configs.Providers {
		clients = append(clients, &Client{Config: &configs.Providers[i]})
	}
	return clients, nil
}

func IsCancelled(ctx context.Context) bool {
	err := ctx.Err()
	return err != nil && (errors.Is(err, context.Canceled) || strings.Contains(err.Error(), constants.Contextcanceled))
}
