package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type parsedQuestionAnswer struct {
	question, answer string
}

func main() {

	quizLocation := parseFlags()
	questions := getQuestions(quizLocation)
	correctAnswers := askQuiz(questions)
	fmt.Printf("You've reached the end of the quiz!\n You scored %v out of %v\n", correctAnswers, len(questions))
}

func getQuestions(location *string) []parsedQuestionAnswer {

	file, err := os.Open(*location)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	var allQuestions []parsedQuestionAnswer

	for scanner.Scan() {
		question := parseQuestion(scanner.Text())
		allQuestions = append(allQuestions, question)
	}
	return allQuestions
}

func parseQuestion(line string) parsedQuestionAnswer {
	q := parsedQuestionAnswer{}

	var splitString []string
	splitString = strings.Split(line, ",")

	if len(splitString) > 2 {
		panic("More than two parts two your question!")
	}
	q.question = splitString[0]
	q.answer = splitString[1]

	return q
}

func askQuiz(questions []parsedQuestionAnswer) int {
	correctAnswers := 0

	for i, item := range questions {

		var givenAnswer string
		fmt.Printf("Question %v: %v\n", i, item.question)

		fmt.Scanf("%v\n", &givenAnswer)

		if item.answer == givenAnswer {
			correctAnswers++
		}
	}

	return correctAnswers
}

func parseFlags() *string {
	csvFlag := flag.String("csv", "problems.csv", "Provide a custom quiz")

	flag.Parse()
	return csvFlag
}
