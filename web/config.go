package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"phone-config/models"
)

type ConfigTemplate struct {
	Lines              *[]models.Line
	ProvisioningServer string
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logger.Tracef("got vars: %#v", vars)

	// Get Phone
	phone, err := models.GetPhone(vars["mac"], vars["vendor"], vars["model"])
	if err != nil {
		msg := fmt.Sprintf("could not get phone: %s", err.Error())
		logger.Errorf(msg)
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	if phone == nil {
		http.Error(w, "phone not found", http.StatusNotFound)
		return
	}

	// Make Template
	tmplVars := ConfigTemplate{}
	tmplVars.ProvisioningServer = extHostname

	tmplVars.Lines, err = phone.GetLines()
	if err != nil {
		msg := fmt.Sprintf("could not get lines: %s", err.Error())
		logger.Errorf(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	templateName := fmt.Sprintf("%s_%s_%s", vars["vendor"], vars["model"], vars["suffix"])
	err = templates.ExecuteTemplate(w, templateName, tmplVars)
	if err != nil {
		msg := fmt.Sprintf("could not render template: %s", err.Error())
		logger.Errorf(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if vars["suffix"] == "xml" {
		w.Header().Add("Content-Type", "text/xml")
	}
}
