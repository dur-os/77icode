package admin

import (
	"net/http"

	"github.com/dur-os/77icode/Server/common"
	"github.com/dur-os/77icode/Server/model/admin"
	"github.com/zenazn/goji/web"
)

func init() {

}

// UserController is admin user Controller
type UserController struct {
	common.Controller
}

func (controller *UserController) getUser(c web.C, r *http.Request) common.ReturnData {
	session := controller.GetSession(c)
	userInfo := session.Get("userInfo")
	userInfo = userInfo.(admin.AdminUser)
	return common.ReturnData{Code: 200, Data: userInfo}
}
