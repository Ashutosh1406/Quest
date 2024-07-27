
### Files and Directories

- `pkg/arena/match.go`: Contains the `Match` struct and methods to simulate the match.
- `pkg/arena/player.go`: Contains the `Player` interface and its implementation.
- `pkg/arena/utils.go`: Contains utility functions used in the match.
- `test/arena_test.go`: Contains unit tests for the Magical Arena implementation. Covered all the Corner Cases as well
- `cmd/main.go`: Entry point for running the simulation.
- `go.mod` and `go.sum`: Go module files for dependency management.
- `README.md`: This file, providing an overview and instructions.

## How to Run the Code

1. **Clone the repository:**

    ```bash
    cd magical-arena
    ```

2. **Build and run the project:**

    ```bash
    go run main.go
    ```

    This will execute the simulation with the default player attributes.

3. **Run the tests:**

    To run the unit tests and see the output: navigate to root folder and run the below command

    ```bash
    go test -v ./...   
    ```

    This will run all the tests in the `test/` directory and display detailed output indicating whether each test passed or failed.

## Player and Match Implementation

### Player

Players are defined by the following attributes:
- `health`: The player's health points.
- `strength`: The player's defense capability.
- `attack`: The player's attack capability.

### Match

A `Match` struct is used to simulate the battle between two players. The match continues until one player's health drops to 0 or below. The player with lower health attacks first at the start of the match. During each turn, both players roll dice to determine the attack and defense outcomes.

### Utilities

The `utils.go` file contains utility functions such as `RollDice` and `Abs`.

## Early Termination

The match has an early termination condition to speed up the simulation in case one player is clearly dominant. This condition checks if the difference in health and strength between the two players is significantly large, allowing the match to terminate early.

## For Readibility

All the Print Statements are logged after every 10 moves until and unless a winner is decided till the 10th move. 

## Customization

You can customize the player attributes in the `main.go` file to simulate different scenarios.

```go
package main

import (
    "fmt"
    "magical-arena/pkg/arena"
)

func main() {
    playerA := arena.NewBasicPlayer(50, 5, 10)
    playerB := arena.NewBasicPlayer(100, 10, 5)

    match := arena.NewMatch(playerA, playerB)
    match.Start()
}
