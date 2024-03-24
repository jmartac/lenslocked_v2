package controllers

import (
	"net/http"

	"github.com/jmartac/lenslocked_v2/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
