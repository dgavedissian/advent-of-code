package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func hasTwoThree(str string) (bool, bool) {
	letterCount := make(map[rune]int)
	for _, char := range str {
		letterCount[char]++
	}
	hasTwoLetters := false
	hasThreeLetters := false
	for _, v := range letterCount {
		if v == 2 {
			hasTwoLetters = true
		}
		if v == 3 {
			hasThreeLetters = true
		}
	}
	return hasTwoLetters, hasThreeLetters
}

func part1() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	numTwoLetters := 0
	numThreeLetters := 0
    for scanner.Scan() {
		hasTwo, hasThree := hasTwoThree(scanner.Text())
		if hasTwo {
			numTwoLetters++
		}
		if hasThree {
			numThreeLetters++
		}
	}
	
	fmt.Println(numTwoLetters * numThreeLetters)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func part2() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	numTwoLetters := 0
	numThreeLetters := 0
    for scanner.Scan() {
		hasTwo, hasThree := hasTwoThree(scanner.Text())
		if hasTwo {
			numTwoLetters++
		}
		if hasThree {
			numThreeLetters++
		}
	}
	
	fmt.Println(numTwoLetters * numThreeLetters)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main() {
	part1()
	part2()
}