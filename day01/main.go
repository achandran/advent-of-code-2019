package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Read input as ints to slice
func ReadInts(r io.Reader) ([]int, error) {
	var out []int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return out, fmt.Errorf("could not parse %v as int, %w", line, err)
		}
		out = append(out, n)
	}
	return out, nil
}

// Compute recursive Fuel amount for a mass
func FuelRecursive(mass int) int {
	n := mass
	sum := 0
	for n > 0 {
		n = FuelOnce(n)
		sum += n
	}
	return sum
}

// Compute single Fuel amount for a mass
func FuelOnce(mass int) int {
	candidate := mass/3 - 2
	if candidate < 0 {
		return 0
	}
	return candidate
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("could not read file input.txt")
	}
	defer f.Close()
	masses, err := ReadInts(f)
	if err != nil {
		log.Fatalf("could not read ints %v", err)
	}

	// part 1
	sum1 := 0
	for _, mass := range masses {
		sum1 += FuelOnce(mass)
	}
	fmt.Println(sum1)

	// part 2
	sum2 := 0
	for _, mass := range masses {
		sum2 += FuelRecursive(mass)
	}
	fmt.Println(sum2)
}
