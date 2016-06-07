package common

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// Makes sure templates are stored in the context
func (application *Application) ApplyTemplates(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["Template"] = application.Template
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (application *Application) ApplyDB(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["DB"] = application.DB
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// Makes sure controllers can have access to session
func (application *Application) ApplySessions(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := application.SessionManager.SessionStart(w, r)
		defer session.SessionRelease(w)
		c.Env["Session"] = session
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
