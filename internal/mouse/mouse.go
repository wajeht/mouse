package mouse

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-vgo/robotgo"
)

type Config struct {
	Size   int
	Delay  int
	DryRun bool
}

type Mouse struct {
	cfg Config
}

func New(cfg Config) *Mouse {
	return &Mouse{cfg: cfg}
}

func (m *Mouse) getMousePos() (int, int) {
	return robotgo.Location()
}

func (m *Mouse) moveMouse(x, y int) {
	robotgo.MoveSmooth(x, y, 1.0, 10.0)
}

func (m *Mouse) pressKey(key string) {
	robotgo.KeyToggle(key, "down")
	robotgo.MilliSleep(50)
	robotgo.KeyToggle(key, "up")
}

func (m *Mouse) sleep(ms int) {
	robotgo.MilliSleep(ms)
}

func (m *Mouse) move(dir string, dist int) error {
	x, y := m.getMousePos()
	switch dir {
	case "left":
		x -= dist
	case "right":
		x += dist
	case "up":
		y -= dist
	case "down":
		y += dist
	default:
		return fmt.Errorf("invalid direction: %s", dir)
	}
	m.moveMouse(x, y)
	return nil
}

func (m *Mouse) pressRandomKey() {
	keys := []string{"escape", "tab", "shift", "alt", "capslock", "up", "down", "left", "right"}
	key := keys[rand.Intn(len(keys))]

	if m.cfg.DryRun {
		fmt.Printf("Would press key: %s\n", key)
		return
	}

	m.pressKey(key)
}

func (m *Mouse) moveInSquare() error {
	for _, dir := range []string{"right", "down", "left", "up"} {
		if err := m.move(dir, m.cfg.Size); err != nil {
			return err
		}
		m.sleep(m.cfg.Delay)
	}
	return nil
}

func (m *Mouse) Run(ctx context.Context) error {
	fmt.Println("Running... (Ctrl+C to stop)")

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			m.pressRandomKey()
			if err := m.moveInSquare(); err != nil {
				fmt.Printf("Move error: %v\n", err)
			}
		}
	}
}
