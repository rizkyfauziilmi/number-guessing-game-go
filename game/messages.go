package game

import "fmt"

func formatScore(score int) string {
	if score == 0 {
		return "-"
	}
	return fmt.Sprintf("%d", score)
}

func DisplayWelcomeMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println()

	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println()
}

func DisplayStartMessage(difficulty string) {
	scores := LoadHighScores()

	fmt.Println()
	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty)

	fmt.Println("🏆 Best Scores:")
	fmt.Printf("- Easy: %s\n", formatScore(scores[Easy]))
	fmt.Printf("- Medium: %s\n", formatScore(scores[Medium]))
	fmt.Printf("- Hard: %s\n", formatScore(scores[Hard]))

	fmt.Println("\nLet's start the game!")
}
