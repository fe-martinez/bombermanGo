package view

import (
	"bombman/model"
	"fmt"
	"sort"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// var d_wall = rl.LoadImage("view/resources/d_wall.png")
// var i_wall = rl.LoadImage("view/resources/i_wall.png")

// var d_wall_texture = rl.LoadTextureFromImage(d_wall)
// var i_wall_texture = rl.LoadTextureFromImage(i_wall)

const TILE_SIZE = 65
const WIDTH = TILE_SIZE * 16
const HEIGHT = TILE_SIZE*10 + OFFSET
const OFFSET = 30

func InitWindow() {
	rl.InitWindow(WIDTH, HEIGHT, "Bomberman Go!")
	rl.SetTargetFPS(30)
}

func WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

func getColorFromString(colorName string) rl.Color {
	switch colorName {
	case "Orange":
		return rl.Orange
	case "Green":
		return rl.Green
	case "Blue":
		return rl.Blue
	case "Violet":
		return rl.Violet
	default:
		return rl.Red
	}
}

func drawPlayers(game model.Game) {
	for _, player := range game.Players {
		if player.Lives == 0 {
			continue
		}
		colorName := game.GetPlayerColor(player.ID)
		color := getColorFromString(colorName)
		rl.DrawRectangle(int32(player.Position.X*TILE_SIZE), int32(player.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, color)
	}
}

func drawBombs(game model.Game) {
	for _, bomb := range game.GameMap.Bombs {
		rl.DrawRectangle(int32(bomb.Position.X*TILE_SIZE), int32(bomb.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Maroon)
	}
}

// Después va a tener que dibujar los distintos powerups según el tipo
func drawPowerUps(game model.Game) {
	for _, powerUp := range game.GameMap.PowerUps {
		rl.DrawRectangle(int32(powerUp.Position.X*TILE_SIZE), int32(powerUp.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Magenta)
	}
}

func drawWalls(game model.Game) {
	for _, wall := range game.GameMap.Walls {
		if wall.Indestructible {
			rl.DrawRectangle(int32(wall.Position.X*TILE_SIZE), int32(wall.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.DarkGray)
		} else {
			rl.DrawRectangle(int32(wall.Position.X*TILE_SIZE), int32(wall.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.LightGray)
		}
	}
}

func drawExplosions(game model.Game) {
	for _, explosion := range game.GameMap.Explosions {
		rl.DrawRectangle(int32(explosion.Position.X*TILE_SIZE), int32(explosion.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Red)
		for _, affectedTile := range explosion.AffectedTiles {
			rl.DrawRectangle(int32(affectedTile.X*TILE_SIZE), int32(affectedTile.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Red)
		}
	}
}

func DrawPlayersLives(game model.Game) {
	// Crear una lista de jugadores a partir del mapa
	var players []*model.Player
	for _, player := range game.Players {
		players = append(players, player)
	}

	// Ordenar los jugadores por nombre
	sort.Slice(players, func(i, j int) bool {
		return players[i].Username < players[j].Username
	})

	var offset int32 = 150
	for _, player := range players {
		playerColor := game.GetPlayerColor(player.ID)
		color := getColorFromString(playerColor)
		lives := strconv.Itoa(int(player.Lives))
		rl.DrawText(fmt.Sprintf("%s: %s <3", player.Username, lives), offset, HEIGHT-OFFSET+5, 20, color)
		offset += 225
	}
}

func DrawGameID(gameID string) {
	rl.DrawRectangle(0, HEIGHT-OFFSET, WIDTH, OFFSET, rl.Black)
	rl.DrawText("Game ID: "+gameID, 3, HEIGHT-OFFSET+5, 20, rl.Red)
}

func DrawGame(game model.Game) {
	if len(game.Players) == 0 {
		return
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	drawWalls(game)

	drawPlayers(game)

	drawBombs(game)

	drawExplosions(game)

	drawPowerUps(game)

	DrawGameID(game.GameId)

	DrawPlayersLives(game)

	rl.EndDrawing()
}
