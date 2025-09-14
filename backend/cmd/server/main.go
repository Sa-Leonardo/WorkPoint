package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Funcionando"))
    }).Methods("GET")

    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
