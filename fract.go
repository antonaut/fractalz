package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var (
	FILENAME string
)

const (
	WIDTH  = 400
	HEIGHT = 400
	MAX_N  = 80
)

// Takes a complex number c
// Returns a number n such that 0 <= n < 1.0f
// Corresponding to c:s existance in the Mandelbrot set
// with respect to the maximum number of iterations, MAX_N
// Returns 1 if MAX_N is reached
func mandel(c complex128) float64 {
	z := complex(0, 0)
	for i := 0; i < MAX_N; i++ {
		if cmplx.Abs(z) > 2 {
			return float64(i-1) / MAX_N
		}
		z = z*z + c
	}

	return 1
}

// Draws the Mandelbrot fractal to FILENAME
func main() {
	b := image.Rect(0, 0, WIDTH, HEIGHT)
	img := image.NewRGBA(b)

	flag.StringVar(&FILENAME, "f", "tmp.png", "The name of the file to write to.")
	flag.Parse()

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			zx := 2.0
			zy := 2.0
			xf := float64(x)/WIDTH*zx - (zx/2.0 + 0.5)
			yf := float64(y)/HEIGHT*zy - (zy / 2.0)
			c := complex(xf, yf)
			//fmt.Printf("(%g, %g) ", xf, yf)
			pix_value := int(mandel(c) * 255)

			// Get a seemingly random color value to put on the pixels
			color_value := color.RGBA{uint8(pix_value), uint8(2 * pix_value % 255), uint8(3 * pix_value % 255), 255}
			if pix_value == 255 {
				color_value = color.RGBA{0, 0, 0, 255} // Black
			}
			img.Set(x, y, color_value)
		}

	}

	file, err := os.Create(FILENAME)
	defer file.Close()

	if err != nil || file == nil {
		file, err = os.Open(FILENAME)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s\n", err)
			return
		}
	}

	err = png.Encode(file, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding image: %s\n", err)
		return
	}
}
