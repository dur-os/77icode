package common

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego/session"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Application app config and common
type Application struct {
	Config         *Configuration
	Template       *template.Template
	SessionManager *session.Manager
	DB             *gorm.DB
}

// Init is app init
func (application *Application) Init(filename *string) {
	application.Config = &Configuration{}
	err := application.Config.Load(*filename)
	if err != nil {
		glog.Fatalf("ini config load failed: %s\n", err)
	}

	application.DB, err = gorm.Open(application.Config.Database.Database, application.Config.Database.URL)
	if err != nil {
		glog.Fatalf("init DB failed: %s\n", err)
	}
	application.DB.DB().SetMaxOpenConns(application.Config.Database.MaxOpenConns)
	application.DB.DB().SetMaxIdleConns(application.Config.Database.MaxIdleConns)
	AutoMigrate(application.DB)

	application.SessionManager, err = session.NewManager(application.Config.Session.SessionEngine,
		application.Config.Session.ToJsonString())
	if err != nil {
		glog.Fatalf("init session failed: %s\n", err)

	}
	go application.SessionManager.GC()

	application.loadTemplates()
}

func (application *Application) loadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	err := filepath.Walk(application.Config.TemplatePath, fn)

	if err != nil {
		return err
	}

	application.Template = template.Must(template.ParseFiles(templates...))
	return nil
}
