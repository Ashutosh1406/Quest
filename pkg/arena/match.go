package arena

type Match struct {
	player1 Player
	player2 Player
}

func NewMatch(player1, player2 Player) *Match {
	return &Match{player1, player2}
}
