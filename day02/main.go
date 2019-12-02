package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const ADD = 1
const MUL = 2
const END = 99

func loadProgram(r io.Reader) ([]int, error) {
	reader := csv.NewReader(r)
	input, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not parse as csv")
	}

	var program []int
	for _, s := range input[0] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("could not parse %v as int", s)
		}
		program = append(program, n)
	}
	return program, nil
}

func run(program []int) (int, error) {
	liveProgram := make([]int, len(program))
	copy(liveProgram, program)
	i := 0
	for {
		opcode, idxA, idxB, dst := liveProgram[i], liveProgram[i+1], liveProgram[i+2], liveProgram[i+3]
		switch opcode {
		case ADD:
			liveProgram[dst] = liveProgram[idxA] + liveProgram[idxB]
		case MUL:
			liveProgram[dst] = liveProgram[idxA] * liveProgram[idxB]
		case END:
			return liveProgram[0], nil
		default:
			return 0, fmt.Errorf("unknown opcode: %v", opcode)
		}
		i = (i + 4) % len(liveProgram)
	}
}

func nounAndVerb(liveProgram []int, target int) (int, int, error) {
	noun, verb := 0, 0

	for {
		for {
			liveProgram[1] = noun
			liveProgram[2] = verb
			answer, err := run(liveProgram)
			if err != nil {
				fmt.Println(err)
				return 0, 0, fmt.Errorf("could not run program with noun = %d, verb = %d", noun, verb)
			}
			switch {
			case answer == target:
				return noun, verb, nil
			case answer < target:
				noun++
				verb++
			case answer > target:
				verb--
			}
		}
	}
}

func main() {
	fd, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fd.Close()

	program, err := loadProgram(fd)
	if err != nil {
		log.Fatalln(err)
	}
	// Override program inputs to reset to desired state
	program[1] = 12
	program[2] = 2

	output, err := run(program)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(output)

	// Part 2
	target := 19690720
	noun, verb, err := nounAndVerb(program, target)
	if err != nil {
		log.Fatalln(err)
	}
	combined := 100*noun + verb
	fmt.Println(combined)
}
