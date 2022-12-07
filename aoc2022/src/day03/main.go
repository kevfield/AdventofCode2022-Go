package main

import "fmt"

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	p1result := part1(puzzleInput)
	p2result := part2(puzzleInput)

	fmt.Println("part 1 = ", p1result)
	fmt.Println("part 2 =", p2result)
}

func alphabetCalc(finalSlice []byte) int {
	// create alphabet/value slice and convert to byte
	var alphaSlice []byte
	for ch := 'a'; ch <= 'z'; ch++ {
		alphaSlice = append(alphaSlice, byte(ch))
	}
	for chUp := 'A'; chUp <= 'Z'; chUp++ {
		alphaSlice = append(alphaSlice, byte(chUp))
	}

	var totalPriority int
	for l := 0; l < len(alphaSlice); l++ {
		for m := 0; m < len(finalSlice); m++ {
			if finalSlice[m] == alphaSlice[l] {
				totalPriority = totalPriority + int(l+1)
			}
		}
	}
	return totalPriority
}

func part1(p1Input []string) int {
	// loop through and compare the first half and second half and grab the match
	var p1finalSlice []byte
	for i := 0; i < len(p1Input); i++ {
		lengthSlice := len(p1Input[i])
		halfSlice := lengthSlice / 2
	innerloop:
		for j := 0; j < halfSlice; j++ {
			for k := halfSlice; k < lengthSlice; k++ {
				if p1Input[i][j] == p1Input[i][k] {
					p1finalSlice = append(p1finalSlice, p1Input[i][j])
					break innerloop
				}
			}

		}
	}
	return alphabetCalc(p1finalSlice)
}

func part2(p2Input []string) int {
	//var elfPriority int
	var p2finalSlice []byte

	for i := 0; i < len(p2Input); {
		elf1 := p2Input[i]
		elf2 := p2Input[i+1]
		elf3 := p2Input[i+2]
	outerloop:
		for j := 0; j < len(elf1); j++ {
			for k := 0; k < len(elf2); k++ {
				if elf1[j] == elf2[k] {
					for l := 0; l < len(elf3); l++ {
						if elf1[j] == elf3[l] {
							p2finalSlice = append(p2finalSlice, elf2[k])
							break outerloop
						}
					}

				}
			}
		}

		i += 3
	}

	return alphabetCalc(p2finalSlice)

}
