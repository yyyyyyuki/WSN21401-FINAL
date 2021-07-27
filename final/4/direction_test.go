package main

import (
	"fmt"
	"testing"
)

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}
}

func TestMovePoint(t *testing.T) {
	var err error
	p0 := Point{X: 5, Y: 5}
	p1 := Point{}

	p1, err = MovePoint(p0, DirectionUp, 0, 0, 10, 10)
	t.Log(p1)
	checkError(t, err)

	if p1.Y != p0.Y-1 {
		t.Fail()
	}
}

func TestSetOfDirections(t *testing.T) {
	var s SetOfDirections

	s.Add(DirectionUp)

	fmt.Println(s.list)

	if s.Len() != 1 {
		t.Error("TestSetOfDirections is failed")
	}

	if s.Contains(DirectionUp) {
		s.Remove(DirectionUp)
	} else {
		s.Add(DirectionUp)
	}

	fmt.Println(s.list)

	if s.Len() != 0 {
		t.Error("TestSetOfDirections is failed")
	}
}
