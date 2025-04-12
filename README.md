# GopherQuiz

GopherQuiz is a command-line based quiz app built in Go to test and reinforce your understanding of Go programming - from beginner syntax to advanced topics like concurrency, pointers, and generics.

Whether you're just starting out or brushing up as an experienced Gopher, GopherQuiz helps you sharpen your skills in an interactive and timed environment.

## 💡 Features

- Quiz questions organized by Go topics (syntax, structs, pointers, goroutines, etc.)
- CSV-based question bank for easy editing and expansion
- Topic-wise and overall quiz modes
- Accurate scoring and time tracking
- Designed for both beginners and advanced learners

## 📁 Project Structure

```bash
gopherquiz/
├── cmd/
│   └── quiz/
│       └── main.go
├── internal/
│   ├── quiz/
│   │   ├── loader.go
│   │   ├── runner.go
│   │   └── timer.go
│   └── utils/
│       └── csv.go
├── questions/
│   ├── syntax.csv
│   ├── pointers.csv
│   └── concurrency.csv
└── README.md
