package controllers

import (
	"chicago-towed-car/model"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (home *homeController) homePageHandler(w http.ResponseWriter, req *http.Request) {
	m := model.NewCar()

	w.Header().Add("Context-Type", "text/html")

	switch req.Method {
	case "GET":
		home.template.Execute(w, m)
	case "POST":
		req.ParseForm()
		plateNumber := req.FormValue("plateNumber")
		plateNumber = strings.Trim(plateNumber, " ")
		plateNumber = strings.ToUpper(plateNumber)

		m := model.SearchByLicensePlate(plateNumber)

		home.template.Execute(w, m)
	default:
		log.Fatal("Could not find appropriate route for home.html")
	}
}
