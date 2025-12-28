package main

import (
	"image/png"
	"log"
	"os"

	"github.com/crazy3lf/colorconv"
)

func decoder(path string) (content string, width int, height int, err error) {

	file, err := os.Open(path)
	if err != nil {
		return "", 0, 0, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	hexValues := ""

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			hexValues += string(colorconv.ColorToHex(img.At(x, y)))[2:]
		}
	}

	return hexValues, img.Bounds().Max.X, img.Bounds().Max.Y, nil
}
