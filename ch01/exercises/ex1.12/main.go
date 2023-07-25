// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 0xff}, color.RGBA{0x00, 0xf0, 0x08, 0xff}}

var parameters = map[string]int{"cycles": 5, "size": 100, "nframes": 64, "delay": 8}

const (
	blackIndex = iota
	greenIndex
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parseQuery(w, r)
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func parseQuery(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k := range parameters {
		n, err := strconv.ParseUint(r.Form.Get(k), 10, 0)
		if err == nil {
			parameters[k] = int(n)
		}
	}
}

func lissajous(out io.Writer) {
	var (
		cycles  = parameters["cycles"]  // number of complete x oscillator revolutions
		res     = 0.001                 // angular resolution
		size    = parameters["size"]    // image canvas covers [-size..+size]
		nframes = parameters["nframes"] // number of animation frames
		delay   = parameters["delay"]   // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
