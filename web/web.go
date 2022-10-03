package web

import (
	"embed"
	"fmt"
	"image"
	"image/jpeg"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kbinani/screenshot"
)

//go:embed static
var FrontendFS embed.FS

func StartWebserver(port int) {
	host := fmt.Sprintf("0.0.0.0:%v", port)

	go func() {
		for {
			UpdateGui()
			time.Sleep(1 * time.Second)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/ws", BuildWebsocket())
	r.HandleFunc("/frame", FrameHandler)
	r.PathPrefix("/").Handler(getStaticFilesHandler(FrontendFS))
	log.Fatal(http.ListenAndServe(host, r))
}

func FrameHandler(w http.ResponseWriter, r *http.Request) {
	if err := jpeg.Encode(w, nextFrame(), nil); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}

//utils
func nextFrame() *image.RGBA {
	screennum := 0
	bounds := screenshot.GetDisplayBounds(screennum)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	return img
}

func getStaticFilesHandler(fefiles embed.FS) http.Handler {
	matches, _ := fs.Glob(fefiles, "static")
	if len(matches) != 1 {
		panic("unable to find frontend build files in FrontendFS")
	}
	feRoot, _ := fs.Sub(fefiles, matches[0])
	buildHandler := http.FileServer(http.FS(feRoot))
	return buildHandler
}
