package main

import (
	"fmt"
	"math"
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

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
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

func f(x, y float64) float64 {
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
