package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	qs := r.FormValue("cycles")
	if qs == "" {
		lissajous(w, 5)
		return
	}

	cycles, err := strconv.Atoi(qs)
	if err != nil {
		fmt.Fprintf(w, "Internal Server Error: %s\n", err)
	} else {
		lissajous(w, cycles)
	}
}

func lissajous(out io.Writer, cycles int) {
	rand.Seed(time.Now().UTC().UnixNano())
	var palette = []color.Color{
		color.Black,
		color.RGBA{0x00, 0xff, 0x00, 0xff}, // 緑
		color.RGBA{0xff, 0x00, 0x00, 0xff}, // 赤
		color.RGBA{0x00, 0x00, 0xff, 0xff}, // 青
	}
	const (
		res              = 0.001 // 回転の分解能
		size             = 100   // 画像キャンバスは[-size..+size]の範囲を扱う
		nframes          = 64    // アニメーションフレーム数
		delay            = 8     // 10ms1単位でのフレーム間の遅延
		firstColorIndex  = 0     // パレットの最初の色
		secondColorIndex = 1     // パレットの次の色
		thirdColorIndex  = 2
		fourthColorIndex = 3
	)

	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), secondColorIndex)
			img.SetColorIndex(size+int(x*size), size+int(y*size), thirdColorIndex)
			img.SetColorIndex(size+int(x*size+0.25), size+int(y*size+0.75), fourthColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
