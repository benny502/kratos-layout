package login

import (
	"context"
	"errors"

	api "github.com/go-kratos/kratos-layout/api/login/v1"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/data"
	"github.com/go-kratos/kratos-layout/internal/pkg/captcha"
	"github.com/go-kratos/kratos-layout/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
)

type BizLogin struct {
	UserRepo data.UserRepo
	c        *conf.Auth
	log      *log.Helper
}

func NewBizLogin(UserData data.UserRepo, c *conf.Auth, logger log.Logger) *BizLogin {
	return &BizLogin{
		UserRepo: UserData,
		c:        c,
		log:      log.NewHelper(logger),
	}
}

// GetCaptcha 验证码
func (l *BizLogin) GetCaptcha(ctx context.Context) (*api.VcodeReply, error) {
	captchaInfo, err := captcha.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return &api.VcodeReply{
		Img: captchaInfo.PicPath,
		S:   captchaInfo.CaptchaId,
	}, nil
}

func (l *BizLogin) VerifyUser(ctx context.Context, req *api.CheckRequest) (*api.CheckReply, error) {
	account, password, vcode := req.GetAccount(), req.GetPassword(), req.GetVcode()
	vCodeId := auth.GetCustomCookie(ctx, "vahash")
	if !captcha.VerifyCaptcha(vCodeId, vcode) {
		return nil, errors.New("vcode error")
	}
	user, err := l.UserRepo.GetUser(ctx, account, password)
	if err != nil {
		return nil, err
	}
	if user.Status == 0 {
		return nil, errors.New("账号已禁用")
	}
	return &api.CheckReply{
		Id:        user.ID,
		Username:  user.Username,
		Realname:  user.Realname,
		RoleId:    user.RoleId,
		Email:     user.Email,
		Cellphone: user.Cellphone,
		CreatedAt: user.CreatedAt.Format("2006/01/02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006/01/02 15:04:05"),
		Status:    user.Status,
		IsDel:     user.IsDel,
		Role:      user.Role,
		Token:     auth.GenerateToken(user.ID, l.c.GetJwtKey(), l.c.GetIssuer()),
	}, nil
}

func (l *BizLogin) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutReply, error) {
	return &api.LogoutReply{}, nil
}
