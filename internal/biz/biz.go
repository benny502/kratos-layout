package biz

import (
	"github.com/go-kratos/kratos-layout/internal/biz/demo"
	"github.com/go-kratos/kratos-layout/internal/biz/login"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(login.NewBizLogin, demo.NewBizDemo)
