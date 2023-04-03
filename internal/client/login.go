package client

import (
	"context"

	api "github.com/go-kratos/kratos-layout/api/login/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type LoginClient struct {
	client *GrpcClient
	logger log.Helper
}

func NewLoginClient(client *GrpcClient, logger log.Logger) *LoginClient {
	return &LoginClient{
		client: client,
		logger: *log.NewHelper(logger),
	}
}

func (d *LoginClient) Check(ctx context.Context, req *api.CheckRequest) (*api.CheckReply, error) {
	client := api.NewLoginClient(d.client.GetConn())
	return client.Check(ctx, req)
}
