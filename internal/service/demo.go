package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/demo/v1"
	biz "github.com/go-kratos/kratos-layout/internal/biz/demo"
	"github.com/go-kratos/kratos-layout/internal/pkg/middleware/localize"
	"github.com/go-kratos/kratos/v2/log"
)

type DemoService struct {
	v1.UnimplementedDemoServer
	demo *biz.BizDemo
	log  *log.Helper
}

func (*DemoService) I18N(ctx context.Context, req *v1.I18NRequest) (*v1.I18NReply, error) {
	message, err := localize.Localize(ctx, "TEAM_10004")
	if err != nil {
		return nil, err
	}
	return &v1.I18NReply{
		Result: message,
	}, err

}

func (d *DemoService) Excel(ctx context.Context, req *v1.ExcelRequest) (*v1.ExcelReply, error) {
	return d.demo.Excel(ctx, req)
}

func NewDemoService(demo *biz.BizDemo, logger log.Logger) *DemoService {
	return &DemoService{
		demo: demo,
		log:  log.NewHelper(logger),
	}
}
