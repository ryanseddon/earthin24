package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"os"
	"time"
)

func main() {
	tileRange := []int{0, 1}
	img := image.NewRGBA(image.Rect(0, 0, 1100, 1100))

	for _, x := range tileRange {
		for _, y := range tileRange {
			url := pathfor(x, y)
			resp, err := http.Get(url)

			defer resp.Body.Close()

			tiledata := resp.Body
			tile, _, err := image.Decode(tiledata)
			if err != nil {
				fmt.Println(err)
			}
			_ = "breakpoint"
			r := image.Rect(1100, 1100, 550*x, 550*y)
			draw.Draw(img, r, tile, image.Point{0, 0}, draw.Over)
		}
	}

	f, _ := os.Create("earth.png")
	defer f.Close()
	png.Encode(f, img)
}

func pathfor(x, y int) string {
	url := "http://himawari8-dl.nict.go.jp/himawari8/img/D531106/2d/550"
	// Images are taken in 10 min intervals and usually lag by ~20mins
	then := time.Now().Add(-20 * time.Minute).UTC()
	minute := then.Minute() - then.Minute()%10

	return fmt.Sprintf("%s/%d/%02d/%02d/%02d%02d00_%d_%d.png", url, then.Year(), then.Month(), then.Day(), then.Hour(), minute, x, y)
}

//func getTile(x, y int) {
//resp, err := http.Get(pathfor("http://himawari8.nict.go.jp/img/D531106/2d/550", x, y))

//defer resp.Body.Close()

//return resp.Body
//}
