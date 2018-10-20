package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 픽셀 단위 캔버스 크기
	cells         = 100                 // 격자 셀의 숫자
	xyrange       = 30.0                // 축 범위( -xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x나 y의 단위 픽셀
	zscale        = height * 0.4        // z 단위 픽셀
	angle         = math.Pi / 6         // x, y축의 각도 (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height) // http://www.w3.org/TR/SVG
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if !(ok1 && ok2 && ok3 && ok4) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// (i,j) 셀 코너에서 (x,y) 점 찾기
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 표면 높이 z 연산
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}
	// (x,y,z)를 3차원 SVG 평면 (sx,sy)에 등각 투영시킴
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}
func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // (0,0)에서의 거리
	if math.IsInf(r,0) || math.IsNaN(r) {
		return 0, false
	}
	return math.Sin(r) / r, true
}
