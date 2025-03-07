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
const BOMB_REACH_MODIFED = 3

const SPEED_INCREMENT = 0.1
const BASE_SPEED = 0

const FRAME_COUNT = 3

var colors = NewQueue[string]()

type GameState string

const (
	NotStarted    GameState = "not-started"
	Started       GameState = "started"
	Finished      GameState = "finished"
	BetweenRounds GameState = "between-rounds"
)

type Game struct {
	State            GameState
	GameId           string
	Round            int8
	Players          map[string]*Player
	PlayerColors     map[string]string
	GameMap          *GameMap
	EliminationOrder []string
	PlayerScores     map[string]int
	CurrentFrame     int
	FrameDuration    time.Duration
	LastFrameTime    time.Time
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
		State:            "not-started",
		GameId:           id,
		Round:            1,
		Players:          make(map[string]*Player),
		PlayerColors:     make(map[string]string),
		GameMap:          GameMap,
		PlayerScores:     make(map[string]int),
		EliminationOrder: []string{},
		CurrentFrame:     0,
		FrameDuration:    200 * time.Millisecond,
		LastFrameTime:    time.Now(),
	}
}

func checkCollision[T Positionable, A Positionable](a T, b A) bool {
	return rl.CheckCollisionRecs(a.GetRect(), b.GetRect())
}

func collidesWithAny[T Positionable, A Positionable](obj T, array []A) bool {
	for _, actual := range array {
		if checkCollision(actual, obj) {
			return true
		}
	}
	return false
}

func (g *Game) IsValidPosition(position Position) bool {
	obj := GameObject{Position: position, Size: 65}
	collidesWithWalls := collidesWithAny(obj, g.GameMap.Walls)
	collidesWithBombs := collidesWithAny(obj, g.GameMap.Bombs)

	players := make([]Positionable, 0, len(g.Players))
	for _, player := range g.Players {
		players = append(players, player)
	}

	collidesWithPlayers := collidesWithAny(obj, players)
	collidesWithPowerUp := collidesWithAny(obj, g.GameMap.PowerUps)
	outOfBounds := g.isOutOfBounds(position)

	return !(collidesWithWalls || collidesWithBombs || collidesWithPlayers || collidesWithPowerUp || outOfBounds)
}

func (g *Game) isOutOfBounds(position Position) bool {
	return position.X < 0 || position.X > float32(g.GameMap.ColumnSize)-1 || position.Y < 0 || position.Y > float32(g.GameMap.RowSize)-1
}

func (g *Game) handleOutOfBounds(position Position) Position {
	newPosition := position

	if position.X < 0 {
		newPosition.X = 0
	}
	if position.X >= float32(g.GameMap.ColumnSize)-1 {
		newPosition.X = float32(g.GameMap.ColumnSize) - 1
	}
	if position.Y < 0 {
		newPosition.Y = 0
	}
	if position.Y >= float32(g.GameMap.RowSize)-1 {
		newPosition.Y = float32(g.GameMap.RowSize) - 1
	}
	return newPosition
}

func (g *Game) CanMove(player *Player, newX float32, newY float32) bool {
	newPlayerPos := Player{Position: &Position{newX, newY}}
	return !collidesWithAny(newPlayerPos, g.GameMap.Walls)
}

func (g *Game) MovePlayer(player *Player, newX float32, newY float32) {
	newPosition := g.handleOutOfBounds(Position{newX, newY})

	if g.CanMove(player, newPosition.X, newPosition.Y) {
		player.Position.X = newPosition.X
		player.Position.Y = newPosition.Y
	} else {
		player.Speed.X, player.Speed.Y = BASE_SPEED, BASE_SPEED
	}
	g.GrabPowerUp(player.ID)
}

func (g *Game) GenerateValidPosition(rowSize int, columnSize int) *Position {
	var ValidPosition = getRandomPosition(rowSize, columnSize)
	for !g.IsValidPosition(*ValidPosition) {
		ValidPosition = getRandomPosition(rowSize, columnSize)
	}
	return ValidPosition
}

func (g *Game) IsPowerUpPosition(obj Positionable) *Position {
	for _, powerUp := range g.GameMap.PowerUps {
		if checkCollision(powerUp, obj) {
			return &Position{powerUp.Position.X, powerUp.Position.Y}
		}
	}
	return nil
}

func (g *Game) PutBomb(player *Player) {
	if player.CanPlantBomb() {
		x := float32(math.Round(float64(player.Position.X)))
		y := float32(math.Round(float64(player.Position.Y)))
		bomb := NewBomb(x, y, player.BombReach, *player)
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
			log.Println("Player has bombs: %i", player.Bombs)
			player.AddBomb()
		}
	}
}

func (g *Game) TransferPowerUpToPlayer(player *Player, powerUpPosition Position) {
	powerUp := g.GameMap.GetPowerUp(powerUpPosition)

	if powerUp != nil {
		powerUp.SetPowerUpStartTime()
		player.AddPowerUp(*powerUp)
		player.ApplyPowerUpBenefit(powerUp.Name)
		g.GameMap.RemovePowerUp(powerUpPosition)
	}
}

func (g *Game) GrabPowerUp(playerId string) {
	player := g.Players[playerId]
	powerUpPosition := g.IsPowerUpPosition(player)
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
	return len(g.Players) >= MAX_PLAYERS
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

func (g *Game) GetPlayer(playerID string) *Player {
	return g.Players[playerID]
}

func (g *Game) RemovePlayer(playerID string) {
	delete(g.Players, playerID)
	color := g.GetPlayerColor(playerID)
	colors.Enqueue(color)
}

func (g *Game) Start() {
	g.State = "started"
	ticker := time.NewTicker(POWERUP_SPAWN_TIME * time.Second)

	go func() {
		for range ticker.C {
			log.Println("tick")
			g.PowerUpSpawn()
		}
	}()
}

func (g *Game) assignScores() {
	score := 12
	for i := len(g.EliminationOrder) - 1; i >= 0; i-- {
		playerID := g.EliminationOrder[i]
		g.PlayerScores[playerID] = score + g.PlayerScores[playerID]
		score -= 3
	}

	g.EliminationOrder = []string{}
}

func (g *Game) passRound() {
	g.assignScores()
	g.State = "between-rounds"
	g.Round++
	g.startRound()
}

func (g *Game) endGame() {
	g.assignScores()
	g.State = Finished
}

func (g *Game) endRound() {
	if g.Round < MAX_ROUNDS {
		g.passRound()
	} else {
		g.endGame()
	}
}

func (g *Game) initializePlayersForRound() {
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
}

func (g *Game) startRound() {
	g.GameMap = GetRoundGameMap(g.Round)
	g.initializePlayersForRound()
	g.State = "started"
}

func (g *Game) IsEmpty() bool {
	return len(g.Players) == 0
}

func (g *Game) RandomPlayerId() string {
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

		for _, neighbor := range g.getTileNeighbors(current) {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return nil
}

func (g *Game) getTileNeighbors(position Position) []Position {
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
	g.State = "stopped"
	log.Println("Game stopped")
}

func (g *Game) verifyExplodingBombs(now time.Time) {
	for _, bomb := range g.GameMap.Bombs {
		if now.After(bomb.PlantedTime.Add(bomb.ExplodeTime)) {
			g.ExplodeBomb(&bomb)
		}
	}
}

func (g *Game) handleExplosion(explosion *Explosion) {
	for _, player := range g.Players {
		if player.Lives == 0 {
			continue
		}
		if explosion.IsTileInRange(*player.Position) && !player.Invencible && !explosion.IsPlayerAlreadyAffected(player.ID) {
			explosion.AddAffectedPlayer(player.ID)
			log.Println("Player ", player.ID, " has {", player.Lives, "} lives left")
			lives_left := player.LoseHealth()
			log.Println("Now player ", player.ID, " has {", player.Lives, "} lives left")
			if lives_left == 0 {
				g.EliminationOrder = append(g.EliminationOrder, player.ID)
			}
		}
	}

	for _, bomb := range g.GameMap.Bombs {
		if explosion.IsTileInRange(bomb.Position) {
			g.ExplodeBomb(&bomb)
		}
	}
}

func (g *Game) verifyExplosions() {
	var expiredExplosions []int

	for i := range g.GameMap.Explosions {
		explosion := &g.GameMap.Explosions[i]
		if explosion.IsExpired() {
			expiredExplosions = append(expiredExplosions, i)
			continue
		}
		g.handleExplosion(explosion)
	}

	g.GameMap.RemoveExplosions(expiredExplosions)
}

func (g *Game) updatePowerUps(now time.Time) {
	for _, player := range g.Players {
		if player.Lives == 0 {
			continue
		}
		for _, powerUp := range player.PowerUps {
			if !powerUp.StartTime.IsZero() {
				if now.After(powerUp.ExpireTime) {
					log.Println("PowerUp expired")
					player.RemovePowerUpBenefit(powerUp.Name)
					player.RemovePowerUp(powerUp)
				}
			}
		}
	}
}

func (g *Game) shouldEndRound() bool {
	return len(g.EliminationOrder) == (len(g.Players)-1) && g.State != "not-started"
}

func (g *Game) Update() {
	if time.Since(g.LastFrameTime) >= g.FrameDuration {
		g.CurrentFrame = (g.CurrentFrame + 1) % FRAME_COUNT
		g.LastFrameTime = time.Now()
	}

	now := time.Now()
	g.verifyExplodingBombs(now)
	g.verifyExplosions()
	g.updatePowerUps(now)

	if len(g.Players) == 1 && g.State != "not-started" {
		g.endGame()
	} else if g.shouldEndRound() {
		g.endRound()
	}
}
