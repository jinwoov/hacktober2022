package hacktoberfest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RequestHandler(mux *mux.Router) {
	mux.HandleFunc("/healthz", getHealthCheck).Methods("GET")
	mux.HandleFunc("/hello", getHelloWorld).Methods("GET")
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	helloWorld := `Hello world!, remember Kindness is not what you do, but who you are!`

	w.WriteHeader(http.StatusOK)
	
	if err := json.NewEncoder(w).Encode(helloWorld); err != nil {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
	}
}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	var healthOjb struct {
		Status string `json:"status"`
	}

	healthOjb.Status = "UP"

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(healthOjb); err != nil {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
	}
}
