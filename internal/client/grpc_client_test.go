package client

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/pkg/etcd"

	v1 "github.com/go-kratos/kratos-layout/api/login/v1"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestCasesDashboardReviewd(t *testing.T) {

	registry := etcd.NewEtcdRegistry(&conf.App_Etcd{
		Url:     "127.0.0.1:2379",
		Timeout: &durationpb.Duration{Seconds: 1},
	})

	client, cancel, err := NewGrpcClient(registry, &conf.App{
		Etcd: &conf.App_Etcd{
			Url:     "127.0.0.1:2379",
			Timeout: &durationpb.Duration{Seconds: 1},
		},
	})
	if err != nil {
		panic(err)
	}
	defer cancel()
	login := NewLoginClient(client, log.GetLogger())
	_, err = login.Check(context.Background(), &v1.CheckRequest{
		Account:  "root",
		Password: "123456",
		Vcode:    "abcd",
	})
	if err != nil {
		panic(err)
	}
}
