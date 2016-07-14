package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/astaxie/beego/session"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

var regRoutes []*Route

var pathMux map[string]*web.Mux = make(map[string]*web.Mux)

// Controller is http base controller
type Controller struct {
}

// GetSession is web.C get session
func (controller *Controller) GetSession(c web.C) session.Store {
	return c.Env["Session"].(session.Store)
}

//Route is web app route info
type Route struct {
	Pre       string
	Controler interface{}
	Handles   []PathToHandle
}

// PathToHandle is path to Handle
type PathToHandle struct {
	Path   string
	Handle string
	Method string
}

// RegRoute Register route
func RegRoute(route *Route) {
	regRoutes = append(regRoutes, route)
}

func commRoute(controller interface{}, route string) interface{} {
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

//Init is route info init
func Init() {
	fmt.Println(len(regRoutes))
	for _, route := range regRoutes {
		var tempMux *web.Mux
		ok := false
		if tempMux, ok = pathMux[route.Pre]; !ok {
			tempMux = web.New()
			tempMux.Use(middleware.SubRouter)
			goji.Handle(route.Pre, tempMux)
			pathMux[route.Pre] = tempMux
		}
		for _, h := range route.Handles {
			switch h.Method {
			case CONNECT:
				tempMux.Connect(h.Path, commRoute(route.Controler, h.Handle))
			case DELETE:
				tempMux.Delete(h.Path, commRoute(route.Controler, h.Handle))
			case GET:
				tempMux.Get(h.Path, commRoute(route.Controler, h.Handle))
			case HEAD:
				tempMux.Handle(h.Path, commRoute(route.Controler, h.Handle))
			case OPTIONS:
				tempMux.Options(h.Path, commRoute(route.Controler, h.Handle))
			case PATCH:
				tempMux.Patch(h.Path, commRoute(route.Controler, h.Handle))
			case PUT:
				tempMux.Put(h.Path, commRoute(route.Controler, h.Handle))
			case TRACE:
				tempMux.Trace(h.Path, commRoute(route.Controler, h.Handle))
			case POST:
				tempMux.Post(h.Path, commRoute(route.Controler, h.Handle))

			default:
				tempMux.Handle(h.Path, commRoute(route.Controler, h.Handle))
			}
		}
	}
}
