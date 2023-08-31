package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
)

func main() {
	file := flag.String("file", "", "Filename to output!")
	flag.Parse()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var (
		red   = uRandInt8()
		green = uRandInt8()
		blue  = uRandInt8()
		// alpha = uRandInt8()
	)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z, red, green, blue))
		}
	}
	if *file == "" {
		png.Encode(os.Stdout, img)
	} else {
		new_file, err := os.Create(*file)
		if err != nil {
			log.Fatal("Ex3.5: File creation error: ", err)
		}
		png.Encode(new_file, img)
	}
}

func uRandInt8() uint8 {
	return uint8(rand.Intn(math.MaxUint8 + 1))
}

func mandelbrot(z complex128, red, green, blue uint8) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				red - constrast*n,
				green - constrast*n,
				blue - constrast*n,
				255 - constrast*n,
			}
		}
	}
	return color.White
}
