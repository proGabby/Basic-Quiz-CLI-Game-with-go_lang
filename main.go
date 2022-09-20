package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	flag.Parse()
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
	for i, prob := range problems {
		fmt.Printf("problem #%d: %s = \n", i+1, prob.question)
		var answer string
		// //using scanf to ensure spaces in input are ignore
		fmt.Scanf("%s\n", &answer)
		if answer == prob.answer {
			correct++
			fmt.Println("correct!")
		} else {
			fmt.Println("incorrect!")
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
