package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	handler "github.com/jinwoov/hacktober2022/4-handlers"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	mux := mux.NewRouter()
	handler.RequestHandler(mux)

	mux.Use(cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"},
			AllowedHeaders: []string{"Accept", "content-type", "Content-Length"},
		}).Handler)

	log.SetFormatter(&log.JSONFormatter{})
	log.Infoln("Serving the website on 4040")
	log.Errorln(http.ListenAndServe(":4040", handlers.CORS()(mux)))
}
