package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"time"
)
// Game implements ebiten.Game interface
type Game struct{}

// PlayingField Store
type PlayingField [10][10]struct{
	xCoord float64
	yCoord float64
}

var playingField = PlayingField{
0:{1:{40,0}},
1:{2:{80,40}},
2:{0:{0,80}, 1:{40,80}, 2:{80,80}},
}

var firstIteration bool = true

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update () error {
	// Write your game's logical update.
	if !firstIteration {
		UpdatePlayingField(&playingField)
	}
	firstIteration = false
	time.Sleep(1 * time.Second)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for _, line := range &playingField {
		for _, cell := range line {
			ebitenutil.DrawRect(screen, cell.xCoord, cell.yCoord, 40, 40,  color.Black)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int)(screenWidth, screenHeight int) {
	// Return the game screen size
	return 404, 404
}

func main() {
	ebiten.SetWindowSize(804, 804)
	ebiten.SetWindowTitle("Game Of Life")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

// The game of life algorithm
func UpdatePlayingField (pf *PlayingField) {
	var yCoord float64 = 0
	for i, line := range pf {
		var xCoord float64 = 0

		for j, column := range line{
			isCellAlive := CheckIsCellAlive(column.xCoord, column.yCoord)
			if !isCellAlive {
				xCoord += 40
				continue
			} else if isCellAlive {
				numOfNeighbors := CheckNumberOfNeighbors(i, j, pf)

				if numOfNeighbors < 2 || numOfNeighbors > 3 {
					pf[i][j].xCoord = 0
					pf[i][j].yCoord = 0
				} else {
					pf[i][j].xCoord = xCoord
					pf[i][j].yCoord = yCoord
				}
			}
			xCoord += 40
		}
		yCoord += 40
	}
}

func CheckIsCellAlive (x float64, y float64 ) bool {
	if x > 0 || y > 0 {
		return true
	}
	return false
}

func CheckNumberOfNeighbors (lineIndex int, columnIndex int, field *PlayingField) int {
	neighborsCount := 0
	if lineIndex > 0 && columnIndex > 0 {
		if field[lineIndex-1][columnIndex-1].xCoord > 0 &&
			field[lineIndex-1][columnIndex-1].yCoord > 0 {
			neighborsCount += 1
		}
	}
	if lineIndex > 0 {
		if field[lineIndex-1][columnIndex].xCoord > 0 &&
			field[lineIndex-1][columnIndex].yCoord > 0 {
			neighborsCount += 1
		}
		if field[lineIndex-1][columnIndex+1].xCoord > 0 &&
			field[lineIndex-1][columnIndex+1].yCoord > 0 {
			neighborsCount += 1
		}
	}
	if columnIndex > 0 {
		if field[lineIndex][columnIndex-1].xCoord > 0 &&
			field[lineIndex][columnIndex-1].yCoord > 0 {
			neighborsCount += 1
		}
	}
	if field[lineIndex][columnIndex + 1].xCoord > 0 &&
		field[lineIndex][columnIndex + 1].yCoord > 0 {
		neighborsCount += 1
	}
	if columnIndex > 0 {
		if field[lineIndex+1][columnIndex-1].xCoord > 0 &&
			field[lineIndex+1][columnIndex-1].yCoord > 0 {
			neighborsCount += 1
		}
	}
	if field[lineIndex + 1][columnIndex].xCoord > 0 &&
		field[lineIndex + 1][columnIndex].yCoord > 0 {
		neighborsCount += 1
	}
	if field[lineIndex + 1][columnIndex + 1].xCoord > 0 &&
		field[lineIndex + 1][columnIndex + 1].yCoord > 0 {
		neighborsCount += 1
	}
	return neighborsCount
}