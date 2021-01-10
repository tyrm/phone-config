package web

import (
	"github.com/gorilla/mux"
	"github.com/juju/loggo"
	"github.com/markbates/pkger"
	"net/http"
	"phone-config/config"
	"text/template"
	"time"
)

var (
	logger     *loggo.Logger
	templates  *template.Template

	extHostname string
)

func Init(conf *config.Config) error {
	newLogger := loggo.GetLogger("web")
	logger = &newLogger

	extHostname = conf.ExtHostname

	// Load Templates
	templateDir := pkger.Include("/web/templates")
	t, err := compileTemplates(templateDir)
	if err != nil {
		return err
	}
	templates = t

	// Setup Router
	r := mux.NewRouter()
	r.Use(Middleware)

	// Fanvil
	r.HandleFunc("/cnf/{vendor}/{model}/{mac}.{suffix}", GetConfig).Methods("GET")
	r.HandleFunc("/directory/{vendor}/home.{suffix}", GetDirectory).Methods("GET")

	go func() {
		srv := &http.Server{
			Handler:      r,
			Addr:         ":5000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		err := srv.ListenAndServe()
		if err != nil {
			logger.Errorf("Could not start web server %s", err.Error())
		}
	}()

	return nil
}