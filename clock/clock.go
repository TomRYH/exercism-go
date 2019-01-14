package clock

import "fmt"

// Clock models a 24hr clock
type Clock struct {
	h, m int
}

const (
	maxMins  = 60
	maxHours = 24
)

// Add minutes to the clock
func (c Clock) Add(m int) Clock {
	c.m += m
	return c.FixOverflow()
}

// Subtract minutes from the clock
func (c Clock) Subtract(m int) Clock {
	c.m -= m
	return c.FixOverflow()
}

// String returns the time in "hh:mm" format
func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// FixOverflow stops minutes and hours overflowing their max
func (c Clock) FixOverflow() Clock {
	c.h += c.m / maxMins

	if c.m %= maxMins; c.m < 0 {
		c.h--
		c.m += maxMins
	}

	if c.h %= maxHours; c.h < 0 {
		c.h += maxHours
	}

	return c
}

// New creates a new clock
func New(h, m int) Clock {
	return Clock{h: h, m: m}.FixOverflow()
}
