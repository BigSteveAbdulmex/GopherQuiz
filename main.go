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

	reader := bufio.NewReader(os.Stdin)

	for {
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
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(topics) {
			fmt.Println("❌ Invalid choice. Please enter a valid topic number.")
			continue
		}

		// Get the selected topic
		selectedTopic := topics[choice-1]
		fmt.Printf("\n✅ You selected: %s\n", selectedTopic.Title)

		// Run quiz session for selected topic
		runQuizLoop(selectedTopic, reader)
	}
}

// runQuizLoop allows restarting same topic, picking a new one, or exiting
func runQuizLoop(topic quiz.Topic, reader *bufio.Reader) {
	for {
		runQuiz(topic, reader) // Start the quiz for the selected topic

		// Ask user what to do after completing the quiz
		fmt.Println("\nWhat would you like to do next?")
		fmt.Println("[1] 🔁 Restart this topic")
		fmt.Println("[2] ➡️ Try another topic")
		fmt.Println("[3] ❌ Exit")

		fmt.Print("\n👉 Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			// Restart current topic
			continue
		case "2":
			// Break this loop, return to topic selection in main
			return
		case "3":
			// Exit the entire program
			fmt.Println("\n👋 Thanks for playing GopherQuiz. See you next time!")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid input. Returning to main menu.")
			return
		}
	}
}

// runQuiz executes the quiz logic with timeout, scoring, and feedback
func runQuiz(topic quiz.Topic, reader *bufio.Reader) {
	// Convert topic title to file name format
	fileName := strings.ToLower(strings.ReplaceAll(topic.Title, " ", "_"))
	filePath := fmt.Sprintf("quizzes/%s.csv", fileName)

	// Load questions from file
	questions, err := quiz.LoadQuestions(filePath)
	if err != nil {
		fmt.Println("❌ Error loading quiz from file:", filePath, "-", err)
		return
	}

	// Set total quiz time limit in seconds
	const totalTime = 210 // 3 minutes, 30 seconds

	// Inform user of time limit
	fmt.Printf("\n⏳ You have %s to complete the quiz. Good luck!\n", utils.FormatDuration(totalTime))

	// Create cancelable context for the timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.SecondsToDuration(totalTime))
	defer cancel()

	score := 0
	utils.StartTimer() // Begin timer

	// Iterate through each question
	for i, question := range questions {
		fmt.Printf("\nQ%d: %s\n", i+1, question.Prompt)
		fmt.Println("1.", question.Options[0])
		fmt.Println("2.", question.Options[1])
		fmt.Println("3.", question.Options[2])
		fmt.Println("4.", question.Options[3])
		fmt.Print("\nYour answer: ")

		answerCh := make(chan string)

		// Read answer in a goroutine to handle timeouts
		go func() {
			answer, _ := reader.ReadString('\n')
			answerCh <- strings.TrimSpace(answer)
		}()

		select {
		case <-ctx.Done():
			// Time expired during question
			fmt.Println("\n⏰ Time's up! You couldn't finish in time.")
			utils.EndTimer()
			showResults(score, len(questions))
			return
		case answer := <-answerCh:
			// Handle user input
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

	// Finished before time ran out
	utils.EndTimer()
	showResults(score, len(questions))
}

// showResults prints the final quiz results
func showResults(score, total int) {
	fmt.Printf("\n🎉 You completed the quiz! Your final score is: %d/%d\n", score, total)
	fmt.Printf("⏰ Time taken: %s\n", utils.FormatDuration(utils.GetElapsedTime()))

	if score == total {
		fmt.Println("🚀 Perfect score! You really know your Go.")
	} else if score == 0 {
		fmt.Println("💡 Don't worry, you'll get better with practice!")
	} else if score < total/2 {
		fmt.Println("📘 Consider reviewing this topic again.")
	} else {
		fmt.Println("👏 Nice effort! Keep improving.")
	}
}
