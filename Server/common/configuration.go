package common

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego/config"
)

const (
	PUBLIC_PATH   string = "public"
	TEMPLATE_PATH string = "views"
)

//Configuration is app all config
type Configuration struct {
	PublicPath   string
	TemplatePath string
	Database     *ConfigurationDatabase
	Session      *ConfigurationSession
}

// ConfigurationDatabase is app use database config
type ConfigurationDatabase struct {
	Database     string
	URL          string
	MaxOpenConns int
	MaxIdleConns int
}

//ConfigurationSession is session message config
type ConfigurationSession struct {
	SessionEngine     string `json:"-"`
	CookieName        string `json:"cookieName"`
	EnableSetCookie   bool   `json:"enableSetCookie,omitempty"`
	Gclifetime        int    `json:"gclifetime"`
	MaxLifetime       int    `json:"maxLifetime"`
	Secure            bool   `json:"secure"`
	SessionIDHashFunc string `json:"sessionIDHashFunc"`
	SessionIDHashKey  string `json:"sessionIDHashKey"`
	CookieLifeTime    int    `json:"cookieLifeTime"`
	ProviderConfig    string `json:"providerConfig"`
}

//Load is load config file
func (configuration *Configuration) Load(filename string) (err error) {
	iniconf, err := config.NewConfig("ini", filename)
	if err != nil {
		return
	}
	configuration.PublicPath = iniconf.DefaultString("general::public_path", PUBLIC_PATH)
	configuration.TemplatePath = iniconf.DefaultString("general::template_path", TEMPLATE_PATH)
	database := &ConfigurationDatabase{}
	database.Database = iniconf.DefaultString("database::database", "postgres")
	database.URL = iniconf.DefaultString("database::url", "postgres://pqgotest:@localhost:5432/postgres?sslmode=verify-full")
	database.MaxOpenConns = iniconf.DefaultInt("database::max_open_conns", 50)
	database.MaxIdleConns = iniconf.DefaultInt("database::max_idle_conns", 10)
	configuration.Database = database
	sessionConf := &ConfigurationSession{}
	sessionConf.SessionEngine = iniconf.DefaultString("session::engine", "memory")
	sessionConf.CookieName = iniconf.DefaultString("session::cookieName", "gosessionid")
	sessionConf.EnableSetCookie = iniconf.DefaultBool("session::enableSetCookie", true)
	sessionConf.Gclifetime = iniconf.DefaultInt("session::gclifetime", 3600)
	sessionConf.MaxLifetime = iniconf.DefaultInt("session::maxLifetime", 3600)
	sessionConf.Secure = iniconf.DefaultBool("session::secure", false)
	sessionConf.SessionIDHashFunc = iniconf.DefaultString("session::sessionIDHashFunc", "sha1")
	sessionConf.SessionIDHashKey = iniconf.DefaultString("session::sessionIDHashKey", "sessionkey")
	sessionConf.CookieLifeTime = iniconf.DefaultInt("session::cookieLifeTime", 0)
	sessionConf.ProviderConfig = iniconf.DefaultString("session::providerConfig", "")
	configuration.Session = sessionConf
	return
}

// ToJsonString is ConfigurationSession convert to json config
func (sessionConf *ConfigurationSession) ToJsonString() string {
	b, err := json.Marshal(sessionConf)
	if err != nil {
		return ""
	}
	return strings.Replace(string(b), "enableSetCookie", "enableSetCookie,omitempty", -1)
}
