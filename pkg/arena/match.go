package arena

import (
	"fmt"
	"magical-arena/pkg/utilities"
)

type Match struct {
	player1 Player
	player2 Player
}

func NewMatch(player1, player2 Player) *Match {
	ValidatePlayers(player1, player2)
	return &Match{player1, player2}
}

func (m *Match) Start() {
	for m.player1.IsAlive() && m.player2.IsAlive() { //checking that both should be alive otherwise terminate the loop / game

		if m.player1.GetHealth() < m.player2.GetHealth() {
			m.AttackSequence(m.player1, m.player2)
			if !m.player2.IsAlive() { // IsAlive returns a boolean true or false
				fmt.Println("Player '1' Wins the GAME , CONGRATULATIONS PLAYER 1")
				return
			}
			m.AttackSequence(m.player2, m.player1)
			if !m.player1.IsAlive() {
				fmt.Println("Player '2' Wins the GAME . CONGRATULATIONS PLAYER 2")
				return
			}
		} else {
			m.AttackSequence(m.player2, m.player1)
			if !m.player1.IsAlive() {
				fmt.Println("Player '2' wins the GAMW , CONGRATULATIONS PLAYER 2 ")
				return
			}
			m.AttackSequence(m.player1, m.player2)
			if !m.player2.IsAlive() {
				fmt.Println("Player '1' Wins the GAME , CONGRATULATIONS PLAYER 1")
				return
			}
		}
	}
}

func (m *Match) AttackSequence(attacker, defender Player) {
	attackerRolles := utilities.RollDice()
	defenderRolles := utilities.RollDice()

	attackDamage := attacker.GetAttack() * attackerRolles
	defenseStrength := defender.GetStrength() * defenderRolles

	damageTaken := max(0, attackDamage-defenseStrength)
	defender.ReduceHealth(damageTaken)

	fmt.Printf("Attacker rolls %d , Defender rolls %d\n ", attackerRolles, defenderRolles)
	fmt.Printf("Attacker Damage: %d , Defender Strength: %d\n", attackDamage, defenseStrength)
	fmt.Printf("Defender Health Reduced By %d to %d \n\n", damageTaken, defender.GetHealth())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
