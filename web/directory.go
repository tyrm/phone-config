package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"phone-config/models"
)

type DirectoryTemplate struct {
	Lines              *[]models.Line
	ProvisioningServer string
}


func GetDirectory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logger.Errorf("got vars: %#v", vars)

	tmplVars := DirectoryTemplate{}

	templateName := fmt.Sprintf("%s_directory", vars["vendor"])
	err := templates.ExecuteTemplate(w, templateName, tmplVars)
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