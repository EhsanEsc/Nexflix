package controller

import (
	"net/http"
	"text/template"
	"viewmodels"
)

type homeController struct {
	template *template.Template
}

func (hc *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content Type", "text/html")
	hc.template.Execute(w, vm)
}
