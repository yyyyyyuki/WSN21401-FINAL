package main

import (
	"fmt"
	"testing"
)

func TestGetRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(get_random(1, 100))
	}
}

func TestGenerateImgMatrix1(t *testing.T) {
	height := 20
	width := 20
	steps := 40

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func TestGenerateImgMatrix2(t *testing.T) {
	height := 10
	width := 20
	steps := 120

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func TestGenerateImgMatrix3(t *testing.T) {
	height := 40
	width := 20
	steps := 40

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func TestGenerateImgMatrix4(t *testing.T) {
	height := 20
	width := 20
	steps := 10000

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}
