package menu

import (
    "fmt"
)

func ShowMainMenu() {
    fmt.Println("Main Menu")
    fmt.Println("1. Play")
    fmt.Println("2. Episodes")
    fmt.Println("3. Quit")

    // input handling
    choice := GetUserInput()
    switch choice {
    case 1:
        SelectEpisode()
    case 2:
        ShowEpisodeMenu()
    case 3:
        QuitGame()
    }
}