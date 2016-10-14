package controllers

import (
	"net/http"
	"strings"
	"text/template"
	"towed-car-locator/model"
	"towed-car-locator/viewmodels"
)

type homeController struct {
	template *template.Template
}

func (home *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	w.Header().Add("Context-Type", "text/html")
	home.template.Execute(w, vm)
}

func (home *homeController) licensePlateSearch(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	plateNumber := req.FormValue("plateNumber")
	plateNumber = strings.Trim(plateNumber, " ")
	plateNumber = strings.ToUpper(plateNumber)

	m := model.SearchByLicensePlate(plateNumber)

	home.template.Execute(w, m)
}
