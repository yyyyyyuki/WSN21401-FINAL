package main

import (
	"errors"
	// "math/rand"
)

type Direction int

const (
	DirectionUp        Direction = 0
	DirectionUpRight   Direction = 1
	DirectionRight     Direction = 2
	DirectionDownRight Direction = 3
	DirectionDown      Direction = 4
	DirectionDownLeft  Direction = 5
	DirectionLeft      Direction = 6
	DirectionUpLeft    Direction = 7
)

func (d Direction) String() string {
	if d == DirectionUp {
		return "DirectionUp"
	}
	if d == DirectionUpRight {
		return "DirectionUpRight"
	}
	if d == DirectionRight {
		return "DirectionRight"
	}
	if d == DirectionDownRight {
		return "DirectionDownRight"
	}
	if d == DirectionDown {
		return "DirectionDown"
	}
	if d == DirectionDownLeft {
		return "DirectionDownLeft"
	}
	if d == DirectionLeft {
		return "DirectionLeft"
	}
	if d == DirectionUpLeft {
		return "DirectionUpLeft"
	}

	return "Unknown Direction"
}

type Point struct {
	X int
	Y int
}

func MovePoint(p Point, dir Direction, min_x, min_y, max_x, max_y int) (Point, error) {
	err := errors.New("Position out of range")

	if dir == DirectionUp {
		p.Y -= 1
	}
	if dir == DirectionDown {
		p.Y += 1
	}
	if dir == DirectionRight {
		p.X += 1
	}
	if dir == DirectionLeft {
		p.X -= 1
	}

	if dir == DirectionUpRight {
		p.X += 1
		p.Y -= 1
	}
	if dir == DirectionDownRight {
		p.X += 1
		p.Y += 1
	}
	if dir == DirectionDownLeft {
		p.X -= 1
		p.Y += 1
	}
	if dir == DirectionUpLeft {
		p.X -= 1
		p.Y -= 1
	}

	if p.X < min_x {
		return p, err
	}

	if p.X > max_x {
		return p, err
	}

	if p.Y < min_y {
		return p, err
	}

	if p.Y > max_y {
		return p, err
	}

	return p, nil
}
