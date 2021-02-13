package pkg

import (
	"errors"
	"fmt"
)

// Direction denote the direction
type Direction string

// Directions
const (
	EAST  Direction = "EAST"
	NORTH Direction = "NORTH"
	WEST  Direction = "WEST"
	SOUTH Direction = "SOUTH"
)

// Rover to explore the Mars
type Rover struct {
	x         int
	y         int
	direction Direction
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction Direction) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r *Rover) ExecuteCommand(command string) (string, error) {
	for _, c := range command {
		switch string(c) {
		case "F":
			r.forward()
		case "B":
			r.backward()
		default:
			return "", errors.New("Invalid command")
		}
	}
	return fmt.Sprintf("(%d, %d) %s", r.x, r.y, r.direction), nil
}

func (r *Rover) forward() error {
	if r.direction == EAST {
		r.x++
	} else if r.direction == NORTH {
		r.y++
	} else if r.direction == WEST {
		r.x--
	} else if r.direction == SOUTH {
		r.y--
	}
	return nil
}

func (r *Rover) backward() error {
	if r.direction == EAST {
		r.x--
	} else if r.direction == NORTH {
		r.y--
	} else if r.direction == WEST {
		r.x++
	} else if r.direction == SOUTH {
		r.y++
	}
	return nil
}
