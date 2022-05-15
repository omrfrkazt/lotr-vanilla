package main

import (
	"net/http"
)

//crud handlers
const fellowship = "fellowship"

func (app *application) getFellowship(w http.ResponseWriter, r *http.Request) {
	result, err := app.models.GetFellowship(r.Header.Get("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, result, fellowship)
}
