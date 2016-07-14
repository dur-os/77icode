package admin

import (
	"net/http"

	"github.com/dur-os/77icode/Server/common"
	"github.com/dur-os/77icode/Server/model/admin"
	"github.com/zenazn/goji/web"
)

func init() {
	route := &common.Route{Pre: "/admin/*", Controler: &UserController{}}
	route.Handles = append(route.Handles, common.PathToHandle{"/userInfo", "GetUser", common.GET})
	route.Handles = append(route.Handles, common.PathToHandle{"/GetUser1", "GetUser1", common.GET})
	common.RegRoute(route)
}

// UserController is admin user Controller
type UserController struct {
	common.Controller
}

func (controller *UserController) GetUser(c web.C, r *http.Request) common.ReturnData {
	session := controller.GetSession(c)
	userInfo := session.Get("userInfo")
	userInfo = userInfo.(admin.AdminUser)
	return common.ReturnData{Code: 200, Data: userInfo}
}

func (controller *UserController) GetUser1(c web.C, r *http.Request) common.ReturnData {
	return common.ReturnData{Code: 200, Data: map[string]string{"aaaa": "aaaa"}}
}
