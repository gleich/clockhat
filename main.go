package main

import (
	"time"

	"github.com/Matt-Gleich/logoru"
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
)

func main() {
	for {
		var (
			x   int
			y   int
			fb  = screen.NewFrameBuffer()
			now = time.Now()
		)

		for i := 2; i <= 60; i++ {
			if x == 8 {
				y++
				x = 0
			}
			if y == 8 {
				break
			}

			if i <= now.Minute() {
				fb.SetPixel(x, y, color.New(255, 255, 255))
				continue
			}

			rgbVal := int(4.25 * now.Second())
			fb.SetPixel(rgbVal, rgbVal, rgbVal)

			x++
		}

		err := screen.Draw(fb)
		if err != nil {
			logoru.Error("Failed to update lights;", err)
		}

		time.Sleep(time.Millisecond * 20)
	}
}
