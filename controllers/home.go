package controllers

import (
	"net/http"
	"text/template"
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
