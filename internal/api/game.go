package api

import (
	"errors"
	"math/rand"
	"time"
)

const GridSize = 5
const NumTypes = 4

type Game struct {
	Grid [][]int `json:"grid"`
}

type MoveRequest struct {
	Row       int    `json:"row"`
	Col       int    `json:"col"`
	Direction string `json:"direction"` // "up", "down", "left", "right"
}

var currentGame *Game

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	grid := make([][]int, GridSize)
	for i := range grid {
		grid[i] = make([]int, GridSize)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(NumTypes)
		}
	}
	return &Game{Grid: grid}
}

func storeGame(game *Game) {
	currentGame = game
}

func getGame() *Game {
	if currentGame == nil {
		currentGame = NewGame()
	}
	return currentGame
}

func (g *Game) MakeMove(row, col int, direction string) error {
	if row < 0 || col < 0 || row >= GridSize || col >= GridSize {
		return errors.New("invalid coordinates")
	}

	var newRow, newCol int
	switch direction {
	case "up":
		newRow, newCol = row-1, col
	case "down":
		newRow, newCol = row+1, col
	case "left":
		newRow, newCol = row, col-1
	case "right":
		newRow, newCol = row, col+1
	default:
		return errors.New("invalid direction")
	}

	if newRow < 0 || newCol < 0 || newRow >= GridSize || newCol >= GridSize {
		return errors.New("move out of bounds")
	}

	// swap
	g.Grid[row][col], g.Grid[newRow][newCol] = g.Grid[newRow][newCol], g.Grid[row][col]
	return nil
}
