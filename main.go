package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	csvFilename := flag.String("csv","problems.csv","a csv file in the format of 'Questions, Answers'.")
	timeLimit := flag.Int("limit",30,"The time limit for the quiz in seconds.")
	flag.Parse()

	file, err := os.Open(*csvFilename)		//The csvFilename should be a pointer to a string, Predefined functionality of the file.
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV File: %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
		if err != nil {
			exit("Failed to parse the csv file.")
		}
	problem := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problem{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh<-answer
		}()

		select {
			case <-timer.C:			//<- used to indicate channel holds the code for specified time limit until it gets the response.
				fmt.Printf("\nYou scored %d out of %d", correct, len(problem))
				return

				case answer:= <- answerCh:
				if answer == p.a {
					correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problem))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range  lines {
		ret[i]= Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}


type Problem struct {
	q, a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
