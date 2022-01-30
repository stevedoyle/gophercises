package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Question struct {
	question string
	answer   string
}

func main() {
	records, err := readData("questions.csv")

	if err != nil {
		log.Fatal(err)
	}

	incorrect := 0
	correct := 0

	for _, record := range records {
		q := Question{
			question: record[0],
			answer:   record[1],
		}

		fmt.Printf("%s: ", q.question)

		var actual string
		fmt.Scan(&actual)
		if q.answer == actual {
			correct += 1
		} else {
			incorrect += 1
		}
	}

	fmt.Printf("Correct: %d, Incorrect: %d\n", correct, incorrect)
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
