package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BigSteveAbdulmex/GopherQuiz/quiz"
	"github.com/BigSteveAbdulmex/GopherQuiz/ui"
)

func main() {
	// Display welcome message
	ui.DisplayWelcomeMessage()

	// Load available topics
	topics, err := quiz.LoadTopics("quizzes/topics.csv")
	if err != nil {
		fmt.Println("❌ Error loading topics:", err)
		return
	}

	// Display available topics to user
	fmt.Println("\nSelect a Topic: ")
	for i, t := range topics {
		fmt.Printf("[%d] %s\n", i+1, t.Title)
	}

	// user input
	fmt.Print("\n👉 Select a topic by number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(topics) {
		fmt.Println("❌ Invalid choice. Please enter a valid topic number.")
		return
	}

	selectedTopic := topics[choice-1]
	fmt.Printf("\n✅ You selected: %s\n", selectedTopic.Title)

	// Convert topic title to lowercase and underscores
	fileName := strings.ToLower(strings.ReplaceAll(selectedTopic.Title, " ", "_"))
	filePath := fmt.Sprintf("quizzes/%s.csv", fileName)
	questions, err := quiz.LoadQuestions(filePath)
	if err != nil {
		fmt.Println("❌ Error loading quiz from file:", filePath, "-", err)
		return
	}

	// Start quiz and track score
	score := 0
	for i, question := range questions {
		fmt.Printf("\nQ%d: %s\n", i+1, question.Prompt)
		fmt.Println("1.", question.Options[0])
		fmt.Println("2.", question.Options[1])
		fmt.Println("3.", question.Options[2])
		fmt.Println("4.", question.Options[3])

		// Get user answer
		fmt.Print("\nYour answer: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answerChoice, err := strconv.Atoi(answer)
		if err != nil || answerChoice < 1 || answerChoice > 4 {
			fmt.Println("❌ Invalid answer choice.")
			continue
		}

		// Check if the answer is correct
		if answerChoice == question.Answer {
			score++
		}
	}

	// Show final score
	fmt.Printf("\n🎉 You completed the quiz! Your final score is: %d/%d\n", score, len(questions))
}
