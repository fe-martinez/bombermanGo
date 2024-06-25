package view

import (
	"bombman/model"
	"fmt"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const RETURN_MAIN_MENU_BUTTON_X = (WIDTH - BUTTON_WIDTH) / 2
const RETURN_MAIN_MENU_BUTTON_Y = HEIGHT - 100

func DrawGameOverScreen(game model.Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	rl.DrawText("Game Over", WIDTH/2-75, 50, 30, rl.Maroon)
	index := 0

	var players []model.Player

	for _, player := range game.Players {
		players = append(players, *player)
	}

	sort.Slice(players, func(i, j int) bool {
		return game.PlayerScores[players[i].ID] > game.PlayerScores[players[j].ID]
	})

	rl.DrawRectangle(WIDTH/2-250, HEIGHT/2-60, 500, 30*4, rl.White)
	rl.DrawRectangleLines(WIDTH/2-250, HEIGHT/2-60, 500, 30*4, rl.Black)

	for _, player := range players {
		color := getColorFromString(game.GetPlayerColor(player.ID))
		playerName := fmt.Sprintf("%s - Points: %d", player.Username, game.PlayerScores[player.ID])
		rl.DrawText(playerName, WIDTH/2-rl.MeasureText(playerName, 30)/2, HEIGHT/2-15+int32(index)*30, 30, color)
		index++
	}

	rl.DrawRectangle(RETURN_MAIN_MENU_BUTTON_X, RETURN_MAIN_MENU_BUTTON_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Return to main menu", RETURN_MAIN_MENU_BUTTON_X+10, RETURN_MAIN_MENU_BUTTON_Y+15, 20, rl.White)

	rl.EndDrawing()
}
