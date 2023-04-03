package demo

import (
	"context"
	"fmt"
	"os"
	"time"

	v1 "github.com/go-kratos/kratos-layout/api/demo/v1"
	"github.com/go-kratos/kratos-layout/internal/data"
	"github.com/go-kratos/kratos-layout/internal/pkg/excel"
	"github.com/go-kratos/kratos/v2/log"
)

type BizDemo struct {
	UserRepo data.UserRepo
	log      *log.Helper
}

func (b *BizDemo) Excel(ctx context.Context, req *v1.ExcelRequest) (*v1.ExcelReply, error) {

	sheetName := "Sheet1"

	excel := excel.NewHelper(log.GetLogger())

	header := []string{
		"User ID",
		"用户名",
		"真实姓名",
		"角色",
		"Email",
		"手机",
		"创建时间",
		"更新时间",
		"状态",
	}

	excel.Header(sheetName, header)

	users := b.UserRepo.ListAll(ctx)

	data := make([][]interface{}, 0, len(users))

	for _, user := range users {
		values := make([]interface{}, 0)
		values = append(values, user.ID)
		values = append(values, user.Username)
		values = append(values, user.Realname)
		values = append(values, user.Realname)
		values = append(values, user.RoleId)
		values = append(values, user.Email)
		values = append(values, user.Cellphone)

		values = append(values, user.CreatedAt.Format("2006-01-02_15:04:05"))

		values = append(values, user.UpdatedAt.Format("2006-01-02_15:04:05"))

		values = append(values, user.Status)

		data = append(data, values)

	}

	excel.Data(sheetName, data)

	now := time.Now().Format("2006-01-02_15:04:05")
	filename := fmt.Sprintf("用户列表_%s.xlsx", now)

	err := excel.SaveAs(os.TempDir() + filename)

	if err != nil {
		return nil, err
	}
	return &v1.ExcelReply{
		Filename: filename,
		Path:     "/file/" + filename,
	}, nil

}

func NewBizDemo(UserData data.UserRepo, logger log.Logger) *BizDemo {
	return &BizDemo{
		UserRepo: UserData,
		log:      &log.Helper{},
	}
}
