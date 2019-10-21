package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main(){
	csvFilename := flag.String("csv","problems.csv","a csv file in the format of 'Questions, Answers'.")
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
	//fmt.Println(problem)
	correct := 0
	for i, p := range problem{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n",&answer)

		if answer == p.a {
			correct++
		} else {
			fmt.Println("Wrong answer!")
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
