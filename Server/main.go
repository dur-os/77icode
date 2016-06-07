package main

import (
	"flag"
	"net/http"

	"github.com/dur-os/77icode/Server/common"
	"github.com/golang/glog"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	_ "github.com/lib/pq"
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
	// graceful.PostHook(func() {
	// 	application.Close()
	// })
	goji.Serve()
}
