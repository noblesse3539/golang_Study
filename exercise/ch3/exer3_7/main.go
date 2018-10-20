package main

import (
	"math/cmplx"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, Newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func Newton(z complex128) color.Color {
	const iterator = 200
	const contrast = 15

	for n := uint8(0); n < iterator; n++ {
		z0 := z
		z = z - f(z)/difF(z)	//Newton-Raphson
		/**
		 *	더 좋은 방법 : 위의 식을 좀더 간결하게 하면 z -= (z - 1/(z*z*z))/4 이다. 
		*/
		if cmplx.Abs(f(z) - f(z0)) < 0.0001 {
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Gray{0}
}

func f(z complex128) complex128 {
	return z*z*z*z - 1
}
func difF(z complex128) complex128 {
	return 4 * z * z * z
}
