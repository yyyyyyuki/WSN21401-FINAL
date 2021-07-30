package main

import (
	// 	"fmt"
	"image/color"
	"testing"
)

func TestDrawImage01(t *testing.T) {
	height := 20
	width := 20
	steps := 40

	data := GenerateImgMatrix(height, width, steps)

	scale := 10
	foregroundColor := color.White
	ForegroundComplementaryColor := color.White
	backgroundColor := color.Black

	DrawImage(data, scale, foregroundColor, ForegroundComplementaryColor, backgroundColor)
}
