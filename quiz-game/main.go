package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func waitTime(t int, c chan int) {
	for i := 0; i < t; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	close(c)
}

func takeQuiz(records [][]string, c chan int) {
	corrects := 0
	for i, el := range records {
		// Check if record has length 2 and no invalid values
		if len(el) == 2 && el[0] != "" && el[1] != "" {
			// Show question and ask for answer
			fmt.Printf("Question %v: %v\n", i+1, el[0])
			fmt.Print("Type your answer: ")
			in := bufio.NewReader(os.Stdin)
			input, err := in.ReadString('\n')
            if err != nil {
                log.Fatal(err)
            }

			// Cleanup
			input = strings.TrimSpace(input)
			input = strings.ToLower(input)
			actualAnswer := strings.TrimSpace(el[1])
			actualAnswer = strings.ToLower(actualAnswer)

			// Compare user answer to correct answer
			if input == actualAnswer {
				corrects += 1
			}
			c <- corrects
		} else {
			fmt.Printf("Question %v was created with wrong values, counts as incorrect\n", i+1)
		}
	}
	close(c)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Get flag from command line
	fileFlag := flag.String("file", "problems.csv", "The csv file to take the quiz from")
	timeFlag := flag.Int("time", 30, "The time the quiz will take")
	shuffleFlag := flag.Bool("shuffle", false, "If set the order of the questions is shuffled")
	flag.Parse()
	path := *fileFlag
	maxTime := *timeFlag
	shuffle := *shuffleFlag

	// Read file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// Parse file
	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// Shuffle if requested
	if shuffle {
		rand.Shuffle(len(records), func(i int, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	// Show quiz
	corrects := 0
	quizChannel := make(chan int)
	timerChannel := make(chan int)
	fmt.Println("Press enter to start the quiz")
	bufio.NewScanner(os.Stdin).Scan()
	fmt.Printf("You have %v seconds\n", maxTime)
	go takeQuiz(records, quizChannel)
	go waitTime(maxTime, timerChannel)

out:
	for {
		select {
		case co, open := <-quizChannel:
			if !open {
				break out
			}
			corrects = co
		case _, open := <-timerChannel:
			if !open {
				fmt.Println("\nYou ran out of time!")
				break out
			}
		}
	}

	// Show number of correct answers and total questions
	total := len(records)
	incorrects := total - corrects
	fmt.Printf("Correct answers: %v / Incorrect answers: %v / Total questions: %v", corrects, incorrects, total)
}
