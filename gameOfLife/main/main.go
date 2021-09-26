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
type PlayingField [100][100]struct{
	xCoord float64
	yCoord float64
}

var playingField = PlayingField{
1:{2:{8,4}},
2:{3:{12,8}},
3:{1:{4,12}, 2:{8,12}, 3:{12,12}},
}
var nextGen = PlayingField{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update () error {
	// Write your game's logical update.
	nextGen = UpdatePlayingField(playingField)
	playingField = nextGen
	nextGen = PlayingField{}
	time.Sleep(200 * time.Millisecond)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for _, line := range playingField {
		for _, cell := range line {
			if cell.xCoord > 0 || cell.yCoord > 0 {
				ebitenutil.DrawRect(screen, cell.xCoord, cell.yCoord, 4, 4,  color.Black)
			}
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
func UpdatePlayingField (pf PlayingField) PlayingField {
	nextGenPlayingField := PlayingField{}

	var yCoord float64 = 0
	for i, line := range pf {
		var xCoord float64 = 0
		for j, cell := range line {
			// ignore edges for neighbour counting
			if i == 0 || j == 0 || i == len(pf) - 1 || j == len(line) - 1 {
				nextGenPlayingField[i][j].xCoord = pf[i][j].xCoord
				nextGenPlayingField[i][j].yCoord = pf[i][j].yCoord
			} else {
				isCellAlive := CheckIsCellAlive(cell.xCoord, cell.yCoord)
				numOfNeighbors := CheckNumberOfNeighbors(i, j, pf, isCellAlive)
				if numOfNeighbors == 3 {
					nextGenPlayingField[i][j].xCoord = xCoord
					nextGenPlayingField[i][j].yCoord = yCoord
				} else if numOfNeighbors < 2 || numOfNeighbors > 3 {
					nextGenPlayingField[i][j].xCoord = 0
					nextGenPlayingField[i][j].yCoord = 0
				} else {
					nextGenPlayingField[i][j].xCoord = pf[i][j].xCoord
					nextGenPlayingField[i][j].yCoord = pf[i][j].yCoord
				}
			}
			xCoord += 4
		}
		yCoord += 4
	}
	return nextGenPlayingField
}

func CheckIsCellAlive (x float64, y float64 ) bool {
	if x > 0 || y > 0 {
		return true
	}
	return false
}

func CheckNumberOfNeighbors (lineIndex int, columnIndex int, field PlayingField, alive bool) int {
	neighborsCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if field[lineIndex + i][columnIndex + j].xCoord > 0 ||
				field[lineIndex + i][columnIndex + j].yCoord > 0 {
				neighborsCount += 1
			}
		}
	}

	// subtract the cell from which the checks are going out
	if alive {
		neighborsCount -= 1
	}

	return neighborsCount
}
