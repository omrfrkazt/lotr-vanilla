package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var app application
	app.init()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
