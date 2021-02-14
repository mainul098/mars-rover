package pkg

import (
	"errors"
	"fmt"
)

// Rover to explore the Mars
type Rover struct {
	coordinate Coordinate
	direction  Direction
	grid       Grid
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction Direction) *Rover {
	return &Rover{
		coordinate: Coordinate{x, y},
		direction:  direction,
		grid:       NewGrid(),
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r Rover) ExecuteCommand(commands string) (string, error) {
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
		c, err := r.grid.moveForward(r.coordinate, r.direction)
		if err != nil {
			return err
		}
		r.coordinate = c
	case "B":
		c, err := r.grid.moveBackward(r.coordinate, r.direction)
		if err != nil {
			return err
		}
		r.coordinate = c
	case "L":
		d, err := r.grid.rotateLeft(r.direction)
		if err != nil {
			return err
		}
		r.direction = d
	case "R":
		d, err := r.grid.rotateRight(r.direction)
		if err != nil {
			return err
		}
		r.direction = d
	default:
		return errors.New("Invalid command")
	}
	return nil
}
