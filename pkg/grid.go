package pkg

import "errors"

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

// Grid for the rover coordinate plane
type Grid struct {
	rotations map[Direction]Rotation
	axes      map[Direction]Coordinate
}

// NewGrid to move the Rover
func NewGrid() Grid {
	rotations := map[Direction]Rotation{
		EAST:  {NORTH, SOUTH},
		NORTH: {WEST, EAST},
		WEST:  {SOUTH, NORTH},
		SOUTH: {EAST, WEST},
	}

	axes := map[Direction]Coordinate{
		EAST:  {1, 0},
		NORTH: {0, 1},
		WEST:  {-1, 0},
		SOUTH: {0, -1},
	}

	return Grid{
		rotations: rotations,
		axes:      axes,
	}
}

func (g Grid) moveForward(coordinate Coordinate, direction Direction) (Coordinate, error) {
	if axis, ok := g.axes[direction]; ok {
		return Coordinate{coordinate.x + axis.x, coordinate.y + axis.y}, nil
	}
	return coordinate, ErrInvalidDirection
}

func (g Grid) moveBackward(coordinate Coordinate, direction Direction) (Coordinate, error) {
	if axis, ok := g.axes[direction]; ok {
		return Coordinate{coordinate.x - axis.x, coordinate.y - axis.y}, nil
	}
	return coordinate, ErrInvalidDirection
}

func (g Grid) rotateLeft(direction Direction) (Direction, error) {
	if rotation, ok := g.rotations[direction]; ok {
		return rotation.left, nil
	}

	return direction, ErrInvalidDirection
}

func (g Grid) rotateRight(direction Direction) (Direction, error) {
	if rotation, ok := g.rotations[direction]; ok {
		return rotation.right, nil
	}

	return direction, ErrInvalidDirection
}
