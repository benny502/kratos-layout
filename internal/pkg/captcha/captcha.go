package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type CaptchaInfo struct {
	CaptchaId string
	PicPath   string
}

// GetCaptcha 生成验证码
func GetCaptcha(ctx context.Context) (*CaptchaInfo, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		return nil, err
	}

	return &CaptchaInfo{
		CaptchaId: id,
		PicPath:   b64s,
	}, nil
}

func VerifyCaptcha(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
