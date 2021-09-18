package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"reflect"
)
// Game implements ebiten.Game interface
type Game struct{}

// PlayingField Store
type PlayingField [100][100]struct{
	xCoord float64
	yCoord float64
}

var playingField = &PlayingField{
0:{1:{4,0}},
1:{2:{8,4}},
3:{0:{0,8}, 1:{4,8}, 2:{8,8}},
}


// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update () error {
	// Write your game's logical update.
	UpdatePlayingField(playingField)

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for _, line := range playingField {
		for _, cell := range line {
			ebitenutil.DrawRect(screen, cell.xCoord, cell.yCoord, 4, 4,  color.Black)
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
	lineIndex := 0
	columnIndex := 0

	for _, line := range pf {
		for _, column := range line{
			fmt.Printf("%s\n", reflect.TypeOf(column))
			isCellAlive := CheckIsCellAlive(column)


			columnIndex += 4
		}
		lineIndex += 4
	}
}

// TODO find out what type cell is => column is of type struct { xCoord float64; yCoord float64 }
func CheckIsCellAlive (cell) bool {
	if (cell.xCoord > 0 && cell.yCoord > 0) {
		return true
	}
	return false
}