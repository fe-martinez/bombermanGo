package model

import (
	"log"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_PLAYERS = 4
const MAX_ROUNDS = 5
const ROUND_DURATION = 2 //minutes
const TICKER_REFRESH = 1 //second
const MAX_POWER_UPS = 4
const POWERUP_SPAWN_TIME = 10 //seconds

const SPEED_INCREMENT = 0.1
const BASE_SPEED = 0

var colors = NewQueue()

var stopChan chan struct{}

type Game struct {
	State        string
	GameId       string
	Round        int8
	Players      map[string]*Player
	PlayerColors map[string]string
	GameMap      *GameMap
}

func initializeColors() {
	colors.Enqueue("Orange")
	colors.Enqueue("Green")
	colors.Enqueue("Violet")
	colors.Enqueue("Blue")
}

func NewGame(id string, GameMap *GameMap) *Game {
	initializeColors()
	return &Game{
		State:        "not-started",
		GameId:       id,
		Round:        1,
		Players:      make(map[string]*Player),
		PlayerColors: make(map[string]string),
		GameMap:      GameMap,
	}
}

func (g *Game) collidesWithWalls(position Position) bool {
	pos := rl.NewRectangle(position.X*65+5, position.Y*65+5, 55, 55)

	for _, wall := range g.GameMap.Walls {
		wallRect := rl.NewRectangle(wall.Position.X*65, wall.Position.Y*65, 65, 65)
		if rl.CheckCollisionRecs(pos, wallRect) {
			return true
		}
	}
	return false
}

func (g *Game) isOutOfBounds(position Position) bool {
	return position.X < 0 || position.X >= float32(g.GameMap.ColumnSize) || position.Y < 0 || position.Y >= float32(g.GameMap.RowSize)
}

func (g *Game) CanMove(player *Player, newX float32, newY float32) bool {
	return !g.collidesWithWalls(Position{newX, newY}) && !g.isOutOfBounds(Position{newX, newY})
}

func (g *Game) collidesWithPlayers(position Position) bool {
	pos := rl.NewRectangle(position.X*65, position.Y*65, 65, 65)

	for _, player := range g.Players {
		playerRect := rl.NewRectangle(player.Position.X*65, player.Position.Y*65, 55, 55)
		if rl.CheckCollisionRecs(pos, playerRect) {
			return true
		}
	}
	return false
}

func (g *Game) collidesWithPowerUp(position Position) bool {
	pos := rl.NewRectangle(position.X*65, position.Y*65, 65, 65)

	for _, powerUp := range g.GameMap.PowerUps {
		powerUpRect := rl.NewRectangle(powerUp.Position.X*65, powerUp.Position.Y*65, 55, 55)
		if rl.CheckCollisionRecs(pos, powerUpRect) {
			return true
		}

	}
	return false
}

func (g *Game) collidesWithBomb(position Position) bool {
	pos := rl.NewRectangle(position.X*65, position.Y*65, 65, 65)

	for _, bomb := range g.GameMap.Bombs {
		bombRect := rl.NewRectangle(bomb.Position.X*65, bomb.Position.Y*65, 55, 55)
		if rl.CheckCollisionRecs(pos, bombRect) {
			return true
		}
	}
	return false
}

func (g *Game) IsValidPosition(ValidPosition Position) bool {
	return !(g.collidesWithWalls(ValidPosition) || g.collidesWithPlayers(ValidPosition) || g.collidesWithPowerUp(ValidPosition) || g.collidesWithBomb(ValidPosition))
}

func (g *Game) GenerateValidPosition(rowSize int, columnSize int) *Position {
	var ValidPosition = getRandomPosition(rowSize, columnSize)
	for !g.IsValidPosition(*ValidPosition) {
		ValidPosition = getRandomPosition(rowSize, columnSize)
	}
	return ValidPosition
}

// Se fija si se tocan no más, habría que arreglarlo para que sea cuando pase medianamente por encima
func (g *Game) IsPowerUpPosition(position Position) *Position {
	// for _, powerUp := range g.GameMap.PowerUps {
	// 	if int(powerUp.Position.X) == int(position.X) && int(powerUp.Position.Y) == int(position.Y) {
	// 		return true
	// 	}
	// }
	// return false
	pos := rl.NewRectangle(position.X*65, position.Y*65, 65, 65)

	for _, powerUp := range g.GameMap.PowerUps {
		powerUpRect := rl.NewRectangle(powerUp.Position.X*65, powerUp.Position.Y*65, 15, 15)
		if rl.CheckCollisionRecs(pos, powerUpRect) {
			return &Position{powerUp.Position.X, powerUp.Position.Y}
		}

	}
	return nil
}

func (g *Game) IsBombPosition(position Position) bool {
	for _, bomb := range g.GameMap.Bombs {
		if bomb.Position == position {
			return true
		}
	}
	return false
}

func (g *Game) PutBomb(player *Player) {
	if player.CanPlantBomb() {
		bomb := NewBomb(player.Position.X, player.Position.Y, 2, *player)
		g.GameMap.PlaceBomb(bomb)
		player.Bombs--
	}
}

func (g *Game) ExplodeBomb(bomb *Bomb) {
	g.GameMap.RemoveBomb(bomb)
	explosion := NewExplosion(bomb.Position, int(bomb.Alcance), *g)
	g.GameMap.Explosions = append(g.GameMap.Explosions, *explosion)

	for _, player := range g.Players {
		if player.ID == bomb.Owner.ID {
			player.Bombs++
		}
	}
}

func (g *Game) TransferPowerUpToPlayer(player *Player, powerUpPosition Position) {
	powerUp := g.GameMap.GetPowerUp(powerUpPosition)

	if powerUp != nil {
		powerUp.SetPowerUpStartTime()
		player.AddPowerUp(*powerUp)
		g.GameMap.RemovePowerUp(powerUpPosition)
	}
}

func (g *Game) GrabPowerUp(playerId string) {
	player := g.Players[playerId]
	powerUpPosition := g.IsPowerUpPosition(*player.Position)
	if powerUpPosition != nil {
		g.TransferPowerUpToPlayer(player, *powerUpPosition)
	}
}

func (g *Game) PowerUpSpawn() {
	if len(g.GameMap.PowerUps) < MAX_POWER_UPS {
		position := g.GenerateValidPosition(g.GameMap.ColumnSize, g.GameMap.RowSize)
		g.GameMap.AddPowerUp(position)
	}

}

func (g *Game) IsFull() bool {
	return len(g.Players) == MAX_PLAYERS
}

func (g *Game) AddPlayer(player *Player) {
	g.Players[player.ID] = player
	color, ok := colors.Dequeue()
	if ok {
		g.PlayerColors[player.ID] = color
	}
}

func (g *Game) GetPlayerColors() map[string]string {
	return g.PlayerColors
}

func (g *Game) GetPlayerColor(id string) string {
	return g.PlayerColors[id]
}

func (g *Game) RemovePlayer(playerID string) {
	delete(g.Players, playerID)
	color := g.GetPlayerColor(playerID)
	colors.Enqueue(color)
}

func (g *Game) Start() {
	g.State = "started"
	ticker := time.NewTicker(POWERUP_SPAWN_TIME * time.Second)
	stopChan = make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("tick")
				g.PowerUpSpawn()
			}
		}
	}()
}

func (g *Game) passRound() {
	if g.Round < MAX_ROUNDS {
		g.Round++
	} else {
		g.State = "finished"
	}
}

func (g *Game) IsEmpty() bool {
	return len(g.Players) == 0
}

func (g *Game) GetAPlayerId() string {
	for key := range g.Players {
		return key
	}
	return ""
}

func (g *Game) Stop() {
	close(stopChan)
	g.State = "stopped"
	log.Println("Game stopped")
}

func (g *Game) Update() {
	now := time.Now()
	for _, bomb := range g.GameMap.Bombs {
		if now.After(bomb.PlantedTime.Add(bomb.ExplodeTime)) {
			g.ExplodeBomb(&bomb)
		}
	}

	for _, explosion := range g.GameMap.Explosions {
		if now.After(explosion.ExplosionTime.Add(1 * time.Second)) {
			g.GameMap.RemoveExplosion(&explosion)
		}
	}

	for _, player := range g.Players {
		for _, powerUp := range player.PowerUps {
			if now.After(powerUp.StartTime.Add(powerUp.ExpireTime * time.Second)) {
				player.RemovePowerUp(powerUp)
			}
		}
	}

}
