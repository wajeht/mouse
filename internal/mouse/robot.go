package mouse

import "github.com/go-vgo/robotgo"

type Roboto struct{}

func (r *Roboto) GetMousePos() (int, int) {
	return robotgo.Location()
}

func (r *Roboto) MoveMouse(x, y int) {
	robotgo.MoveSmooth(x, y, 1.0, 10.0)
}

func (r *Roboto) PressKey(key string) {
	robotgo.KeyToggle(key, "down")
	robotgo.MilliSleep(50)
	robotgo.KeyToggle(key, "up")
}

func (r *Roboto) Sleep(ms int) {
	robotgo.MilliSleep(ms)
}
