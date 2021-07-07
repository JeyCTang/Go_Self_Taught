package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	// size of picture
	const size = 300
	// create the gray picture according to the size
	pic := image.NewGray(image.Rect(0, 0, size, size))

	// traverse every pixel
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// fill with white color
			pic.SetGray(x, y, color.Gray{Y: 255})
		}
	}

	// create X axis from 0
	for x := 0; x < size; x++ {
		// regular the size of Sin from 0 - 2Pi
		s := float64(x) * 2 * 2 * math.Pi / size
		// the range of sin is half of the size, we will move down it and turnover
		y := size / 2 - math.Sin(s) * size / 2

		// draw the sin curve with black color
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// create file
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// write picture into PNG file
	png.Encode(file, pic)

	// close file
	file.Close()
}