package admin

import "github.com/dur-os/77icode/Server/common"

func init() {
	common.RegisterDBModel(&AdminUser{})
}

type AdminUser struct {
	//gorm.Model
	ID       int    `gorm:"column:fid;primary_key;AUTO_INCREMENT;not null"`
	Username string `gorm:"column:fusername;size:50;not null"`
	Password string `gorm:"column:fpassword;size:50;not null"`
	Name     string `gorm:"column:fname;size:50;not null"`
}

func (AdminUser) TableName() string {
	return "t_admin_user"
}
