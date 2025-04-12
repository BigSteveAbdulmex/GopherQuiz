# GopherQuiz

GopherQuiz is a command-line-based quiz app built in Go to test and reinforce your understanding of Go programming — from beginner syntax to advanced topics like concurrency, pointers, and generics.

Whether you're just starting out or brushing up on your skills, GopherQuiz helps you sharpen your knowledge in an interactive and timed environment.

## 💡 Features

- Quiz questions organized by Go topics (syntax, structs, pointers, goroutines, etc.)
- CSV-based question bank for easy editing and expansion
- Topic-wise and overall quiz modes
- Accurate scoring and time tracking
- Designed for both beginners and advanced learners

## 🛠️ How to Run

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/BigSteveAbdulmex/GopherQuiz.git
   cd GopherQuiz
   ```

2. Install dependencies (Go modules):
   ```bash
   go mod tidy
   ```

3. Run the quiz:
   ```bash
   go run main.go
   ```

4. Select a topic and start the quiz! Answer questions, track your score, and test your knowledge.

## 📝 How to Add New Questions

- The question bank is stored in CSV files located in the `quizzes/` directory.
- To add a new topic, simply create a new CSV file for the topic (e.g., `concurrency.csv`).
- The format for each CSV file is:
  ```
  question, option1, option2, option3, option4, correct_option_number
  ```

---