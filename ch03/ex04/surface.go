package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // キャンバスの大きさ（画素数）
	cells         = 100                 // 格子のます目の数
	xyRange       = 30.0                // 軸の範囲（-xyrange..+xyrange）
	xyScale       = width / 2 / xyRange // x単位およびy単位当たりの画素数
	zScale        = height * 0.4        // z単位当たりの画素数
	angle         = math.Pi / 6         // x,y軸の角度（=30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type zFunc func(x, y float64) float64

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	svg(w, defaultZFunc)
}

func svg(w io.Writer, fn zFunc) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	min, max := calcMinMax(fn)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, fn)
			bx, by := corner(i, j, fn)
			cx, cy := corner(i, j+1, fn)
			dx, dy := corner(i+1, j+1, fn)
			if isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				color := toColor(i, j, min, max, fn)
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f zFunc) (float64, float64) {
	// ます目(i, j int)のかどの点(x, y)を見つける。
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)

	// (x,y,z)を2-D SVGキャンバス(sx, sy)へ等角的に投影。
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}

func defaultZFunc(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)からの距離
	return math.Sin(r) / r
}

func isValid(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func calcMinMax(fn zFunc) (float64, float64) {
	min := math.Inf(+1)
	max := math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyRange * (float64(i)/cells - 0.5)
			y := xyRange * (float64(j)/cells - 0.5)
			z := fn(x, y)

			if math.IsNaN(z) || math.IsInf(z, 0) {
				continue
			}

			min = math.Min(min, z)
			max = math.Max(max, z)
		}
	}
	return min, max
}

func toColor(i, j int, min, max float64, fn zFunc) string {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)
	z := fn(x, y)
	coeff := (z - min) / (max - min)
	red := uint8(255 * coeff)
	blue := uint8(255 * (1 - coeff))
	return fmt.Sprintf("#%02x00%02x", red, blue)
}
