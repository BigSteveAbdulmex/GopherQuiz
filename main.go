package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BigSteveAbdulmex/GopherQuiz/quiz"
	"github.com/BigSteveAbdulmex/GopherQuiz/ui"
	"github.com/BigSteveAbdulmex/GopherQuiz/utils"
)

func main() {
	// Show welcome message
	ui.DisplayWelcomeMessage()

	// Load the list of quiz topics from the file
	topics, err := quiz.LoadTopics("quizzes/topics.csv")
	if err != nil {
		fmt.Println("❌ Error loading topics:", err)
		return
	}

	// Show all available topics to the user
	fmt.Println("\nSelect a Topic: ")
	for i, t := range topics {
		fmt.Printf("[%d] %s\n", i+1, t.Title)
	}

	// Ask user to choose a topic by number
	fmt.Print("\n👉 Select a topic by number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(topics) {
		fmt.Println("❌ Invalid choice. Please enter a valid topic number.")
		return
	}

	// Get the selected topic
	selectedTopic := topics[choice-1]
	fmt.Printf("\n✅ You selected: %s\n", selectedTopic.Title)

	// Convert topic title to file name format
	fileName := strings.ToLower(strings.ReplaceAll(selectedTopic.Title, " ", "_"))
	filePath := fmt.Sprintf("quizzes/%s.csv", fileName)
	questions, err := quiz.LoadQuestions(filePath)
	if err != nil {
		fmt.Println("❌ Error loading quiz from file:", filePath, "-", err)
		return
	}

	// Set how much time (in seconds) user has to finish the quiz
	const totalTime = 210 // 3 minutes, 30 seconds

	// Show how much time user has
	fmt.Printf("\n⏳ You have %s to complete the quiz. Good luck!\n", utils.FormatDuration(totalTime))

	// Start a timer that cancels when time runs out
	ctx, cancel := context.WithTimeout(context.Background(), utils.SecondsToDuration(totalTime))
	defer cancel()

	// Track the score
	score := 0

	// Start the timer for tracking how long user took
	utils.StartTimer()

	// Go through each question one by one
	for i, question := range questions {
		fmt.Printf("\nQ%d: %s\n", i+1, question.Prompt)
		fmt.Println("1.", question.Options[0])
		fmt.Println("2.", question.Options[1])
		fmt.Println("3.", question.Options[2])
		fmt.Println("4.", question.Options[3])

		fmt.Print("\nYour answer: ")

		answerCh := make(chan string)

		// Run user input in a separate goroutine
		go func() {
			answer, _ := reader.ReadString('\n')
			answerCh <- strings.TrimSpace(answer)
		}()

		select {
		case <-ctx.Done():
			// If time runs out while waiting for input
			fmt.Println("\n⏰ Time's up! You couldn't finish in time.")
			showResults(score, len(questions))
			return
		case answer := <-answerCh:
			// Parse and check user's answer
			answerChoice, err := strconv.Atoi(answer)
			if err != nil || answerChoice < 1 || answerChoice > 4 {
				fmt.Println("❌ Invalid answer choice.")
				continue
			}
			if answerChoice == question.Answer {
				score++
			}
		}
	}

	// Quiz finished before time ran out
	utils.EndTimer()
	showResults(score, len(questions))
}

// Print final results
func showResults(score, total int) {
	fmt.Printf("\n🎉 You completed the quiz! Your final score is: %d/%d\n", score, total)
	fmt.Printf("⏰ Time taken: %s\n", utils.FormatDuration(utils.GetElapsedTime()))
}
