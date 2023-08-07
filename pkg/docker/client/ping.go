package client

import (
	"context"

	"github.com/docker/docker/api/types"
)

func (cli *Client) Ping(ctx context.Context) (types.Ping, error) {
	var ping types.Ping

	req, err := cli.buildRequest
}
