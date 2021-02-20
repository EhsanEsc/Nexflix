package controller

import (
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
)

type filmsController struct {
	template *template.Template
}

func (fc *filmsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetFilms()
	if req.FormValue("film_score") == "" {
		vm.Filtered_Film_score = 0
	} else {
		vm.Filtered_Film_score, _ = strconv.Atoi(req.FormValue("film_score"))
	}

	w.Header().Add("Content Type", "text/html")
	fc.template.Execute(w, vm)
}
