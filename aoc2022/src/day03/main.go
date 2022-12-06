package main

import "fmt"

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)
	var finalSlice []byte

	// create alphabet/value slice and convert to byte
	var alphaSlice []byte
	for ch := 'a'; ch <= 'z'; ch++ {
		alphaSlice = append(alphaSlice, byte(ch))
	}
	for chUp := 'A'; chUp <= 'Z'; chUp++ {
		alphaSlice = append(alphaSlice, byte(chUp))
	}

	// loop through and compare the first half and second half and grab the match
	for i := 0; i < len(puzzleInput); i++ {
		lengthSlice := len(puzzleInput[i])
		halfSlice := lengthSlice / 2
	innerloop:
		for j := 0; j < halfSlice; j++ {
			for k := halfSlice; k < lengthSlice; k++ {
				if puzzleInput[i][j] == puzzleInput[i][k] {
					finalSlice = append(finalSlice, puzzleInput[i][j])
					break innerloop
				}
			}

		}
	}

	//loop through finalSlice to get values from alphaSlice
	var totalPriority int
	for l := 0; l < len(alphaSlice); l++ {
		for m := 0; m < len(finalSlice); m++ {
			if finalSlice[m] == alphaSlice[l] {
				totalPriority = totalPriority + int(l+1)
			}
		}
	}
	fmt.Println(totalPriority)
}
