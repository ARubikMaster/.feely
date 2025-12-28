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

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			hexValues = hexValues + string(colorconv.ColorToHex(img.At(x, y)))[2:]
		}
	}

	return hexValues, img.Bounds().Max.X, img.Bounds().Max.Y, nil
}
