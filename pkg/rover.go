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
func (r *Rover) ExecuteCommand(commands string) (string, error) {
	for _, c := range commands {
		if err := r.performAction(string(c)); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("(%d, %d) %s", r.x, r.y, r.direction), nil
}

func (r *Rover) performAction(command string) error {
	switch string(command) {
	case "F":
		return r.moveForward()
	case "B":
		return r.moveBackward()
	case "L":
		return r.rotateLeft()
	case "R":
		return r.rotateRight()
	default:
		return errors.New("Invalid command")
	}
}

func (r *Rover) moveForward() error {
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

func (r *Rover) moveBackward() error {
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
