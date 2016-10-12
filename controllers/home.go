package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (home *homeController) apiSearch(w http.ResponseWriter, req *http.Request) {
	m := model.GetTowedCarData()
	fmt.Println(m)
	w.Header().Add("Context-Type", "application/json")

	data, err := json.Marshal(m)

	if err != nil {
		w.WriteHeader(404)
	}
	w.Write(data)
}
