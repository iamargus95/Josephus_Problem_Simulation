package main

import (
	"flag"
	"fmt"
	"iamargus95/josephus/simulator"
	"os"
	"strconv"
	"strings"
	"time"
)

// Get N from the command line
// Arguments: <None>
// Return value: []int
func getCSVInput() []int {
	n := flag.String("N", "", "Total number of soldiers in the circle [where N > 0]")

	flag.Parse()
	if *n == "" {
		flag.PrintDefaults()
		return nil
	}

	inputsString := strings.Split(*n, ",")

	var inputsInt []int
	for _, v := range inputsString {
		i, err := strconv.Atoi(v)
		if err != nil {
			flag.PrintDefaults()
			return nil
		}
		inputsInt = append(inputsInt, i)
	}

	return inputsInt
}

func main() {

	circlesOfDeath := getCSVInput()

	for _, circle := range circlesOfDeath {
		newCircle := new(simulator.CircleOfDeath)

		newCircle.Init(circle, 1)

		for _, v := range newCircle.Execute().KillingOrder {
			fmt.Println("\t", v+1, "is dead")
			time.Sleep(50 * time.Millisecond)
		}

		fmt.Fprintf(
			os.Stderr,
			"\n\tJosephus is standing at position no. %d in the circle of death among %d soldiers.\n\n",
			newCircle.Execute().LastAlive+1,
			circle,
		)

		fmt.Printf(
			"\n\tJosephus is standing at position no. %d in the circle of death among %d soldiers.\n\n",
			newCircle.Execute().LastAlive+1,
			circle,
		)
	}

}
