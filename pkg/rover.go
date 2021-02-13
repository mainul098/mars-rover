package pkg

import (
	"errors"
	"fmt"
)

// Rover to explore the Mars
type Rover struct {
	x         int
	y         int
	direction string
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction string) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r *Rover) ExecuteCommand(command string) error {
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
			return errors.New("Invalid command")
		}
	}
	return nil
}

func (r *Rover) forward() error {
	r.x++
	return nil
}

func (r *Rover) backward() error {
	r.x--
	return nil
}

func (r *Rover) rotateLeft() error {
	return nil
}

func (r *Rover) rotateRight() error {
	return nil
}

// GetCurrentPosition will give the current position for the rover
func (r *Rover) GetCurrentPosition() (string, error) {
	return fmt.Sprintf("(%d, %d) %s", r.x, r.y, r.direction), nil
}
