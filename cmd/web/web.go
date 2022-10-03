package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartWebserver(port int) {
	host := fmt.Sprintf("0.0.0.0:%v", port)

	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	log.Fatal(http.ListenAndServe(host, r))
}
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
