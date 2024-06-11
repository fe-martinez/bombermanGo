package model

import (
	"log"
	"math"
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

type GameState string

const (
	NotStarted    GameState = "not-started"
	Started       GameState = "started"
	Finished      GameState = "finished"
	BetweenRounds GameState = "between-rounds"
)

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
	return position.X < 0.0 || position.X >= float32(g.GameMap.ColumnSize) || position.Y < 0.0 || position.Y >= float32(g.GameMap.RowSize)
}

func (g *Game) CanMove(player *Player, newX float32, newY float32) bool {
	return !g.collidesWithWalls(Position{newX, newY}) && !g.isOutOfBounds(Position{newX, newY})
}

func (g *Game) MovePlayer(player *Player, newX float32, newY float32) {
	if g.CanMove(player, newX, newY) {
		player.Position.X = newX
		player.Position.Y = newY
	} else {
		player.Speed.X, player.Speed.Y = BASE_SPEED, BASE_SPEED
	}
	g.GrabPowerUp(player.ID)
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
	if g.GameMap.isUnbreakableWall(ValidPosition) || g.GameMap.isBreakableWall(ValidPosition) {
		return false
	}

	return !(g.collidesWithWalls(ValidPosition) || g.collidesWithPlayers(ValidPosition) || g.collidesWithPowerUp(ValidPosition) || g.collidesWithBomb(ValidPosition))
}

func (g *Game) GenerateValidPosition(rowSize int, columnSize int) *Position {
	var ValidPosition = getRandomPosition(rowSize, columnSize)
	for !g.IsValidPosition(*ValidPosition) {
		ValidPosition = getRandomPosition(rowSize, columnSize)
	}
	return ValidPosition
}

func (g *Game) IsPowerUpPosition(position Position) *Position {
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
		x := float32(math.Round(float64(player.Position.X)))
		y := float32(math.Round(float64(player.Position.Y)))
		bomb := NewBomb(x, y, 2, *player)
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
		log.Println("Power up start time is setted")
		player.AddPowerUp(*powerUp)
		g.ApplyPowerUpBenefit(powerUp.name, player.ID)
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

func (g *Game) ApplyPowerUpBenefit(powerUp PowerUpType, playerID string) {
	switch powerUp {
	case Invencibilidad:
		log.Println("Invencibilidad")
		g.Players[playerID].Invencible = true
	case CaminarSobreParedes:
		log.Println("Caminar sobre paredes not yet implemented")
	case MasBombasEnSimultaneo:
		log.Println("Mas bombas en simultaneo")
		g.Players[playerID].Bombs = 2
	case AlcanceMejorado:
		log.Println("Alcance mejorado")
		for _, bomb := range g.GameMap.Bombs {
			if bomb.IsOwner(playerID) {
				bomb.Alcance = 3
			}
		}
	default:
	}
}

func (g *Game) RemovePowerUpBenefit(powerUp PowerUpType, playerID string) {
	switch powerUp {
	case Invencibilidad:
		log.Println("Removiendo invencibilidad")
		g.Players[playerID].Invencible = false
	case CaminarSobreParedes:
		log.Println("Caminar sobre paredes not yet implemented")
	case MasBombasEnSimultaneo:
		log.Println("Removiendo mas bombas en simultaneo")
		g.Players[playerID].Bombs = 1
	case AlcanceMejorado:
		log.Println("Removiendo alcance mejorado")
		for _, bomb := range g.GameMap.Bombs {
			if bomb.IsOwner(playerID) {
				bomb.Alcance = 1
			}
		}
	default:
	}
}

func (g *Game) endRound() {
	if g.Round < MAX_ROUNDS {
		g.State = "between-rounds"
		g.Round++
		g.startRound()
	} else {
		g.State = "finished"
	}
}

func (g *Game) startRound() {
	g.GameMap = GetRoundGameMap(g.Round)
	count := 0
	for _, player := range g.Players {
		player.Lives = 2
		player.Position = g.GetPlayerPosition(count)
		player.Invencible = false
		player.Bombs = 1
		player.PowerUps = []PowerUp{}
		player.Speed = Speed{X: 0.0, Y: 0.0}
		count++
	}

	g.State = "started"
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

func (g *Game) GetPlayerPosition(playerIndex int) *Position {
	corners := []Position{
		{X: 0, Y: 0},
		{X: float32(g.GameMap.ColumnSize - 1), Y: 0},
		{X: 0, Y: float32(g.GameMap.RowSize - 1)},
		{X: float32(g.GameMap.ColumnSize - 1), Y: float32(g.GameMap.RowSize - 1)},
	}

	if playerIndex < 0 || playerIndex >= len(corners) {
		return nil
	}

	start := corners[playerIndex]
	return g.findNearestFreeSpace(start)
}

func (g *Game) findNearestFreeSpace(start Position) *Position {
	queue := []Position{start}
	visited := make(map[Position]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if g.IsValidPosition(current) {
			log.Println(current)
			return &current
		}

		for _, neighbor := range g.getNeighbors(current) {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return nil
}

func (g *Game) getNeighbors(position Position) []Position {
	neighbors := []Position{}
	if position.X-1 >= 0 {
		neighbors = append(neighbors, Position{X: position.X - 1, Y: position.Y})
	}
	if position.X+1 < float32(g.GameMap.ColumnSize) {
		neighbors = append(neighbors, Position{X: position.X + 1, Y: position.Y})
	}
	if position.Y-1 >= 0 {
		neighbors = append(neighbors, Position{X: position.X, Y: position.Y - 1})
	}
	if position.Y+1 < float32(g.GameMap.RowSize) {
		neighbors = append(neighbors, Position{X: position.X, Y: position.Y + 1})
	}
	return neighbors
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

	for i := range g.GameMap.Explosions {
		explosion := &g.GameMap.Explosions[i]
		if explosion.IsExpired() {
			g.GameMap.RemoveExplosion(explosion)
		}

		for _, player := range g.Players {
			if explosion.IsTileInRange(*player.Position) && !player.Invencible && !explosion.IsPlayerAlreadyAffected(player.ID) {
				log.Println("Player affected by explosion")
				explosion.AddAffectedPlayer(player.ID)
				player.LoseHealth()
			}
		}
	}

	for _, player := range g.Players {
		for _, powerUp := range player.PowerUps {
			log.Println("power up start time: ", powerUp.StartTime, "power up expire time: ", powerUp.ExpireTime, "now: ", now)
			if !powerUp.StartTime.IsZero() {
				if now.After(powerUp.StartTime.Add(powerUp.ExpireTime * time.Second)) {
					log.Println("PowerUp expired")
					g.RemovePowerUpBenefit(powerUp.name, player.ID)
					player.RemovePowerUp(powerUp)
				}
			}
		}
	}

	// Check if there is a winner
	deadPlayers := 0
	for _, player := range g.Players {
		if player.Lives == 0 {
			deadPlayers++
		}

		if deadPlayers == len(g.Players) {
			g.endRound()
		}
	}

}
