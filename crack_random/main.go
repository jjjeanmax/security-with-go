package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func nxt(seed int) int {
	seed2 := (seed*25214903917 + 11) % int((math.Pow(2, 48)))
	return seed2 >> 16
}

func entrees() int {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)

	}

	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")

	}
	return guess
}

func gestSeed(n1, n2 int) int {

	for i := 0; i < 65536; i++ {
		seed := (n1 * 65536) + i
		if nxt(seed) == n2 {
			fmt.Println(">>> Found seed", seed)
			return seed
		}
	}
	return 0
}

// a>>b = a/(2^b)
// a<<b = a*(2^b)
func main() {
	fmt.Println("n1 and n2")
	n1 := entrees()
	n2 := entrees()

	seed := gestSeed(n1, n2)
	for i := 0; i < 5; i++ {
		x := seed >> 16
		if x >= int((math.Pow(2, 31))) {
			x -= int((math.Pow(2, 32)))
		}
		fmt.Println(x)
		seed = (seed*25214903917 + 11) % int((math.Pow(2, 48)))
	}

}
