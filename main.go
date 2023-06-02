package main

import (
	"fmt"
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var controlsShowing = false

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - custom uniform variable")

	shader := rl.LoadShader("", "glsl330/shader.fs")
	timeLoc := rl.GetShaderLocation(shader, "iTime")

	resLoc := rl.GetShaderLocation(shader, "iResolution")
	resolution := []float32{float32(screenWidth), float32(screenHeight)}
	rl.SetShaderValue(shader, resLoc, resolution, rl.ShaderUniformVec2)

	paLoc := rl.GetShaderLocation(shader, "pa")
	pa := []float32{0.5, 0.5, 0.5}
	rl.SetShaderValue(shader, paLoc, pa, rl.ShaderUniformVec3)
	pbLoc := rl.GetShaderLocation(shader, "pb")
	pb := []float32{0.5, 0.5, 0.5}
	rl.SetShaderValue(shader, pbLoc, pb, rl.ShaderUniformVec3)
	pcLoc := rl.GetShaderLocation(shader, "pc")
	pc := []float32{1.0, 1.0, 1.0}
	rl.SetShaderValue(shader, pcLoc, pc, rl.ShaderUniformVec3)
	pdLoc := rl.GetShaderLocation(shader, "pd")
	pd := []float32{0.263, 0.416, 0.557}
	rl.SetShaderValue(shader, pdLoc, pd, rl.ShaderUniformVec3)

	rl.SetTargetFPS(144)

	for !rl.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		//if rl.IsWindowResized() {
		//	screenWidth = int32(rl.GetScreenWidth())
		//	screenHeight = int32(rl.GetScreenHeight())
		//
		//	resolution[0] = float32(screenWidth)
		//	resolution[1] = float32(screenHeight)
		//	rl.SetShaderValue(shader, resLoc, resolution, rl.ShaderUniformVec2)
		//}

		if rl.IsKeyPressed(rl.KeyF1) {
			controlsShowing = !controlsShowing
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		time := []float32{float32(rl.GetTime())}
		rl.SetShaderValue(shader, timeLoc, time, rl.ShaderUniformFloat)

		rl.SetShaderValue(shader, paLoc, pa, rl.ShaderUniformVec3)
		rl.SetShaderValue(shader, pbLoc, pb, rl.ShaderUniformVec3)
		rl.SetShaderValue(shader, pcLoc, pc, rl.ShaderUniformVec3)
		rl.SetShaderValue(shader, pdLoc, pd, rl.ShaderUniformVec3)

		rl.DrawText("TEXT DRAWN IN RENDER TEXTURE", 200, 10, 30, rl.Red)

		rl.BeginShaderMode(shader)
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Black)
		rl.EndShaderMode()

		rl.DrawFPS(screenWidth-80, 10)

		if controlsShowing {
			pa[0] = gui.Slider(rl.NewRectangle(25, 10, 200, 20), "pa1", fmt.Sprintf("%2.5f", pa[0]), pa[0], 0.0, 1.0)
			pa[1] = gui.Slider(rl.NewRectangle(25, 40, 200, 20), "pa2", fmt.Sprintf("%2.5f", pa[1]), pa[1], 0.0, 1.0)
			pa[2] = gui.Slider(rl.NewRectangle(25, 70, 200, 20), "pa3", fmt.Sprintf("%2.5f", pa[2]), pa[2], 0.0, 1.0)
			pb[0] = gui.Slider(rl.NewRectangle(25, 100, 200, 20), "pb1", fmt.Sprintf("%2.5f", pb[0]), pb[0], 0.0, 1.0)
			pb[1] = gui.Slider(rl.NewRectangle(25, 130, 200, 20), "pb2", fmt.Sprintf("%2.5f", pb[1]), pb[1], 0.0, 1.0)
			pb[2] = gui.Slider(rl.NewRectangle(25, 160, 200, 20), "pb3", fmt.Sprintf("%2.5f", pb[2]), pb[2], 0.0, 1.0)
			pc[0] = gui.Slider(rl.NewRectangle(25, 190, 200, 20), "pc1", fmt.Sprintf("%2.5f", pc[0]), pc[0], 0.0, 1.0)
			pc[1] = gui.Slider(rl.NewRectangle(25, 220, 200, 20), "pc2", fmt.Sprintf("%2.5f", pc[1]), pc[1], 0.0, 1.0)
			pc[2] = gui.Slider(rl.NewRectangle(25, 250, 200, 20), "pc3", fmt.Sprintf("%2.5f", pc[2]), pc[2], 0.0, 1.0)
			pd[0] = gui.Slider(rl.NewRectangle(25, 280, 200, 20), "pd1", fmt.Sprintf("%2.5f", pd[0]), pd[0], 0.0, 1.0)
			pd[1] = gui.Slider(rl.NewRectangle(25, 310, 200, 20), "pd2", fmt.Sprintf("%2.5f", pd[1]), pd[1], 0.0, 1.0)
			pd[2] = gui.Slider(rl.NewRectangle(25, 340, 200, 20), "pd3", fmt.Sprintf("%2.5f", pd[2]), pd[2], 0.0, 1.0)
		}

		rl.EndDrawing()
	}

	rl.UnloadShader(shader) // Unload shader

	rl.CloseWindow()
}
