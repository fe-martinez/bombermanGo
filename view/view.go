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
const HEIGHT = TILE_SIZE*10 + OFFSET
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
	once                    sync.Once
	player1Model            rl.Texture2D
	player2Model            rl.Texture2D
	player3Model            rl.Texture2D
	player4Model            rl.Texture2D
	destructibleWallModel   rl.Texture2D
	indestructibleWallModel rl.Texture2D
	counter                 = 0
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

func drawPlayers(game model.Game) {
	loadPlayerModel()
	for _, player := range game.Players {
		if player.Lives == 0 {
			continue
		}

		sourceRect := getSourceRect(player.Direction, game)
		/*sourceRect := rl.NewRectangle(
					//float32(65*counter),
					//float32(68),
					float32(game.CurrentFrame * 65),
		        	float32(game.CurrentFrame * 68),
		        	float32(65),
		        	float32(68),                              // Height of the frame
				)*/
		player_color := game.GetPlayerColor(player.ID)
		var playerModel rl.Texture2D
		switch player_color {
		case "Orange":
			playerModel = player1Model
		case "Green":
			playerModel = player2Model
		case "Violet":
			playerModel = player3Model
		case "Blue":
			playerModel = player4Model
		}
		position := rl.NewVector2(player.Position.X*TILE_SIZE, player.Position.Y*TILE_SIZE)
		rl.DrawTextureRec(playerModel, sourceRect, position, rl.White)
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

const LOBBY_TITLE_POS_X = WIDTH/2 - 70
const LOBBY_TITLE_POS_Y = 200

const CREATE_GAME_POS_X = WIDTH/2 - 170
const CREATE_GAME_POS_Y = 300

const JOIN_GAME_POS_X = WIDTH/2 - 170
const JOIN_GAME_POS_Y = 400

const BUTTON_WIDTH = 350
const BUTTON_HEIGHT = 50

func DrawMainMenuScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	rl.DrawText("Lobby screen", LOBBY_TITLE_POS_X, LOBBY_TITLE_POS_Y, 20, rl.Maroon)
	rl.DrawRectangle(CREATE_GAME_POS_X, CREATE_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Create game", CREATE_GAME_POS_X+10, CREATE_GAME_POS_Y+15, 20, rl.White)
	rl.DrawRectangle(JOIN_GAME_POS_X, JOIN_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Join game", JOIN_GAME_POS_X+10, JOIN_GAME_POS_Y+15, 20, rl.White)
	rl.EndDrawing()
}

const (
	INPUT_BOX_POS_X  = WIDTH/2 - 100
	INPUT_BOX_POS_Y  = HEIGHT/2 - 25
	INPUT_BOX_WIDTH  = 350
	INPUT_BOX_HEIGHT = 50
)

// Raylib no tiene cajas de texto, este es un intento de simular una
func DrawLobbySelectionScreen(lobbyID string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	rl.DrawText("Enter Game ID", INPUT_BOX_POS_X, INPUT_BOX_POS_Y-40, 20, rl.Maroon)
	rl.DrawRectangleLines(INPUT_BOX_POS_X-95, INPUT_BOX_POS_Y, INPUT_BOX_WIDTH, INPUT_BOX_HEIGHT, rl.DarkPurple)

	rl.DrawText(lobbyID, INPUT_BOX_POS_X-90, INPUT_BOX_POS_Y+15, 20, rl.Maroon)

	rl.EndDrawing()
}

const (
	START_TITLE_POS_X  = 400
	START_TITLE_POS_Y  = 200
	START_GAME_POS_X   = 370
	START_GAME_POS_Y   = 350
	PLAYER_START_POS_X = 100
	PLAYER_START_POS_Y = 150
	PLAYER_GAP_Y       = 30
	GAME_ID_POS_X      = 50  // Posición X para el texto del ID del juego
	GAME_ID_POS_Y      = 50  // Posición Y para el texto del ID del juego
	TEXTBOX_WIDTH      = 200 // Ancho del recuadro para el ID del juego
	TEXTBOX_HEIGHT     = 30  // Alto del recuadro para el ID del juego
)

func DrawWaitingMenu(players []string, lobbyId string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	rl.DrawText("Are you ready for the game?", START_TITLE_POS_X, START_TITLE_POS_Y, 20, rl.Maroon)
	// Draw Game ID
	rl.DrawText("Game ID:", GAME_ID_POS_X, GAME_ID_POS_Y, 20, rl.Black)
	rl.DrawRectangle(GAME_ID_POS_X, GAME_ID_POS_Y+25, TEXTBOX_WIDTH, TEXTBOX_HEIGHT, rl.DarkGray)
	rl.DrawText(lobbyId, GAME_ID_POS_X+5, GAME_ID_POS_Y+30, 20, rl.White)
	// Draw Connected players
	rl.DrawText("Connected players:", PLAYER_START_POS_X, PLAYER_START_POS_Y-30, 20, rl.Black)
	// Draw players
	for i, player := range players {
		yPos := PLAYER_START_POS_Y + int32(i)*PLAYER_GAP_Y
		rl.DrawText(player, PLAYER_START_POS_X, yPos, 20, rl.Black)
	}
	// Draw Start Game button
	rl.DrawRectangle(START_GAME_POS_X, START_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Start Game", START_GAME_POS_X+10, START_GAME_POS_Y+15, 20, rl.White)
	rl.EndDrawing()
}

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

	// Draw a white rectangle with a black border to put the players on
	rl.DrawRectangle(WIDTH/2-250, HEIGHT/2-60, 500, 30*4, rl.White)
	rl.DrawRectangleLines(WIDTH/2-250, HEIGHT/2-60, 500, 30*4, rl.Black)

	for _, player := range players {
		color := getColorFromString(game.GetPlayerColor(player.ID))
		playerName := fmt.Sprintf("%s - Points: %d", player.Username, game.PlayerScores[player.ID])
		rl.DrawText(playerName, WIDTH/2-rl.MeasureText(playerName, 30)/2, HEIGHT/2-15+int32(index)*30, 30, color)
		index++
	}
	rl.EndDrawing()
}
