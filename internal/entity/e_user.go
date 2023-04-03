package entity

import (
	"time"
)

// 系统用户表
type User struct {
	ID        int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Username  string    `gorm:"column:username;NOT NULL"`
	Password  string    `gorm:"column:password;NOT NULL"`
	Realname  string    `gorm:"column:realname"`
	RoleId    int64     `gorm:"column:roleId;NOT NULL"`
	Email     string    `gorm:"column:email;NOT NULL"`
	Cellphone string    `gorm:"column:cellphone;NOT NULL"`
	CreatedAt time.Time `gorm:"column:createdAt"`        // 创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt"`        // 更新时间
	Status    int32     `gorm:"column:status;default:1"` // 1为启用，0为禁用
	IsDel     int32     `gorm:"column:isDel;default:0"`
}

func (m *User) TableName() string {
	return "m_user"
}
