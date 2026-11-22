package main

import (
	"encoding/json"
	"os"
	"strings"
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

	for i := range words {
		words[i].Text = strings.ToUpper(words[i].Text)
	}

	return words, nil
}
