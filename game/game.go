package game

import (
    "fmt"
)

var CurrentLevel Level
var Player Player

func StartLevel(levelName string) {
    CurrentLevel = LoadLevel(fmt.Sprintf("assets/levels/%s.json", levelName))
    
    Player = Player{
        X: CurrentLevel.PlayerStart.X,
        Y: CurrentLevel.PlayerStart.Y,
    }

    fmt.Printf("Starting level: %s\n", CurrentLevel.Name)
    GameLoop()
}