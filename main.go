package main

import (
	"fmt"
	"image/color"
)

func main() {
	height := 20
	width := 20
	steps := 40
	scale := 10
	pattern := 0
	fc := color.White
	fcc := fc
	bc := color.Black

	fmt.Println("Step 1: Input of height, width, steps and scale of avator (default is 20, 20, 20, 10).")
	fmt.Println("Please enter height and width (0 is default):")
	t1 := 0
	t2 := 0
	fmt.Scan(&t1, &t2)
	if t1 > 0 {
		height = t1
	}
	if t2 > 0 {
		width = t2
	}
	fmt.Print("Please enter steps (0 is default):")
	t1 = 0
	fmt.Scan(&t1)
	if t1 > 0 {
		steps = t1
	}
	fmt.Print("Please enter scale (0 is default):")
	t1 = 0
	fmt.Scan(&t1)
	if t1 > 0 {
		scale = t1
	}

	data := GenerateImgMatrix(height, width, steps)

	fmt.Println("Step 2: Select color pattern.")
	fmt.Print("Please select color pattern. 0 is black background, 1 is white background:")
	t1 = 0
	fmt.Scan(&t1)
	if t1 == 1 {
		pattern = 1
	}
	if pattern == 1 {
		fc = bc
		bc = fcc
		fcc = fc
	}

	DrawImage(data, scale, fc, fcc, bc)

	fmt.Println("Done!")
}
