package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Prompt  string
	Options []string
	Answer  int // 1-based index of the correct option
}

// LoadQuestions loads questions from a given CSV file.
func LoadQuestions(path string) ([]Question, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var questions []Question
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}

		// Parse the correct answer (assuming it's in column 5, index 4)
		answerIndex, err := strconv.Atoi(record[5])
		if err != nil {
			fmt.Printf("⚠️ Error parsing answer on line %d: %v\n", i+1, err)
			continue
		}

		// Make sure the answer is within a valid range (1-4)
		if answerIndex < 1 || answerIndex > 4 {
			fmt.Printf("⚠️ Invalid answer on line %d: %d\n", i+1, answerIndex)
			continue
		}

		// Create the Question struct
		q := Question{
			Prompt:  record[0],
			Options: record[1:5], // Assumes 4 options
			Answer:  answerIndex, // Correct answer (1-based)
		}
		questions = append(questions, q)
	}

	return questions, nil
}
