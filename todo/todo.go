package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}
	return Todo{
		Text: content,
	}, nil
}
