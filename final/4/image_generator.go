package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateImgMatrix(height, width, steps int) [][]int {
	data := make([][]int, height)
	for i := range data {
		data[i] = make([]int, width)
	}

	min_x := 1
	max_x := height/2 - 1
	min_y := 1
	max_y := width - 2
	random_y := get_random(min_y, max_y)

	pos := Point{X: max_x, Y: random_y}

	data[pos.X][pos.Y] = 1

	for i := 0; i < steps; i++ {
		for {
			potential_direction := RandomDirection()
			potential_position, err := MovePoint(pos, potential_direction, min_x, min_y, max_x, max_y)
			if err != nil {
				continue
			}

			pos = potential_position
			break
		}

		data[pos.X][pos.Y] = 1
	}

	for i := 0; i < len(data)/2; i++ {
		data[len(data)-i-1] = data[i]
	}

	return data
}

func get_random(from, to int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(to-from) + from
}

func GenerateImgColorMatrix(data [][]int, br, bg, bb, fr, fg, fb int) [][][]int {
	height := len(data)
	width := len(data[0])

	fmt.Println(height)
	fmt.Println(width)

	color_data := make([][][]int, 3)
	for i := range color_data {
		color_data[i] = make([][]int, height)
		for j := range color_data[i] {
			color_data[i][j] = make([]int, width)
		}
	}

	for x := range color_data[0] {
		for y := range color_data[0][x] {
			if data[x][y] == 1 {
				color_data[0][x][y] = fr
			} else {
				color_data[0][x][y] = br
			}
		}
	}
	for x := range color_data[1] {
		for y := range color_data[0][x] {
			if data[x][y] == 1 {
				color_data[0][x][y] = fg
			} else {
				color_data[0][x][y] = bg
			}
		}
	}
	for x := range color_data[2] {
		for y := range color_data[0][x] {
			if data[x][y] == 1 {
				color_data[0][x][y] = fb
			} else {
				color_data[0][x][y] = bb
			}
		}
	}

	return color_data

}
