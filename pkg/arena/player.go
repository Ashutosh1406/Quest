package arena

import (
	"log"
	"magical-arena/pkg/utilities"
)

type Player interface { //using go interfaces for method extensibility in future...
	//some methods to be called according to tghe functionality

	GetHealth() int //for health check {rest all are self explanatory methods}
	GetStrength() int
	GetAttack() int
	ReduceHealth(amount int) bool // amount is to be calculated according to the problem statement as contains important Game Logic
	IsAlive() bool
}

type BasicPlayer struct { //basic player strcuct for attributes of a player / schema
	health   int
	strength int
	attack   int
}

func NewBasicPlayer(health, strength, attack int) *BasicPlayer {
	return &BasicPlayer{health, strength, attack}
}

//Method defination starts Here :=>

func (p *BasicPlayer) GetHealth() int {
	return p.health
}

func (p *BasicPlayer) GetStrength() int {
	return p.strength
}

func (p *BasicPlayer) GetAttack() int {
	return p.attack
}

func (p *BasicPlayer) ReduceHealth(amount int) bool {
	p.health -= amount
	return p.health > 0
}

func (p *BasicPlayer) IsAlive() bool {
	return p.health > 0
}

// basic validation for custom testing/unit tests values accepted are b/w [1,1000]

func ValidatePlayers(player1, player2 Player) {
	healthDifference := utilities.Abs(player1.GetHealth() - player2.GetHealth())
	strengthDifference := utilities.Abs(player1.GetStrength() - player2.GetStrength())
	attackDifference := utilities.Abs(player1.GetAttack() - player2.GetAttack())

	if healthDifference > 200 || strengthDifference > 30 || attackDifference > 30 {
		log.Fatalf("Huge Difference in player attribues : Health Difference: %d ,  Strength Difference: %d, Attack Difference: %d", healthDifference, strengthDifference, attackDifference)
	}
}

//absolute[abs] moved to utilities as a helper package
