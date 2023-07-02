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

var pallete = []color.Color{color.White}

const (
	blackIndex = 0 // First color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	for i := 1; i < math.MaxUint8; i++ {
		red := uint8(rand.Float64() * float64(256))
		green := uint8(rand.Float64() * float64(256))
		blue := uint8(rand.Float64() * float64(256))
		pallete = append(pallete, color.RGBA{red, green, blue, 0xFF})
	}
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
