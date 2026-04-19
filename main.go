package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rizkyfauziilmi/number-guessing-game-go/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	game.DisplayWelcomeMessage()

	var difficulty int

	for {
		fmt.Print("Enter your choice: ")

		input, err := game.ReadInt(reader)
		if err != nil {
			fmt.Println("Error: Please enter a valid number (1, 2, or 3).")
			continue
		}

		difficulty = input

		if difficulty >= game.Easy && difficulty <= game.Hard {
			game.DisplayStartMessage(game.DifficultyToString[difficulty])
			break
		}

		fmt.Println("Invalid choice. Please select 1, 2, or 3.")
	}

	game.PlayWithReplay(difficulty, reader)
}
