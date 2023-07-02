// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{}

const (
	blackIndex = 0 // First color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	for red := 0; red < math.MaxUint8; red++ {
		for green := 0; green < math.MaxUint8; green++ {
			for blue := 0; blue < math.MaxUint8; blue++ {
				pallete = append(pallete, color.RGBA{uint8(red), uint8(green), uint8(blue), 0xff})
			}
		}
	}
	// fmt.Printf("len(pallete): %v\n", len(pallete))
	// index := uint8(rand.Float64() * float64(len(pallete)))
	// fmt.Printf("index: %v\n", index)
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas corvers [-size..+size]
		nframes = 64    // number of animations frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.00 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			index := uint8(rand.Float64() * float64(len(pallete)))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size*0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
