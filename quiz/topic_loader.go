package quiz

import (
	"encoding/csv"
	"os"
)

type Topic struct {
	ID    string
	Title string
}

// LoadTopics reads topics from quizzes/topics.csv and returns a slice of Topic structs.
func LoadTopics(path string) ([]Topic, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create new reader to read the csv files
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var topics []Topic
	for i, record := range records {
		if i == 0 {
			continue //this is to skip the topic.csv file header
		}
		topics = append(topics, Topic{
			ID:    record[0],
			Title: record[1],
		})
	}

	return topics, nil

}
