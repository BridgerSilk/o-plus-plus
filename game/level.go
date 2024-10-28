package game

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

type Level struct {
    Name        string      `json:"name"`
    PlayerStart Position    `json:"player_start"`
    Tiles       []Tile      `json:"tiles"`
    Obstacles   []Obstacle  `json:"obstacles"`
    Goals       []Position  `json:"goals"`
}

type Position struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type Tile struct {
    Type   string  `json:"type"`
    X      float64 `json:"x"`
    Y      float64 `json:"y"`
    Width  float64 `json:"width"`
    Height float64 `json:"height"`
}

type Obstacle struct {
    Type string  `json:"type"`
    X    float64 `json:"x"`
    Y    float64 `json:"y"`
}

func LoadLevel(filename string) Level {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Failed to read level file: %v", err)
    }

    var level Level
    err = json.Unmarshal(data, &level)
    if err != nil {
        log.Fatalf("Failed to parse level data: %v", err)
    }

    fmt.Printf("Loaded level: %s\n", level.Name)
    return level
}