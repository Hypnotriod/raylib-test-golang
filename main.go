package main

import (
	_ "embed"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_TITLE string = "test"
const SCREEN_WIDTH int32 = 800
const SCREEN_HEIGHT int32 = 480
const FPS_MAX int32 = 60
const SCREEN_WIDTH_HALF int32 = SCREEN_WIDTH / 2
const SCREEN_HEIGHT_HALF int32 = SCREEN_HEIGHT / 2
const SCREEN_WIDTH_QUATER int32 = SCREEN_WIDTH / 4
const SCREEN_HEIGHT_QUATER int32 = SCREEN_HEIGHT / 4

//go:embed resources/kitty.png
var kittyBytes []byte
var txrKitty rl.Texture2D

//go:embed resources/superhero.png
var superheroBytes []byte
var txrSuperhero rl.Texture2D

func orbitalposition(texture rl.Texture2D, angle float32, distanceDelta float32, mousePosition rl.Vector2) rl.Vector2 {
	p := rl.Vector2{}
	p.X = float32(math.Cos(float64(angle)))*float32(texture.Width)/2*distanceDelta + mousePosition.X - float32(texture.Width)/2
	p.Y = float32(math.Sin(float64(angle)))*float32(texture.Width)/2*distanceDelta/2 + mousePosition.Y - float32(texture.Height)/2
	return p
}

func initialize() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(FPS_MAX)
	rl.SetExitKey(0)
	rl.HideCursor()

	txrSuperhero = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", kittyBytes, int32(len(kittyBytes))))
	txrKitty = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", superheroBytes, int32(len(superheroBytes))))
}

func loop() {
	var mousePosition rl.Vector2
	var superheroPosition rl.Vector2
	var kittyPosition rl.Vector2
	var angle float32 = 0
	var distanceDelta float32 = 4
	var scaleDelta float32 = 1

	for !rl.WindowShouldClose() {
		angle += (math.Pi * 0.5) * rl.GetFrameTime()
		if angle > math.Pi*2 {
			angle -= math.Pi * 2
		}

		distanceDelta += rl.GetMouseWheelMove() * 0.1

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		mousePosition = rl.GetMousePosition()

		superheroPosition = orbitalposition(txrSuperhero, angle, distanceDelta, mousePosition)
		kittyPosition = orbitalposition(txrKitty, angle+math.Pi, distanceDelta, mousePosition)
		scaleDelta = (superheroPosition.Y - kittyPosition.Y) * 0.001

		if superheroPosition.Y < kittyPosition.Y {
			rl.DrawTextureEx(txrSuperhero, superheroPosition, 0, 1+scaleDelta, rl.White)
			rl.DrawCircle(int32(mousePosition.X), int32(mousePosition.Y), float32(SCREEN_HEIGHT_QUATER), rl.Red)
			rl.DrawText("Ninja", int32(mousePosition.X)-30, int32(mousePosition.Y)-15, 30, rl.White)
			rl.DrawTextureEx(txrKitty, kittyPosition, 0, 1-scaleDelta, rl.White)
		} else {
			rl.DrawTextureEx(txrKitty, kittyPosition, 0, 1-scaleDelta, rl.White)
			rl.DrawCircle(int32(mousePosition.X), int32(mousePosition.Y), float32(SCREEN_HEIGHT_QUATER), rl.Red)
			rl.DrawText("Ninja", int32(mousePosition.X)-30, int32(mousePosition.Y)-15, 30, rl.White)
			rl.DrawTextureEx(txrSuperhero, superheroPosition, 0, 1+scaleDelta, rl.White)
		}

		rl.EndDrawing()
	}
}

func main() {
	initialize()
	loop()
	rl.CloseWindow()
}
