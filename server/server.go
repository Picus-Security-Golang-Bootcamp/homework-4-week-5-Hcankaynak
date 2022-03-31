package server

import (
	"log"
	"net/http"
)

type App struct {
	client *http.Client
}

func StartServer() {
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/send", withQueryParam)
	http.HandleFunc("/form", withFormValue)

	log.Println(http.ListenAndServe("localhost:8080", nil))
}

func ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("READY"))
}

func withQueryParam(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param))

}

func withFormValue(w http.ResponseWriter, r *http.Request) {
	param1 := r.FormValue("param1")
	param2 := r.FormValue("param2")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param1 + param2))

}
