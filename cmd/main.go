package main

import (
	"magical-arena/pkg/arena"
)

func main() {

	//calling for new player properties where:
	//(health, strength, attack) are propeties infering to struct [BasicPlayer]

	//using player "A" and "B" as a convention for game battle

	playerA := arena.NewBasicPlayer(50, 5, 10)
	playerB := arena.NewBasicPlayer(100, 10, 5)

}
