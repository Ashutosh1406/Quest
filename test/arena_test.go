package test

import (
	"magical-arena/pkg/arena"
	"testing"
)

func TestMoveCounter(t *testing.T) {
	playerA := arena.NewBasicPlayer(50, 5, 10)
	playerB := arena.NewBasicPlayer(100, 10, 5)

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	if match.Moves() == 0 {
		t.Errorf("Move counter should be greater than 0 after the match")
	}
}

func TestEarlyTermination(t *testing.T) {
	playerA := arena.NewBasicPlayer(200, 20, 30) // Clearly stronger player
	playerB := arena.NewBasicPlayer(50, 5, 10)   // Clearly weaker player

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	// Since playerA is clearly dominant, the match should terminate early
	if match.Moves() == 0 {
		t.Errorf("Move counter should be greater than 0 after the match")
	}

	if playerA.IsAlive() && playerB.IsAlive() {
		t.Errorf("One of the players should be dead after early termination")
	}
}

func TestPlayerA50PlayerB100(t *testing.T) {
	playerA := arena.NewBasicPlayer(50, 5, 10)
	playerB := arena.NewBasicPlayer(100, 10, 5)

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	if match.Moves() == 0 {
		t.Errorf("Move counter should be greater than 0 after the match")
	}

	if !playerA.IsAlive() && !playerB.IsAlive() {
		t.Errorf("At least one player should still be alive after the match")
	}
}

func TestExampleTestCase(t *testing.T) {
	playerA := arena.NewBasicPlayer(50, 5, 10)
	playerB := arena.NewBasicPlayer(100, 10, 5)

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	if playerA.IsAlive() && playerB.IsAlive() {
		t.Errorf("One Player should be dead after the Match")
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

func TestSignificantDifferenceInHealth(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when there was a significant difference in health")
		}
	}()

	playerA := arena.NewBasicPlayer(50, 5, 10)
	playerB := arena.NewBasicPlayer(300, 5, 10) // Significant difference in health

	_ = arena.NewMatch(playerA, playerB)
}

func TestSignificantDifferenceInStrength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when there was a significant difference in strength")
		}
	}()

	playerA := arena.NewBasicPlayer(100, 5, 10)
	playerB := arena.NewBasicPlayer(100, 30, 10) // Significant difference in strength

	_ = arena.NewMatch(playerA, playerB)
}

func TestSignificantDifferenceInAttack(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when there was a significant difference in attack")
		}
	}()

	playerA := arena.NewBasicPlayer(100, 5, 10)
	playerB := arena.NewBasicPlayer(100, 5, 40) // Significant difference in attack

	_ = arena.NewMatch(playerA, playerB)
}

func TestEqualStrengthButDifferentHealth(t *testing.T) {
	playerA := arena.NewBasicPlayer(150, 10, 10)
	playerB := arena.NewBasicPlayer(100, 10, 10)

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	if match.Moves() == 0 {
		t.Errorf("Move counter should be greater than 0 after the match")
	}

	if !playerA.IsAlive() && !playerB.IsAlive() {
		t.Errorf("At least one player should still be alive after the match")
	}
}

func TestEqualStrengthAndHealth(t *testing.T) {
	playerA := arena.NewBasicPlayer(100, 10, 10)
	playerB := arena.NewBasicPlayer(100, 10, 10)

	match := arena.NewMatch(playerA, playerB)
	match.Start()

	if match.Moves() == 0 {
		t.Errorf("Move counter should be greater than 0 after the match")
	}

	if !playerA.IsAlive() && !playerB.IsAlive() {
		t.Errorf("At least one player should still be alive after the match")
	}
}
