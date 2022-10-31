package simulator

import "fmt"

type Soldier struct {
	isAlive  bool
	position int
}

// Init initializes the Soldier.
func (s *Soldier) Init(position int) {
	s.isAlive = true
	s.position = position
}

// Kill changes state of soldier from alive to dead
func (s *Soldier) Kill() {
	if !s.isAlive {
		panic(fmt.Sprintf("soldier already dead in position: %d", s.position))
	}
	s.isAlive = false
}

// IsAlive indicates whether this soldier is alive or not
func (s *Soldier) IsAlive() bool {
	return s.isAlive
}

// GetPosition gets the position of the Soldier
func (s *Soldier) GetPosition() int {
	return s.position
}
