package web

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static
var FrontendFS embed.FS

func StartWebserver(port int) {
	host := fmt.Sprintf("0.0.0.0:%v", port)

	r := mux.NewRouter()
	r.HandleFunc("/ws", BuildWebsocket())
	r.PathPrefix("/").Handler(getStaticFilesHandler(FrontendFS))
	log.Fatal(http.ListenAndServe(host, r))
}

//TODO
//r.HandleFunc("/screencount", ScreenCountHandler)
//func ScreenCountHandler(w http.ResponseWriter, r *http.Request) {
//w.Write([]byte("Gorilla!\n"))
//}

//utils
func getStaticFilesHandler(fefiles embed.FS) http.Handler {
	matches, _ := fs.Glob(fefiles, "static")
	if len(matches) != 1 {
		panic("unable to find frontend build files in FrontendFS")
	}
	feRoot, _ := fs.Sub(fefiles, matches[0])
	buildHandler := http.FileServer(http.FS(feRoot))
	return buildHandler
}
