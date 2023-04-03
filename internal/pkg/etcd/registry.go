package etcd

import (
	"github.com/go-kratos/kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/google/wire"
	etcdclient "go.etcd.io/etcd/client/v3"
)

func NewEtcdRegistry(conf *conf.App_Etcd) *etcd.Registry {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{conf.Url},
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

var ProviderSet = wire.NewSet(NewEtcdRegistry)
