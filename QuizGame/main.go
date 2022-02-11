package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Questions struct {
	Question string
	Answer   string
}

func readFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Welcome to the quiz game")
	defer file.Close()
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvLines
}

func parseFileToStruct(fileArray [][]string) []Questions {
	var QuestionArray []Questions

	for _, line := range fileArray {
		emp := Questions{
			Question: line[0],
			Answer:   line[1],
		}
		QuestionArray = append(QuestionArray, emp)
	}
	return QuestionArray
}

func askQuestionsAndGetAnswers(QuestionArray []Questions) {
	newtimer := time.NewTimer(30 * time.Second)
	var answers string
	var correctAnswers = 0

	for i := range QuestionArray {
		go func() {
			<-newtimer.C

			// Printed when timer is fired
			fmt.Println("sure bitti")
			fmt.Printf("you made %d correct answers inside %d of questions", correctAnswers, len(QuestionArray))
			os.Exit(0)
		}()

		fmt.Printf("what is value of %s: ", QuestionArray[i].Question)
		fmt.Scanf("%s", &answers)
		if answers == QuestionArray[i].Answer {
			correctAnswers++
		}
	}
	fmt.Printf("you made %d correct answers inside %d of questions", correctAnswers, len(QuestionArray))
}
func main() {

	readedFile := readFile("problems.csv")
	QuestionArray := parseFileToStruct(readedFile)
	askQuestionsAndGetAnswers(QuestionArray)

}
