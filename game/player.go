package game

import (
    "fmt"
)

func GameLoop() {
    running := true
    for running {
        Player.Update()
        fmt.Printf("Player Position: (%.2f, %.2f)\n", Player.X, Player.Y)
    }
}