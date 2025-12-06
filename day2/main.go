package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func isInvalid(id int) bool {
	idString := strconv.Itoa(id)
	if len(idString)%2 != 0 {
		return false
	}

	firstHalf := idString[:len(idString)/2]
	secondHalf := idString[len(idString)/2:]

	return firstHalf == secondHalf
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, ",")
	invalidSum := 0
	for _, part := range parts {
		startAndEnd := strings.Split(part, "-")
		if len(startAndEnd) != 2 {
			log.Fatal("Wrong length startAndEnd")
		}
		start, err := strconv.Atoi(startAndEnd[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(startAndEnd[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := start; i <= end; i++ {
			if isInvalid(i) {
				invalidSum += i
			}
		}
	}
	log.Println("Invalid sum:", invalidSum)
}
