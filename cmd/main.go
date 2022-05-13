package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := []int{}

	// fmt.Println(isArrayTwoSquare(data))
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Is your number is Square by 2 (\"exit\" for stop)")
	fmt.Println("---------------------")
	fmt.Print("integer : ")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("%s not integer\n", text)
			fmt.Print("integer : ")
			continue
		}

		data = append(data, value)
		fmt.Println(isTwoSquare(value))
		fmt.Print("integer : ")
	}

	fmt.Println(isArrayTwoSquare(data))
}

func isArrayTwoSquare(input []int) []bool {
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

func isTwoSquare(input int) bool {

	if input == 1 {
		return false
	}

	for {
		if input == 1 {
			return true
		} else if input%2 != 0 {
			return false
		} else {
			input = input / 2
			continue
		}

	}
}
