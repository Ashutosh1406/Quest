package arena

import (
	"fmt"
	"magical-arena/pkg/utilities"
)

type Match struct {
	player1 Player
	player2 Player
	moves   int
	//moves int // for future extensibilty it will store the moves taken to end the game and printing the statements for every 10 moves redundant print opertaion is Ignored
}

const ( //for better readibilty
	MoveFormat    = "Move %d: Attacker rolls %d, Defender rolls %d\n"
	DamageFormat  = "Attack damage: %d, Defense strength: %d\n"
	HealthFormat  = "Defender health reduced by %d to %d\n\n"
	WinnerFormat  = "Player %d wins! Total moves: %d\n"
	InitialHealth = "Initial Health: Player 1 = %d, Player 2 = %d\n\n"
)

func NewMatch(player1, player2 Player) *Match {
	ValidatePlayers(player1, player2)
	return &Match{player1: player1, player2: player2, moves: 0}
}

func (m *Match) Moves() int {
	return m.moves
}
func (m *Match) Start() {
	fmt.Printf(InitialHealth, m.player1.GetHealth(), m.player2.GetHealth())
	for m.player1.IsAlive() && m.player2.IsAlive() {
		if m.player1.GetHealth() < m.player2.GetHealth() {
			if !m.AttackSequence(m.player1, m.player2) {
				fmt.Printf(WinnerFormat, 1, m.moves)
				return
			}
			if !m.AttackSequence(m.player2, m.player1) {
				fmt.Printf(WinnerFormat, 2, m.moves)
				return
			}
		} else {
			if !m.AttackSequence(m.player2, m.player1) {
				fmt.Printf(WinnerFormat, 2, m.moves)
				return
			}
			if !m.AttackSequence(m.player1, m.player2) {
				fmt.Printf(WinnerFormat, 1, m.moves)
				return
			}
		}

		// Early termination check
		if m.shouldTerminateEarly() {
			if m.player1.GetHealth() > m.player2.GetHealth() {
				fmt.Printf("Early Termination: Player 1 is likely to win! Total moves: %d\n", m.moves)
			} else {
				fmt.Printf("Early Termination: Player 2 is likely to win! Total moves: %d\n", m.moves)
			}
			return
		}
	}
}

func (m *Match) AttackSequence(attacker, defender Player) bool {
	m.moves++ //counting the moves until the winner is decided

	attackRoll := utilities.RollDice()
	defenseRoll := utilities.RollDice()

	attackDamage := attacker.GetAttack() * attackRoll
	defenseStrength := defender.GetStrength() * defenseRoll

	damageTaken := max(0, attackDamage-defenseStrength)
	isAlive := defender.ReduceHealth(damageTaken)

	if m.moves%10 == 0 {
		fmt.Printf(MoveFormat, m.moves, attackRoll, defenseRoll)
		fmt.Printf(DamageFormat, attackDamage, defenseStrength)
		fmt.Printf(HealthFormat, damageTaken, defender.GetHealth())
	}

	return isAlive
}

func (m *Match) shouldTerminateEarly() bool {
	healthDifference := utilities.Abs(m.player1.GetHealth() - m.player2.GetHealth())
	strengthDifference := utilities.Abs(m.player1.GetStrength() - m.player2.GetStrength())

	// If one player's health and strength are significantly higher (feature can be waived off)
	if healthDifference > 100 && strengthDifference > 10 {
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
