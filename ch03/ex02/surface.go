package main

// mathパッケージの別の関数で可視化を試してみなさい。
// みなさんは、鶏卵の箱、モーグルのこぶ、乗馬用の鞍などを生成できますか。

import (
	"fmt"
	"io"
	"math"
	"os"
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

func eggbox(x, y float64) float64 {
	return 0.15 * (math.Cos(x) + math.Cos(y))
}

func main() {
	svg(os.Stdout, eggbox)
}

func svg(w io.Writer, fn zFunc) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, fn)
			bx, by := corner(i, j, fn)
			cx, cy := corner(i, j+1, fn)
			dx, dy := corner(i+1, j+1, fn)
			if isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func isValid(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
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
