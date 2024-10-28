package menu

import (
    "fmt"
    "oplusplus/game"
)

func ShowEpisodeMenu() {
    fmt.Println("Select Episode")

    // ep list
    episodes := []string{"Episode 1", "Episode 2", "Episode 3"}
    for i, episode := range episodes {
        fmt.Printf("%d. %s\n", i+1, episode)
    }

    choice := GetUserInput()
    if choice >= 1 && choice <= len(episodes) {
        game.StartLevel(episodes[choice-1])
    }
}