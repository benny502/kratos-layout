package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos-layout/internal/entity"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUser(ctx context.Context, username string, password string) (*UserData, error)
	FindByRealname(ctx context.Context, realname string) []string
	ListAll(ctx context.Context) []*entity.User
}

type userRepo struct {
	db  *Data
	log *log.Helper
}

type UserData struct {
	ID        int64
	Username  string
	Realname  string
	Role      string
	RoleId    int64
	Email     string
	Cellphone string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int32
	IsDel     int32
}

func (s *userRepo) GetUser(ctx context.Context, username string, password string) (*UserData, error) {
	user := &UserData{}

	rs := s.db.GetDB().Model(entity.User{}).
		Select("m_user.id,m_user.username,m_user.realname,m_user.roleId RoleId,m_user.email,m_user.cellphone,m_user.createdAt CreatedAt,m_user.updatedAt UpdatedAt,m_user.status,m_user.isDel,m_role.role").
		Where("(username = ? OR email = ?) AND password = ?", username, username, password).
		Joins("LEFT JOIN m_role on m_user.roleId=m_role.roleId").
		First(&user)
	if rs.Error != nil {
		s.log.Errorf("Mysql error: %v", rs.Error)
		return nil, rs.Error
	}
	return user, nil
}

// FindByRealname implements systemUserRepo
func (s *userRepo) FindByRealname(ctx context.Context, realname string) []string {
	ids := make([]string, 0, 10)
	s.db.GetDB().Model(entity.User{}).Where("realname LIKE ?", fmt.Sprintf("%%%s%%", realname)).Find(ids)
	return ids
}

func (s *userRepo) ListAll(ctx context.Context) []*entity.User {
	users := make([]*entity.User, 0, 10)
	s.db.GetDB().Model(&entity.User{}).Find(&users)
	return users
}

func NewUserData(data *Data, logger log.Logger) UserRepo {
	return &userRepo{
		db:  data,
		log: log.NewHelper(logger),
	}
}
