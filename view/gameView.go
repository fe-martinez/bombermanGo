package view

import (
	"bombman/model"
	"fmt"
	"sort"
	"strconv"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILE_SIZE = 65
const WIDTH = TILE_SIZE * 16
const HEIGHT = TILE_SIZE*10 + OFFSET + OFFSET
const OFFSET = 30

var directionMap = map[string]int{
	"down":  0,
	"left":  1,
	"right": 2,
	"up":    3,
}

func InitWindow() {
	rl.InitWindow(WIDTH, HEIGHT, "Bomberman Go!")
	rl.SetTargetFPS(30)
}

func WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

var (
	once                       sync.Once
	onceBomb				   sync.Once
	player1Model               rl.Texture2D
	player2Model               rl.Texture2D
	player3Model               rl.Texture2D
	player4Model               rl.Texture2D
	destructibleWallModel      rl.Texture2D
	indestructibleWallModel    rl.Texture2D
	powerUpAlcanceModel        rl.Texture2D
	powerUpMasBombasModel      rl.Texture2D
	powerUpInvencibilidadModel rl.Texture2D
	bombModel				   rl.Texture2D
)

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

func loadPlayerModel() {
	once.Do(func() {
		player1Model = rl.LoadTexture("./view/resources/player1.png")
		player2Model = rl.LoadTexture("./view/resources/player2.png")
		player3Model = rl.LoadTexture("./view/resources/player3.png")
		player4Model = rl.LoadTexture("./view/resources/player4.png")
		destructibleWallModel = rl.LoadTexture("./view/resources/d_wall.png")
		indestructibleWallModel = rl.LoadTexture("./view/resources/i_wall.png")
		powerUpAlcanceModel = rl.LoadTexture("./view/resources/powerup_alcance.png")
		powerUpMasBombasModel = rl.LoadTexture("./view/resources/powerup_masbombas.png")
		powerUpInvencibilidadModel = rl.LoadTexture("./view/resources/powerup_invencibility.png")
	})
}

func getSourceRect(direction string, game model.Game) rl.Rectangle {
	dirInt := directionMap[direction]

	sourceRect := rl.NewRectangle(
		float32(game.CurrentFrame*65),
		float32(dirInt*68),
		float32(65),
		float32(68), // Height of the frame
	)

	return sourceRect
}

func getPlayerModel(playerColor string) rl.Texture2D {
	switch playerColor {
	case "Orange":
		return player1Model
	case "Green":
		return player2Model
	case "Violet":
		return player3Model
	case "Blue":
		return player4Model
	default:
		return player1Model
	}
}

func drawPlayers(game model.Game) {
	loadPlayerModel()
	for _, player := range game.Players {
		if player.Lives == 0 {
			continue
		}

		sourceRect := getSourceRect(player.Direction, game)
		playerModel := getPlayerModel(game.GetPlayerColor(player.ID))
		position := rl.NewVector2(player.Position.X*TILE_SIZE, player.Position.Y*TILE_SIZE)
		rl.DrawTextureRec(playerModel, sourceRect, position, rl.White)
	}
}

func loadBombModel() {
	onceBomb.Do(func() {
		bombModel = rl.LoadTexture("./view/resources/res.png")
	})
}

func drawBombs(game model.Game) {
	loadBombModel()
	for _, bomb := range game.GameMap.Bombs {
		//rl.DrawRectangle(int32(bomb.Position.X*TILE_SIZE), int32(bomb.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Maroon)
		sourceRect := rl.NewRectangle(float32(0), float32(0),float32(54), float32(65),)
		position := rl.NewVector2(bomb.Position.X*TILE_SIZE, bomb.Position.Y*TILE_SIZE)
		rl.DrawTextureRec(bombModel, sourceRect, position, rl.White)
	}
}

func getPowerUpModel(powerUp model.PowerUpType) *rl.Texture2D {
	switch powerUp {
	case model.AlcanceMejorado:
		return &powerUpAlcanceModel
	case model.MasBombasEnSimultaneo:
		return &powerUpMasBombasModel
	case model.Invencibilidad:
		return &powerUpInvencibilidadModel
	}
	return nil
}

func drawPowerUps(game model.Game) {
	for _, powerUp := range game.GameMap.PowerUps {
		var powerUpModel = getPowerUpModel(powerUp.Name)
		rl.DrawTexture(*powerUpModel, int32(powerUp.Position.X*TILE_SIZE), int32(powerUp.Position.Y*TILE_SIZE), rl.White)
	}
}

func drawWalls(game model.Game) {
	for _, wall := range game.GameMap.Walls {
		x := int32(wall.Position.X * TILE_SIZE)
		y := int32(wall.Position.Y * TILE_SIZE)

		if wall.Indestructible {
			rl.DrawTexture(indestructibleWallModel, x, y, rl.White)
		} else {
			rl.DrawTexture(destructibleWallModel, x, y, rl.White)
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

	var offset int32 = 5
	for _, player := range players {
		playerColor := game.GetPlayerColor(player.ID)
		color := getColorFromString(playerColor)
		lives := strconv.Itoa(int(player.Lives))
		rl.DrawText(fmt.Sprintf("%s: %s <3", player.Username, lives), offset, HEIGHT-OFFSET*2+5, 20, color)
		offset += 230
	}
}

func DrawGameID(gameID string) {
	// Dibujar Game ID separado de los jugadores
	rl.DrawRectangle(0, HEIGHT-OFFSET-OFFSET, WIDTH, OFFSET*2, rl.Black)
	rl.DrawText("Game ID: "+gameID, 5, HEIGHT-OFFSET+5, 20, rl.Red)
}

func DrawGameRound(gameRound string) {
	rl.DrawText("Round: "+gameRound, 150, HEIGHT-OFFSET+5, 20, rl.SkyBlue)
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

	DrawGameRound(strconv.Itoa(int(game.Round)))

	DrawPlayersLives(game)

	rl.EndDrawing()
}
