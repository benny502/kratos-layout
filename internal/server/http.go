package server

import (
	"context"
	"os"

	demoApi "github.com/go-kratos/kratos-layout/api/demo/v1"
	loginApi "github.com/go-kratos/kratos-layout/api/login/v1"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/pkg/errors"
	"github.com/go-kratos/kratos-layout/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos-layout/internal/pkg/middleware/localize"
	"github.com/go-kratos/kratos-layout/internal/pkg/response"
	"github.com/go-kratos/kratos-layout/internal/pkg/websocket"
	"github.com/go-kratos/kratos-layout/internal/service"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/mux"

	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, login *service.LoginService, demo *service.DemoService, websocket *websocket.Websocket, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(ac.JwtKey)).Match(NewSkipRoutersMatcher()).Build(),
			localize.I18N(),
		),
		http.ErrorEncoder(errors.ErrorEncoder),
		http.ResponseEncoder(response.ResponseEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	loginApi.RegisterLoginHTTPServer(srv, login)
	demoApi.RegisterDemoHTTPServer(srv, demo)

	router := mux.NewRouter()

	openapi := openapiv2.NewHandler()

	router.HandleFunc("/ws", websocket.Handler)

	srv.HandlePrefix("/q/", openapi)

	srv.HandlePrefix("/file/", nethttp.StripPrefix("/file/", nethttp.FileServer(nethttp.Dir(os.TempDir()))))

	srv.HandlePrefix("/", router)
	return srv
}

func NewSkipRoutersMatcher() selector.MatchFunc {
	skipRouters := map[string]struct{}{
		"/login.v1.Login/Check": {},
		"/demo.v1.Demo/I18N":    {},
		"/demo.v1.Demo/Excel":   {},
	}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}
