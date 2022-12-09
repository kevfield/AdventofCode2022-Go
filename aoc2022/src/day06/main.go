package main

import "fmt"

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	fmt.Println("Part 1 =", getMarker(puzzleInput))
	fmt.Println("Part 2 =", getMessage(puzzleInput))
}

func getMarker(gmInput []string) int {
	markPos := 0
	for i := 3; i < len(gmInput[0]); i++ {

		if gmInput[0][i] != gmInput[0][i-3] && gmInput[0][i] != gmInput[0][i-2] && gmInput[0][i] != gmInput[0][i-1] {
			if gmInput[0][i-1] != gmInput[0][i-2] && gmInput[0][i-1] != gmInput[0][i-3] {
				if gmInput[0][i-2] != gmInput[0][i-3] {
					markPos = i + 1
					break
				}
			}

		}
	}

	return markPos
}

// markPos = i + 1
// break
func getMessage(gmInput []string) int {
	markPos := 0
	for i := 14; i < len(gmInput[0]); i++ {

		if gmInput[0][i] != gmInput[0][i-14] && gmInput[0][i] != gmInput[0][i-13] && gmInput[0][i] != gmInput[0][i-12] && gmInput[0][i] != gmInput[0][i-11] && gmInput[0][i] != gmInput[0][i-10] && gmInput[0][i] != gmInput[0][i-9] && gmInput[0][i] != gmInput[0][i-8] && gmInput[0][i] != gmInput[0][i-7] && gmInput[0][i] != gmInput[0][i-6] && gmInput[0][i] != gmInput[0][i-5] && gmInput[0][i] != gmInput[0][i-4] && gmInput[0][i] != gmInput[0][i-3] && gmInput[0][i] != gmInput[0][i-2] && gmInput[0][i] != gmInput[0][i-1] {
			if gmInput[0][i-13] != gmInput[0][i-12] && gmInput[0][i-13] != gmInput[0][i-11] && gmInput[0][i-13] != gmInput[0][i-10] && gmInput[0][i-13] != gmInput[0][i-9] && gmInput[0][i-13] != gmInput[0][i-8] && gmInput[0][i-13] != gmInput[0][i-7] && gmInput[0][i-13] != gmInput[0][i-6] && gmInput[0][i-13] != gmInput[0][i-5] && gmInput[0][i-13] != gmInput[0][i-4] && gmInput[0][i-13] != gmInput[0][i-3] && gmInput[0][i-13] != gmInput[0][i-2] && gmInput[0][i-13] != gmInput[0][i-1] {
				if gmInput[0][i-12] != gmInput[0][i-11] && gmInput[0][i-12] != gmInput[0][i-10] && gmInput[0][i-12] != gmInput[0][i-9] && gmInput[0][i-12] != gmInput[0][i-8] && gmInput[0][i-12] != gmInput[0][i-7] && gmInput[0][i-12] != gmInput[0][i-6] && gmInput[0][i-12] != gmInput[0][i-5] && gmInput[0][i-12] != gmInput[0][i-4] && gmInput[0][i-12] != gmInput[0][i-3] && gmInput[0][i-12] != gmInput[0][i-2] && gmInput[0][i-12] != gmInput[0][i-1] {
					if gmInput[0][i-11] != gmInput[0][i-10] && gmInput[0][i-11] != gmInput[0][i-9] && gmInput[0][i-11] != gmInput[0][i-8] && gmInput[0][i-11] != gmInput[0][i-7] && gmInput[0][i-11] != gmInput[0][i-6] && gmInput[0][i-11] != gmInput[0][i-5] && gmInput[0][i-11] != gmInput[0][i-4] && gmInput[0][i-11] != gmInput[0][i-3] && gmInput[0][i-11] != gmInput[0][i-2] && gmInput[0][i-11] != gmInput[0][i-1] {
						if gmInput[0][i-10] != gmInput[0][i-9] && gmInput[0][i-10] != gmInput[0][i-8] && gmInput[0][i-10] != gmInput[0][i-7] && gmInput[0][i-10] != gmInput[0][i-6] && gmInput[0][i-10] != gmInput[0][i-5] && gmInput[0][i-10] != gmInput[0][i-4] && gmInput[0][i-10] != gmInput[0][i-3] && gmInput[0][i-10] != gmInput[0][i-2] && gmInput[0][i-10] != gmInput[0][i-1] {
							if gmInput[0][i-9] != gmInput[0][i-8] && gmInput[0][i-9] != gmInput[0][i-7] && gmInput[0][i-9] != gmInput[0][i-6] && gmInput[0][i-9] != gmInput[0][i-5] && gmInput[0][i-9] != gmInput[0][i-4] && gmInput[0][i-9] != gmInput[0][i-3] && gmInput[0][i-9] != gmInput[0][i-2] && gmInput[0][i-9] != gmInput[0][i-1] {
								if gmInput[0][i-8] != gmInput[0][i-7] && gmInput[0][i-8] != gmInput[0][i-6] && gmInput[0][i-8] != gmInput[0][i-5] && gmInput[0][i-8] != gmInput[0][i-4] && gmInput[0][i-8] != gmInput[0][i-3] && gmInput[0][i-8] != gmInput[0][i-2] && gmInput[0][i-8] != gmInput[0][i-1] {
									if gmInput[0][i-7] != gmInput[0][i-6] && gmInput[0][i-7] != gmInput[0][i-5] && gmInput[0][i-7] != gmInput[0][i-4] && gmInput[0][i-7] != gmInput[0][i-3] && gmInput[0][i-7] != gmInput[0][i-2] && gmInput[0][i-7] != gmInput[0][i-1] {
										if gmInput[0][i-6] != gmInput[0][i-5] && gmInput[0][i-6] != gmInput[0][i-4] && gmInput[0][i-6] != gmInput[0][i-3] && gmInput[0][i-6] != gmInput[0][i-2] && gmInput[0][i-6] != gmInput[0][i-1] {
											if gmInput[0][i-5] != gmInput[0][i-4] && gmInput[0][i-5] != gmInput[0][i-3] && gmInput[0][i-5] != gmInput[0][i-2] && gmInput[0][i-5] != gmInput[0][i-1] {
												if gmInput[0][i-4] != gmInput[0][i-3] && gmInput[0][i-4] != gmInput[0][i-2] && gmInput[0][i-4] != gmInput[0][i-1] {
													if gmInput[0][i-3] != gmInput[0][i-2] && gmInput[0][i-3] != gmInput[0][i-1] {
														if gmInput[0][i-2] != gmInput[0][i-1] {
															markPos = i + 1
															break
														}

													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return markPos
}
