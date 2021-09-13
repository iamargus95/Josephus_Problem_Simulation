package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	flag.Parse()
	input := flag.Arg(0)
	noOfSoldiers, err := validateInput(input)
	if err != nil {
		fmt.Println("Enter valid Integer value.")
	}
	result := josephusSimulation(noOfSoldiers)
	fmt.Printf("\nThe surviving soldier is standing at position no : %d\n\n", result)
}

func validateInput(input string) (int, error) {

	val, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return val, err
}

func josephusSimulation(noOfSoldiers int) int {

	val := math.Log2(float64(noOfSoldiers))
	intoInt := int(val) / 1
	rem := noOfSoldiers - int(math.Pow(2, float64(intoInt)))
	result := (2 * rem) + 1

	return result
}
