package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Quiz contains all questions and responses
type Quiz struct {
	Question      string
	CorrectAnswer string
	ClientAnswer  string
	IsCorrect     bool
}

// Quizes contains a bunch of quiz
type Quizes struct {
	Quizes []Quiz
}

// NewQuiz implements a quiz based on a csv file
func NewQuiz(file string) (*Quizes, error) {
	var exercice Quizes
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range records {
		exercice.Quizes = append(exercice.Quizes, Quiz{Question: row[0], CorrectAnswer: row[1]})
	}
	return &exercice, err
}

// IsClientAnswerCorrect tells whether the quiz got the right answer
func (q Quiz) IsClientAnswerCorrect() bool {
	result := false

	if strings.Compare(q.ClientAnswer, q.CorrectAnswer) == 0 {
		return true
	}

	return result
}

// Run quiz
func (q Quiz) Run() *Quiz {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s=", q.Question)
	result, _ := reader.ReadString('\n')
	q.ClientAnswer = strings.TrimSpace(strings.TrimSuffix(result, "\n"))
	q.IsCorrect = q.IsClientAnswerCorrect()
	return &q
}

// CountWrongAnswer counts the number of wrong answer
func (q Quizes) CountWrongAnswer() int {
	var result int
	for _, item := range q.Quizes {
		if !item.IsCorrect {
			result++
		}
	}
	return result
}
