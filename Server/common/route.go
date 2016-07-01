package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/astaxie/beego/session"
	"github.com/zenazn/goji/web"
)

// Controller is http base controller
type Controller struct {
}

// GetSession is web.C get session
func (controller *Controller) GetSession(c web.C) session.Store {
	return c.Env["Session"].(session.Store)
}

type Route struct {
	Pre string
	*Controller
	Path map[string]string
}

//Route reg route
func (application *Application) Route(controller interface{}, route string) interface{} {
	methodValue := reflect.ValueOf(controller).MethodByName(route)
	methodInterface := methodValue.Interface()
	method := methodInterface.(func(c web.C, r *http.Request) ReturnData)
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		returnData := method(c, r)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(returnData); err != nil {
			fmt.Fprintln(w, err)
		}
	}
	return fn
}
