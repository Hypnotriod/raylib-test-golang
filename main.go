package main

import (
	_ "embed"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_TITLE string = "Ninja"
const SCREEN_WIDTH int32 = 800
const SCREEN_HEIGHT int32 = 480
const FPS_MAX int32 = 60
const SCREEN_WIDTH_HALF int32 = SCREEN_WIDTH / 2
const SCREEN_HEIGHT_HALF int32 = SCREEN_HEIGHT / 2
const SCREEN_WIDTH_QUARTER int32 = SCREEN_WIDTH / 4
const SCREEN_HEIGHT_QUARTER int32 = SCREEN_HEIGHT / 4

//go:embed resources/icon32.png
var icon32PNG []byte

//go:embed resources/kitty.png
var kittyPNG []byte
var kittyTexture rl.Texture2D

//go:embed resources/superhero.png
var superheroPNG []byte
var superheroTexture rl.Texture2D

func orbitalPosition(texture rl.Texture2D, angle float32, distanceDelta float32, mousePosition rl.Vector2) rl.Vector2 {
	p := rl.Vector2{}
	p.X = float32(math.Cos(float64(angle)))*float32(texture.Width)/2*distanceDelta + mousePosition.X - float32(texture.Width)/2
	p.Y = float32(math.Sin(float64(angle)))*float32(texture.Width)/2*distanceDelta/2 + mousePosition.Y - float32(texture.Height)/2
	return p
}

func startup() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(FPS_MAX)
	rl.SetExitKey(0)
	rl.HideCursor()
	rl.SetWindowIcon(*rl.LoadImageFromMemory(".png", icon32PNG, int32(len(icon32PNG))))

	superheroTexture = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", kittyPNG, int32(len(kittyPNG))))
	kittyTexture = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", superheroPNG, int32(len(superheroPNG))))
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

		superheroPosition = orbitalPosition(superheroTexture, angle, distanceDelta, mousePosition)
		kittyPosition = orbitalPosition(kittyTexture, angle+math.Pi, distanceDelta, mousePosition)
		scaleDelta = (superheroPosition.Y - kittyPosition.Y) * 0.001

		if superheroPosition.Y < kittyPosition.Y {
			rl.DrawTextureEx(superheroTexture, superheroPosition, 0, 1+scaleDelta, rl.White)
			rl.DrawCircle(int32(mousePosition.X), int32(mousePosition.Y), float32(SCREEN_HEIGHT_QUARTER), rl.Red)
			rl.DrawText("Ninja", int32(mousePosition.X)-30, int32(mousePosition.Y)-15, 30, rl.White)
			rl.DrawTextureEx(kittyTexture, kittyPosition, 0, 1-scaleDelta, rl.White)
		} else {
			rl.DrawTextureEx(kittyTexture, kittyPosition, 0, 1-scaleDelta, rl.White)
			rl.DrawCircle(int32(mousePosition.X), int32(mousePosition.Y), float32(SCREEN_HEIGHT_QUARTER), rl.Red)
			rl.DrawText("Ninja", int32(mousePosition.X)-30, int32(mousePosition.Y)-15, 30, rl.White)
			rl.DrawTextureEx(superheroTexture, superheroPosition, 0, 1+scaleDelta, rl.White)
		}

		rl.EndDrawing()
	}
}

func main() {
	startup()
	loop()
	rl.CloseWindow()
}
