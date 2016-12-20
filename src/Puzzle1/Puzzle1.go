package Puzzle1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Step ... Counting directed distance
type Step struct {
	Direction string
	N         int
}

// Run ...
func Run() {
	solve("./Puzzle1/Puzzle1_Test1.txt")
	solve("./Puzzle1/Puzzle1_Test2.txt")
	solve("./Puzzle1/Puzzle1_Test3.txt")

	solve("./Puzzle1/Puzzle1_Input.txt")
}

func solve(fileName string) {
	var HQ [4]int
	var step Step

	directions := readProblem(fileName)
	compass := 0

	for i := 0; i < len(directions); i++ {
		step = getInstruction(directions[i])

		if step.Direction == "L" {
			compass = turnLeft(compass)
		}

		if step.Direction == "R" {
			compass = turnRight(compass)
		}

		HQ[compass] += step.N
	}

	fmt.Println(HQ[0] - HQ[2] + HQ[1] - HQ[3])
	fmt.Println("Done...")
}

func readProblem(fileName string) (result []string) {
	dat, err := ioutil.ReadFile(fileName)
	check(err)

	result = strings.Split(string(dat), ",")
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInstruction(x string) (s Step) {
	pivot := strings.TrimPrefix(x, " ")

	d := pivot[0:1]
	n, err := strconv.Atoi(strings.TrimPrefix(pivot, d))
	check(err)
	return Step{Direction: d, N: n}
}

func turnLeft(compass int) (newCompass int) {
	if compass == 0 {
		newCompass = 3
	} else {
		newCompass = compass - 1
	}
	return
}

func turnRight(compass int) (newCompass int) {
	if compass == 3 {
		newCompass = 0
	} else {
		newCompass = compass + 1
	}
	return
}
