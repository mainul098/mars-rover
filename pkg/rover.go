package pkg

import (
	"errors"
	"fmt"
)

// ErrInvalidDirection to define a custome error on invalid direction
var ErrInvalidDirection = errors.New("Direction is not valid")

// Direction over the grid
type Direction string

// Directions for the allowed movement on the Grid
const (
	EAST  Direction = "EAST"
	NORTH Direction = "NORTH"
	WEST  Direction = "WEST"
	SOUTH Direction = "SOUTH"
)

// Rotation define the left or right direction for a axis on Grid
type Rotation struct {
	left  Direction
	right Direction
}

// Coordinate points on the Grid
type Coordinate struct {
	x int
	y int
}

// Rover to explore the Mars
type Rover struct {
	coordinate Coordinate
	direction  Direction
	rotations  map[Direction]Rotation
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
		coordinate: Coordinate{x, y},
		direction:  direction,
		rotations:  rotations,
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r *Rover) ExecuteCommand(commands string) (string, error) {
	for _, c := range commands {
		if err := r.performAction(string(c)); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("(%d, %d) %s", r.coordinate.x, r.coordinate.y, r.direction), nil
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
	c, err := nextCoordinate(r.coordinate, r.direction, true)
	if err != nil {
		return err
	}

	r.coordinate = c
	return nil
}

func (r *Rover) moveBackward() error {
	c, err := nextCoordinate(r.coordinate, r.direction, false)
	if err != nil {
		return err
	}

	r.coordinate = c
	return nil
}

func (r *Rover) rotateLeft() error {
	if rotation, ok := r.rotations[r.direction]; ok {
		r.direction = rotation.left
	} else {
		return ErrInvalidDirection
	}
	return nil
}

func (r *Rover) rotateRight() error {
	if rotation, ok := r.rotations[r.direction]; ok {
		r.direction = rotation.right
	} else {
		return ErrInvalidDirection
	}
	return nil
}

func nextCoordinate(coordinate Coordinate, direction Direction, isMoveForward bool) (Coordinate, error) {
	x := coordinate.x
	y := coordinate.y
	step := 1

	if !isMoveForward {
		step = -1
	}

	if direction == EAST {
		return Coordinate{x + step, y}, nil
	}

	if direction == NORTH {
		return Coordinate{x, y + step}, nil
	}

	if direction == WEST {
		return Coordinate{x - step, y}, nil
	}

	if direction == SOUTH {
		return Coordinate{x, y - step}, nil
	}

	return coordinate, ErrInvalidDirection
}
