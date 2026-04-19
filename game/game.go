package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
)

type Game struct {
	SecretNumber int
	Chances      int
	Attempts     int
	Reader       *bufio.Reader
	IsWinning    bool
}

func NewGame(difficulty int, reader *bufio.Reader) *Game {
	return &Game{
		SecretNumber: rand.Intn(100) + 1,
		Chances:      DifficultyToChances[difficulty],
		Reader:       reader,
	}
}

func (g *Game) Play() {
	for g.Chances > 0 {
		fmt.Printf("\nYou have %d chances left. Enter your guess: ", g.Chances)

		guess, err := ReadInt(g.Reader)
		if err != nil {
			fmt.Println("Error: Please enter a numeric value.")
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("Out of range! Please guess between 1 and 100.")
			continue
		}

		g.Attempts++
		g.Chances--

		if guess == g.SecretNumber {
			g.IsWinning = true
			break
		} else if guess > g.SecretNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		}
	}

	g.printResult()
}

func (g *Game) printResult() {
	if g.IsWinning {
		fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts.\n", g.Attempts)
	} else {
		fmt.Printf("\nGame Over! The secret number was %d.\n", g.SecretNumber)
	}
}

func PlayWithReplay(difficulty int, reader *bufio.Reader) {
	for {
		game := NewGame(difficulty, reader)
		game.Play()

		fmt.Print("\nDo you want to play again? (y/n): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" {
			fmt.Println("Thanks for playing! 👋")
			break
		}
	}
}
