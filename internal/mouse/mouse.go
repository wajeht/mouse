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

type Controller interface {
	GetMousePos() (int, int)
	MoveMouse(x, y int)
	PressKey(key string)
	Sleep(ms int)
}

type RobotController struct{}

func (r *RobotController) GetMousePos() (int, int) {
	return robotgo.Location()
}

func (r *RobotController) MoveMouse(x, y int) {
	robotgo.MoveSmooth(x, y, 1.0, 10.0)
}

func (r *RobotController) PressKey(key string) {
	robotgo.KeyToggle(key, "down")
	robotgo.MilliSleep(50)
	robotgo.KeyToggle(key, "up")
}

func (r *RobotController) Sleep(ms int) {
	robotgo.MilliSleep(ms)
}

type Mover struct {
	cfg  Config
	rng  *rand.Rand
	ctrl Controller
}

func NewMover(cfg Config, ctrl Controller, rng *rand.Rand) *Mover {
	return &Mover{
		cfg:  cfg,
		rng:  rng,
		ctrl: ctrl,
	}
}

func (m *Mover) move(dir string, dist int) error {
	x, y := m.ctrl.GetMousePos()
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
	m.ctrl.MoveMouse(x, y)
	return nil
}

func (m *Mover) pressRandomKey() {
	keys := []string{"escape", "tab", "shift", "alt", "capslock", "up", "down", "left", "right"}
	key := keys[m.rng.Intn(len(keys))]

	if m.cfg.DryRun {
		fmt.Printf("Would press key: %s\n", key)
		return
	}

	m.ctrl.PressKey(key)
}

func (m *Mover) moveInSquare() error {
	for _, dir := range []string{"right", "down", "left", "up"} {
		if err := m.move(dir, m.cfg.Size); err != nil {
			return err
		}
		m.ctrl.Sleep(m.cfg.Delay)
	}
	return nil
}

func (m *Mover) Run(ctx context.Context) error {
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
