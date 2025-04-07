package mouse

import (
	"fmt"
	"time"
)

type MouseMover struct {
	size  int
	delay time.Duration
}

func (m *MouseMover) PressRandomKey() error {
	return nil
}

func (m *MouseMover) Move(direction string, distance int) error {
	return nil
}

func (m *MouseMover) MoveInSquare() error {
	return nil
}

func (m *MouseMover) Start() {
	fmt.Println("Moving the mouse in a square... (press Ctrl + C to stop)")

	for {
		if err := m.PressRandomKey(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if err := m.MoveInSquare(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
