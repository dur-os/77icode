package user

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/dur-os/77icode/Server/common"
	"github.com/jinzhu/gorm"
)

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

func (user *AdminUser) Login(db *gorm.DB) bool {
	m := md5.Sum([]byte(user.Password))
	password := hex.EncodeToString(m[:])
	var userData AdminUser
	db.Where("fusername = ? ", user.Username).First(&userData)
	if &userData == nil || password != userData.Password {
		return false
	} else {
		user.Password = ""
		user.Name = userData.Name
		return true
	}
}
