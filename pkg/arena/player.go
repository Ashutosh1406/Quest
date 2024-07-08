package arena

type Player interface { //using go interfaces for method extensibility in future...
	//some methods to be called according to tghe functionality

	GetHealth() int //for health check {rest all are self explanatory methods}
	GetStrength() int
	GetAttack() int
	ReduceHealth(amount int) // amount is to be calculated according to the problem statement as contains important Game Logic
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

func (p *BasicPlayer) ReduceHealth(amount int) {
	p.health -= amount
	// TODO:
}

func (p *BasicPlayer) IsAlive() bool {
	return p.health > 0
}
