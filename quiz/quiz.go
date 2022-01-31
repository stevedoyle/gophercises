package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Question struct {
	question string
	answer   string
}

func main() {
	// Flag processing
	timeoutFlag := flag.Int("time", 30, "Quiz time limit")
	flag.Parse()

	correct := 0
	incorrect := 0

	fmt.Printf("Press Enter to start the Quiz ...")
	var tmp string
	fmt.Scanln(&tmp)
	countdownTimer := time.NewTimer(
		time.Duration(*timeoutFlag) * time.Second)

	quizCh := make(chan int)
	go do_quiz(&correct, &incorrect, quizCh)

	select {
	case <-countdownTimer.C:
		fmt.Println("Quiz ountdown timer expired.")
	case <-quizCh:
		fmt.Println("Quiz completed!")
	}

	fmt.Printf("Correct: %d, Incorrect: %d\n", correct, incorrect)
}

func do_quiz(correct *int, incorrect *int, ch chan<- int) {
	records, err := readData("questions.csv")

	if err != nil {
		log.Fatal(err)
	}

	*incorrect = 0
	*correct = 0

	for _, record := range records {
		q := Question{
			question: record[0],
			answer:   record[1],
		}

		fmt.Printf("%s: ", q.question)

		var actual string
		fmt.Scan(&actual)
		if q.answer == actual {
			*correct += 1
		} else {
			*incorrect += 1
		}
	}
	ch <- 1
}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
