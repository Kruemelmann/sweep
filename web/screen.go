package web

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func NextFrame() {
	screennum := 0
	bounds := screenshot.GetDisplayBounds(screennum)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d.png", screennum, bounds.Dx(), bounds.Dy())
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
	fmt.Printf("#%d : %v \"%s\"\n", screennum, bounds, fileName)
}
