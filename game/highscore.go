package game

import (
	"encoding/json"
	"fmt"
	"os"
)

const highScoreFile = "highscore.json"

type HighScores map[int]int

func LoadHighScores() HighScores {
	file, err := os.ReadFile(highScoreFile)
	if err != nil {
		return HighScores{}
	}

	scores := HighScores{}
	if err := json.Unmarshal(file, &scores); err != nil {
		return HighScores{}
	}

	return scores
}

func SaveHighScores(scores HighScores) {
	file, err := os.Create(highScoreFile)
	if err != nil {
		fmt.Println("failed to create highscore file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(scores); err != nil {
		fmt.Println("failed to save high scores:", err)
	}
}
