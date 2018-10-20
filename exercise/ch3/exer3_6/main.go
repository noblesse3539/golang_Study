package main

import (
	"os"
	"math/cmplx"
	"image/color"
	"image/png"
	"image"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)

			// Image의 점 (px, py)는 복소수 값 z를 나타낸다.
			img.Set(px, py, color.Gray{uint8((uint16(mandelbrot(z1).Y) + uint16(mandelbrot(z2).Y) + uint16(mandelbrot(z3).Y) + uint16(mandelbrot(z4).Y))/4)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: 오류 무시
}

func mandelbrot(z complex128) color.Gray {
	const iterations = 200
	const contrast = 15
	
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Gray{0}
}