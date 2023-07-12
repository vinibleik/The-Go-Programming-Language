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

var pallete = []color.Color{color.White}

var defaultValues = make(map[string]int)

func main() {
	populatePallete()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parseQuery(w, r)
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:5500", nil))
}

func populatePallete() {
	for i := 1; i < math.MaxUint8; i++ {
		red := uint8(rand.Float64() * float64(256))
		green := uint8(rand.Float64() * float64(256))
		blue := uint8(rand.Float64() * float64(256))
		pallete = append(pallete, color.RGBA{red, green, blue, 0xFF})
	}
}

func parseQuery(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	defaultValues["cycles"] = 5
	defaultValues["size"] = 100
	defaultValues["nframes"] = 64
	defaultValues["delay"] = 8

	for param, _ := range defaultValues {
		if formData := r.Form.Get(param); formData != "" {
			value, err := strconv.Atoi(formData)
			if err != nil {
				log.Print(err)
			}
			defaultValues[param] = value
		}
	}
}

func lissajous(out io.Writer, r *http.Request) {
	// var (
	// 	cycles  = 5     // number of complete x oscillator revolutions
	// 	res     = 0.001 // angular resolution
	// 	size    = 100   // image canvas corvers [-size..+size]
	// 	nframes = 64    // number of animations frames
	// 	delay   = 8     // delay between frames in 10ms units
	// )
	var (
		cycles  = defaultValues["cycles"]
		res     = 0.001
		size    = defaultValues["size"]
		nframes = defaultValues["nframes"]
		delay   = defaultValues["delay"]
	)
	freq := rand.Float64() * 3.00 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			index := uint8(rand.Float64() * float64(len(pallete)))
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)*0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
