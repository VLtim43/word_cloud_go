package main

import (
	"encoding/json"
	"os"
)

type Word struct {
	Text   string `json:"text"`
	Weight int    `json:"weight"`
}

func readJson(filename string) ([]Word, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var words []Word
	if err := json.Unmarshal(data, &words); err != nil {
		return nil, err
	}

	return words, nil
}
