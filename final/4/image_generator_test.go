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

func TestGenerateImgColorMatrix1(t *testing.T) {
	height := 20
	width := 20
	steps := 40

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	color_data := GenerateImgColorMatrix(data, 255, 255, 255, 1, 1, 1)

	matrix_r := color_data[0]
	fmt.Println("matrix_r")
	for i := 0; i < len(matrix_r); i++ {
		fmt.Println(matrix_r[i])
	}
	matrix_g := color_data[1]
	fmt.Println("matrix_g")
	for i := 0; i < len(matrix_g); i++ {
		fmt.Println(matrix_r[i])
	}
	matrix_b := color_data[2]
	fmt.Println("matrix_b")
	for i := 0; i < len(matrix_b); i++ {
		fmt.Println(matrix_r[i])
	}

}

func TestGenerateImgColorMatrix2(t *testing.T) {
	height := 40
	width := 20
	steps := 40

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	color_data := GenerateImgColorMatrix(data, 255, 255, 255, 190, 100, 120)

	matrix_r := color_data[0]
	fmt.Println("matrix_r")
	for i := 0; i < len(matrix_r); i++ {
		fmt.Println(matrix_r[i])
	}
	matrix_g := color_data[1]
	fmt.Println("matrix_g")
	for i := 0; i < len(matrix_g); i++ {
		fmt.Println(matrix_r[i])
	}
	matrix_b := color_data[2]
	fmt.Println("matrix_b")
	for i := 0; i < len(matrix_b); i++ {
		fmt.Println(matrix_r[i])
	}
}

func TestGenerateImgColorMatrix3(t *testing.T) {
	height := 20
	width := 20
	steps := 10000

	data := GenerateImgMatrix(height, width, steps)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	color_data := GenerateImgColorMatrix(data, 255, 255, 255, 55, 55, 55)

	matrix_r := color_data[0]
	fmt.Println("matrix_r")
	for i := 0; i < len(matrix_r); i++ {
		fmt.Println(matrix_r[i])
	}
	matrix_g := color_data[1]
	fmt.Println("matrix_g")
	for i := 0; i < len(matrix_g); i++ {
		fmt.Println(matrix_r[i])
	}
	matrix_b := color_data[2]
	fmt.Println("matrix_b")
	for i := 0; i < len(matrix_b); i++ {
		fmt.Println(matrix_r[i])
	}
}
