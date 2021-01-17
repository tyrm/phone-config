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
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			logger.Tracef("  %s: %s", name, value)
		}
	}


	vars := mux.Vars(r)
	mode := "phone"

	switch vars["mac"] {
	case "F0V0X5U00000":
		mode = "generic"
	}

	// Get Phone

	phone, err := models.GetPhone(vars["mac"], vars["vendor"], vars["model"])
	if err != nil {
		msg := fmt.Sprintf("could not get phone: %s", err.Error())
		logger.Errorf(msg)
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	if phone == nil && mode == "phone" {
		http.Error(w, "phone not found", http.StatusNotFound)
		return
	}

	// Make Template
	tmplVars := ConfigTemplate{}
	tmplVars.ProvisioningServer = extHostname

	if mode == "phone" {
		tmplVars.Lines, err = phone.GetLines()
		if err != nil {
			msg := fmt.Sprintf("could not get lines: %s", err.Error())
			logger.Errorf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	}

	// Select Template
	templateName := fmt.Sprintf("%s_%s_%s_%s", vars["vendor"], vars["model"], vars["suffix"], mode)

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
