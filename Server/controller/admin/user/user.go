package user

import (
	"net/http"

	"github.com/dur-os/77icode/Server/common"
	"github.com/dur-os/77icode/Server/model/admin/user"
	"github.com/zenazn/goji/web"
)

func init() {
	route := &common.Route{Pre: "/admin/*", Controler: &UserController{}}
	route.Handles = append(route.Handles, common.PathToHandle{Path: "/userInfo", Handle: "GetUser", Method: common.GET})
	route.Handles = append(route.Handles, common.PathToHandle{Path: "/GetUser1", Handle: "GetUser1", Method: common.GET})
	route.Handles = append(route.Handles, common.PathToHandle{Path: "/login", Handle: "Login", Method: common.POST})
	common.RegRoute(route)
}

// UserController is admin user Controller
type UserController struct {
	common.Controller
}

func (controller *UserController) Login(c web.C, r *http.Request) common.ReturnData {
	err := r.ParseForm()
	if err != nil {
		return common.ReturnData{Code: 502, Msg: "参数错误"}
	}
	userName := r.Form.Get("userName")
	passWord := r.Form.Get("passWord")
	user := &user.AdminUser{Username: userName, Password: passWord}
	if user.Login(controller.GetDB(c)) {
		controller.GetSession(c).Set("userInfo", user)
		return common.ReturnData{Code: 200, Msg: "登陆成功", Data: user}
	}
	return common.ReturnData{Code: 302, Msg: "用户名或密码错误"}

}

func (controller *UserController) GetUser(c web.C, r *http.Request) common.ReturnData {
	session := controller.GetSession(c)
	userInfo := session.Get("userInfo")
	userInfo = userInfo.(user.AdminUser)
	return common.ReturnData{Code: 200, Data: userInfo}
}

func (controller *UserController) GetUser1(c web.C, r *http.Request) common.ReturnData {
	return common.ReturnData{Code: 200, Data: map[string]string{"name": "aaaa", "userName": "ccccc"}}
}
