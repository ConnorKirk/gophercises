package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type parsedQuestionAnswer struct {
	question, answer string
}

func main() {

	csvFlag := flag.String("csv", "problems.csv", "Provide a custom quiz")
	timerFlag := flag.Int("timer", 60, "Time limit for the quiz, in seconds")
	flag.Parse()

	questions := getQuestions(csvFlag)
	timer := time.NewTimer(time.Duration(*timerFlag) * time.Second)
	correctAnswers := askQuiz(questions, *timer)
	fmt.Printf("You've reached the end of the quiz!\nYou scored %v out of %v\n", correctAnswers, len(questions))
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

func askQuiz(questions []parsedQuestionAnswer, timer time.Timer) int {
	correctAnswers := 0
problemLoop:
	for i, item := range questions {

		fmt.Printf("Question %v: %v\n", i, item.question)
		answerChan := make(chan string)
		go func() {
			var givenAnswer string
			fmt.Scanf("%v\n", &givenAnswer)
			answerChan <- givenAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime's up!\n")
			break problemLoop

		case answer := <-answerChan:
			if item.answer == answer {
				correctAnswers++
			}

		}

	}
	return correctAnswers
}
