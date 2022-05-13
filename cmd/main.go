package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 8, 16, 20, 32, 84, 50, 24, 32, 68, 128, 256}

	fmt.Println(isTwoSquare(data))
}

func isTwoSquare(input []int) []bool {
	var results []bool

	for _, i := range input {

		if i == 1 {
			results = append(results, false)
			continue
		}

		for {
			if i == 1 {
				results = append(results, true)
				break
			} else if i%2 != 0 {
				results = append(results, false)
				break
			} else {
				i = i / 2
				continue
			}

		}
	}

	return results
}
