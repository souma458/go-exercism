package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	m += h * 60
	h = m / 60

	m = m % 60
	if m < 0 {
		m += 60
		h--
	}

	h = h % 24
	if h < 0 {
		h += 24
	}

	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	totalMin := c.hour*60 + c.minute + m
	return New(totalMin/60, totalMin%60)
}

func (c Clock) Subtract(m int) Clock {
	totalMin := c.hour*60 + c.minute - m
	return New(totalMin/60, totalMin%60)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
