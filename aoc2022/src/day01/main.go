package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	var highestCalories1 int
	var currentCalories int
	var calorieSlice []int

	// loop through input and calculate the calories

	for i := 0; i < len(puzzleInput); i++ {
		if puzzleInput[i] != "" {
			intConverted, _ := strconv.Atoi(puzzleInput[i])
			currentCalories = currentCalories + intConverted
		} else {
			if currentCalories > highestCalories1 {
				calorieSlice = append(calorieSlice, currentCalories)
				highestCalories1 = currentCalories
				currentCalories = 0
			} else {
				calorieSlice = append(calorieSlice, currentCalories)
				currentCalories = 0
			}
		}
	}

	sort.Slice(calorieSlice, func(i, j int) bool {
		return calorieSlice[i] < calorieSlice[j]
	})

	calsliceLength := len(calorieSlice)

	fmt.Println("part 1 answer =", highestCalories1)
	fmt.Println("part 2 answer =", calorieSlice[calsliceLength-1]+calorieSlice[calsliceLength-2]+calorieSlice[calsliceLength-3])

}
