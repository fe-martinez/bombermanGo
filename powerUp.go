package main

type PowerUpType int8

const (
	Invencibilidad PowerUpType = iota
	CaminarSobreParedes
	BombasRodantes
	AlcanceMejorado
	PotenciaMejorada
	MasBombasEnSimultaneo
)

type PowerUp struct {
	Position Position
	name     PowerUpType
	timer    float32
}
