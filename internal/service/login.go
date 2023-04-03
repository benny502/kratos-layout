package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/login/v1"
	biz "github.com/go-kratos/kratos-layout/internal/biz/login"
	"github.com/go-kratos/kratos-layout/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/log"
)

type LoginService struct {
	v1.UnimplementedLoginServer
	loginBiz *biz.BizLogin
	log      *log.Helper
}

func NewLoginService(biz *biz.BizLogin, logger log.Logger) *LoginService {
	return &LoginService{
		loginBiz: biz,
		log:      log.NewHelper(logger),
	}
}

func (s *LoginService) Vcode(ctx context.Context, req *v1.VcodeRequest) (*v1.VcodeReply, error) {
	capInfo, err := s.loginBiz.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	// cookie
	auth.SetCustomCookie(ctx, "vahash="+capInfo.S)
	return capInfo, nil
}
func (s *LoginService) Check(ctx context.Context, req *v1.CheckRequest) (*v1.CheckReply, error) {
	r, err := s.loginBiz.VerifyUser(ctx, req)
	if err != nil {
		return nil, err
	}
	auth.SetCustomCookie(ctx, "stoken="+r.Token)
	return r, nil
}
func (s *LoginService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return s.loginBiz.Logout(ctx, req)
}
