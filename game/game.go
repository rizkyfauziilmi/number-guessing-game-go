package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	SecretNumber int
	Chances      int
	Attempts     int
	Reader       *bufio.Reader
	IsWinning    bool
	Difficulty   int
	HintUsed     bool
}

func NewGame(difficulty int, reader *bufio.Reader) *Game {
	return &Game{
		SecretNumber: rand.Intn(100) + 1,
		Chances:      DifficultyToChances[difficulty],
		Reader:       reader,
		Difficulty:   difficulty,
	}
}

func (g *Game) shouldGiveHint() bool {
	switch g.Difficulty {
	case Easy:
		return g.Attempts >= 6
	case Medium:
		return g.Attempts >= 3
	case Hard:
		return g.Attempts >= 2
	default:
		return false
	}
}

func (g *Game) giveHint() {
	if g.HintUsed || !g.shouldGiveHint() {
		return
	}

	g.HintUsed = true

	fmt.Println("\n💡 Hint:")

	// Hint 1: selalu ada
	if g.SecretNumber%2 == 0 {
		fmt.Println("- The number is EVEN")
	} else {
		fmt.Println("- The number is ODD")
	}

	// Hint 2: hanya untuk Easy & Medium
	if g.Difficulty != Hard {
		switch {
		case g.SecretNumber <= 25:
			fmt.Println("- Range: 1-25")
		case g.SecretNumber <= 50:
			fmt.Println("- Range: 26-50")
		case g.SecretNumber <= 75:
			fmt.Println("- Range: 51-75")
		default:
			fmt.Println("- Range: 76-100")
		}
	}
}

func (g *Game) Play() {
	startTime := time.Now()

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

		g.giveHint()

		if guess == g.SecretNumber {
			g.IsWinning = true
			break
		} else if guess > g.SecretNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		}
	}

	duration := time.Since(startTime)
	g.printResult(duration)
}

func (g *Game) printResult(duration time.Duration) {
	if g.IsWinning {
		fmt.Printf("\n🎉 Congratulations! You guessed the correct number in %d attempts.\n", g.Attempts)
		fmt.Printf("⏱️ Time taken: %s\n", duration)
	} else {
		fmt.Printf("\n💀 Game Over! The secret number was %d.\n", g.SecretNumber)
		fmt.Printf("⏱️ Time taken: %s\n", duration)
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
