package pkg

import (
	"errors"
	"fmt"
)

// Direction over the grid
type Direction string

// Directions for the allowed movement
const (
	EAST  Direction = "EAST"
	NORTH Direction = "NORTH"
	WEST  Direction = "WEST"
	SOUTH Direction = "SOUTH"
)

// Rotation define the left or right direction for a axis
type Rotation struct {
	left  Direction
	right Direction
}

// Rover to explore the Mars
type Rover struct {
	x         int
	y         int
	direction Direction
	rotations map[Direction]Rotation
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction Direction) Rover {
	rotations := map[Direction]Rotation{
		EAST:  {NORTH, SOUTH},
		NORTH: {WEST, EAST},
		WEST:  {SOUTH, NORTH},
		SOUTH: {EAST, WEST},
	}

	return Rover{
		x:         x,
		y:         y,
		direction: direction,
		rotations: rotations,
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
		case "L":
			r.rotateLeft()
		case "R":
			r.rotateRight()
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

func (r *Rover) rotateLeft() error {
	r.direction = r.rotations[r.direction].left
	return nil
}

func (r *Rover) rotateRight() error {
	r.direction = r.rotations[r.direction].right
	return nil
}
