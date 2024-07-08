package test

import (
	"magical-arena/pkg/arena"
	"magical-arena/pkg/utilities"
	"testing"
)

func TestRollDice(t *testing.T) {
	rollValue := utilities.RollDice()

	if rollValue < 1 || rollValue > 6 {
		t.Errorf("RollDice() = %d not returning value b/w 1 and 6 [valid dice count Not Found]", rollValue)
	}
}

func TestAttack(t *testing.T) {
	attacker := arena.NewBasicPlayer(50, 5, 10)
	defender := arena.NewBasicPlayer(100, 10, 5)

	m := arena.NewMatch(attacker, defender)
	m.AttackSequence(attacker, defender)

	if defender.GetHealth() > 100 {
		t.Errorf("Defender Health %d , should not increase", defender.GetHealth())
	}
}

func TestSimulateBattle(t *testing.T) {
	playerA := arena.NewBasicPlayer(50, 5, 10)
	playerB := arena.NewBasicPlayer(100, 10, 5)

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	if playerA.IsAlive() && playerB.IsAlive() {
		t.Errorf("One Player should be dead after the Match")
	}
}

func TestValidatePlayers(t *testing.T) {
	// Test with acceptable differences
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ValidatePlayers() should not panic for acceptable differences")
		}
	}()
	playerA := arena.NewBasicPlayer(100, 10, 20)
	playerB := arena.NewBasicPlayer(120, 15, 25)
	arena.ValidatePlayers(playerA, playerB)

	// Test with significant differences
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ValidatePlayers() should panic for significant differences")
		}
	}()
	playerA = arena.NewBasicPlayer(1000, 50, 100)
	playerB = arena.NewBasicPlayer(100, 5, 10)
	arena.ValidatePlayers(playerA, playerB)
}
