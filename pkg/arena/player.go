package arena

type Player interface { //using go interfaces for extensibility in future...
	//some methods to be called according to tghe functionality
}

type BasicPlayer struct { //basic player strcuct for attributes of a player / schema
	health   int
	strength int
	attack   int
}

func NewBasicPlayer(health, strength, attack int) *BasicPlayer {
	return &BasicPlayer{health, strength, attack}
}
