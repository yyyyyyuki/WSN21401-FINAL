package main

import (
	"math/rand"
	"time"
)

func GenerateImgMatrix(height, width, steps int) [][]uint8 {
	data := make([][]uint8, height)
	for i := range data {
		data[i] = make([]uint8, width)
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
