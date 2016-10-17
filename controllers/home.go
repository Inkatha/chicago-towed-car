package controllers

import (
	"chicago-towed-car/model"
	"chicago-towed-car/viewmodels"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (home *homeController) homePageHandler(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	w.Header().Add("Context-Type", "text/html")

	switch req.Method {
	case "GET":
		home.template.Execute(w, vm)
	case "POST":
		req.ParseForm()
		plateNumber := req.FormValue("plateNumber")
		plateNumber = strings.Trim(plateNumber, " ")
		plateNumber = strings.ToUpper(plateNumber)

		m := model.SearchByLicensePlate(plateNumber)

		fmt.Println(m.Plate)
		home.template.Execute(w, m)
	default:
		log.Fatal("Could not find appropriate route for home.html")
	}
}
