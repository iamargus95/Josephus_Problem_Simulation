package simulator

import "fmt"

type Result struct {
	LastAlive    int
	KillingOrder []int
}

type CircleOfDeath struct {
	soldiers    []Soldier
	stepPerKill int
}

// Init initializes the CircleOfDeath.
// It takes two argument, first is the number of soldiers in the circle
// Second is steps per kill.
func (c *CircleOfDeath) Init(numberOfSoldiers, stepPerKill int) {

	if numberOfSoldiers < 1 {
		panic(
			fmt.Errorf(
				"number of soldiers: %d should be greater than or equal to 1",
				numberOfSoldiers,
			),
		)
	}

	if stepPerKill >= numberOfSoldiers {
		panic(
			fmt.Errorf(
				"step per kill: %d should be less then number of soldiers: %d",
				numberOfSoldiers,
				stepPerKill,
			),
		)
	}

	c.stepPerKill = stepPerKill

	c.soldiers = make([]Soldier, numberOfSoldiers)

	for index := range c.soldiers {
		c.soldiers[index].Init(index)
	}
}

func (c *CircleOfDeath) findFirstAlive() int {
	firstAlive := -1

	for index, soldier := range c.soldiers {
		if soldier.IsAlive() {
			firstAlive = index
			break
		}
	}

	return firstAlive
}

func (c *CircleOfDeath) getNextIndex(index int) int {

	nextIndex := index + 1

	if nextIndex >= len(c.soldiers) {
		nextIndex = 0
	}

	return nextIndex
}

func (c *CircleOfDeath) findNextAlive(index int) int {
	nextAliveIndex := index

	stepRemaining := c.stepPerKill

	for i := c.getNextIndex(index); i != index; i = c.getNextIndex(i) {
		if c.soldiers[i].IsAlive() {
			nextAliveIndex = i
			stepRemaining -= 1
		}

		if stepRemaining == 0 {
			break
		}
	}

	return nextAliveIndex
}

// Execute executes the simulation
// It returns Result struct instance.
func (c *CircleOfDeath) Execute() Result {

	i := c.findFirstAlive()

	result := Result{}
	result.KillingOrder = make([]int, 0)

	if i == -1 {
		panic("nobody is alive")
	}

	for j := c.findNextAlive(i); i != j; {
		c.soldiers[j].Kill()

		result.KillingOrder = append(result.KillingOrder, j)

		i = c.findNextAlive(i)
		j = c.findNextAlive(i)
	}

	result.LastAlive = c.soldiers[i].GetPosition()

	return result
}
