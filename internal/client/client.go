package client

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	grpcinsecure "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn grpc.ClientConnInterface
	etcd *etcd.Registry
}

func NewGrpcClient(etcd *etcd.Registry, c *conf.App) (*GrpcClient, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Etcd.GetTimeout().AsDuration())
	conn, err := getConn(ctx, etcd, "discovery:///demo")
	if err != nil {
		return nil, cancel, err
	}
	return &GrpcClient{
		conn: conn,
		etcd: etcd,
	}, cancel, nil
}

func (g *GrpcClient) GetConn() grpc.ClientConnInterface {
	return g.conn
}

func getConn(ctx context.Context, dis *etcd.Registry, endpoint string) (*grpc.ClientConn, error) {

	return grpcinsecure.DialInsecure(ctx, grpcinsecure.WithDiscovery(dis), grpcinsecure.WithEndpoint(endpoint))
}

var ProviderSet = wire.NewSet(NewGrpcClient, NewLoginClient)
