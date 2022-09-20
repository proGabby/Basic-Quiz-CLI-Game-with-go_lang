package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//defining cli flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for your quiz in seconds")

	//parsing the flags on the cli
	flag.Parse()
	//opein the file
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s\n", *csvFilename))
	}

	// a reader to read the csc file
	r := csv.NewReader(file)

	//reading all records in the file
	lines, err := r.ReadAll()
	if err != nil {
		exit("faild to manipulate the provided cs file")
	}
	//variable to keep track of the correct answers
	correct := 0
	problems := parseLine(lines)

	//declare a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, prob := range problems {
		fmt.Printf("problem #%d: %s = \n", i+1, prob.question)

		answerchanel := make(chan string)

		//create and call an anonymous go-routine to accept user input and send the data to answerchannel
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerchanel <- ans
		}()

		select {
		//check cases where the timer is out
		case <-timer.C:
			fmt.Printf("time your! \n you scrored %d out of %d.\n", correct, len(problems))
			return

			//check case where anserchannel has data
		case ansr := <-answerchanel:
			if ansr == prob.answer {
				correct++
				fmt.Println("correct!")
			} else {
				fmt.Println("incorrect!")
			}
		}

	}
	// fmt.Println(lines)
	fmt.Printf("you scrored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	question string
	answer   string
}

//exit function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// iterate through all the record and parse each of them into a problem object
func parseLine(lines [][]string) []problem {
	//make a problem slice of problem object with size of lenght of the file record
	retn := make([]problem, len(lines))
	for i, line := range lines {
		retn[i] = problem{
			question: line[0],
			//triming the spaces in the answers on the csv file
			answer: strings.TrimSpace(line[1]),
		}
	}
	return retn
}
