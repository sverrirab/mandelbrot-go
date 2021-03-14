package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

const (
	_ViewWidth  = 640
	_ViewHeight = 480
	_MaxEscape  = 64
)

var indexTemplate *template.Template
var palette []color.RGBA
var escapeColor color.RGBA

func init() {
	var err error
	var indexData string
	indexData = string(MustAsset("resources/index.html"))
	indexTemplate, err = template.New("index.html").Parse(indexData)
	if err != nil {
		panic(err)
	}

	palette = make([]color.RGBA, _MaxEscape)
	for i := 0; i < _MaxEscape-1; i++ {
		palette[i] = color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255}
	}
	escapeColor = color.RGBA{R:0, G:0, B:0, A:0}
}

func escape(c complex128) int {
	z := c
	for i := 0; i < _MaxEscape-1; i++ {
		if cmplx.Abs(z) > 2 {
			return i
		}
		z = cmplx.Pow(z, 2) + c
	}
	return _MaxEscape - 1
}

func generate(imgWidth int, imgHeight int, viewCenter complex128, radius float64) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	zoomWidth := radius * 2
	pixelWidth := zoomWidth / float64(imgWidth)
	pixelHeight := pixelWidth
	viewHeight := (float64(imgHeight) / float64(imgWidth)) * zoomWidth
	left := (real(viewCenter) - (zoomWidth / 2)) + pixelWidth/2
	top := (imag(viewCenter) - (viewHeight / 2)) + pixelHeight/2

	var wgx sync.WaitGroup
	wgx.Add(imgWidth)
	for x := 0; x < imgWidth; x++ {
		go func(xx int) {
			defer wgx.Done()
			for y := 0; y < imgHeight; y++ {
				coord := complex(left+float64(xx)*pixelWidth, top+float64(y)*pixelHeight)
				f := escape(coord)
				if f == _MaxEscape-1 {
					m.Set(xx, y, escapeColor)
				}
				m.Set(xx, y, palette[f])
			}
		}(x)
	}
	wgx.Wait()
	return m
}

func _SafeFloat64(s string, def float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}
	return f
}

func pic(w http.ResponseWriter, r *http.Request) {
	mx := _SafeFloat64(r.FormValue("mx"), 0.0)
	my := _SafeFloat64(r.FormValue("my"), 0.0)
	radius := _SafeFloat64(r.FormValue("radius"), 2.0)

	m := generate(_ViewWidth, _ViewHeight, complex(mx, my), radius)
	w.Header().Set("Content-Type", "image/png")
	err := png.Encode(w, m)
	if err != nil {
		log.Println("png.Encode:", err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	vars := make(map[string]interface{})
	vars["Title"] = "Mandelbrot"
	vars["Height"] = _ViewHeight
	vars["Width"] = _ViewWidth
	err := indexTemplate.Execute(w, vars)
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Listening - open http://localhost:8090/ in browser")
	defer log.Println("Exiting")

	http.HandleFunc("/", index)
	http.HandleFunc("/pic", pic)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
