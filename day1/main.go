package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction string

const (
	Right Direction = "R"
	Left  Direction = "L"
)

type Instruction struct {
	Direction Direction
	Steps     int
}

type Dial struct {
	Position  int
	ZeroCount int
}

func (d *Dial) Rotate(direction Direction) {
	switch direction {
	case Right:
		d.Position++
		if d.Position == 100 {
			d.Position = 0
		}
	case Left:
		d.Position--
		if d.Position == -1 {
			d.Position = 99
		}
	}
}

func (d *Dial) RotateN(n int, direction Direction) {
	for i := 0; i < n; i++ {
		d.Rotate(direction)
		if d.Position == 0 {
			d.ZeroCount++
		}
	}
}

func main() {
	lines, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	instructions, err := parseInstructions(lines)
	if err != nil {
		log.Fatal(err)
	}

	dial := Dial{Position: 50, ZeroCount: 0}
	for _, instruction := range instructions {
		dial.RotateN(instruction.Steps, instruction.Direction)
	}

	fmt.Printf("Final Dial Position: %d\n", dial.Position)
	fmt.Printf("Number of times dial was at zero: %d\n", dial.ZeroCount)
}

func readInput() ([]string, error) {
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}

func parseInstruction(line string) (Instruction, error) {
	var dir Direction
	if line[0] == 'R' {
		dir = Right
	} else {
		dir = Left
	}
	step, err := strconv.Atoi(line[1:])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{Direction: dir, Steps: step}, nil
}

func parseInstructions(lines []string) ([]Instruction, error) {
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}
