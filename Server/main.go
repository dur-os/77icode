package main

import (
	"flag"
	"net/http"

	"github.com/dur-os/77icode/Server/common"
	_ "github.com/dur-os/77icode/Server/model/admin"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func main() {
	filename := flag.String("config", "config.ini", "Path to configuration file")
	flag.Parse()
	defer glog.Flush()
	var application = &common.Application{}
	application.Init(filename)

	static := web.New()
	static.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(application.Config.PublicPath))))

	http.Handle("/assets/", static)

	// Apply middleware
	goji.Use(application.ApplyTemplates)
	goji.Use(application.ApplySessions)
	goji.Use(application.ApplyDB)

	goji.Get("/", func(c web.C, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Thanks ok !!!!!!!!!"))
	})

	admin := web.New()
	admin.Use(middleware.SubRouter)
	admin.Get("/", func(c web.C, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1111111111111111"))
	})

	goji.Handle("/admin/*", admin)

	// graceful.PostHook(func() {
	// 	application.Close()
	// })
	goji.Serve()
}
