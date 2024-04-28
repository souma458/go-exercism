package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery = c.battery - c.batteryDrain
		c.distance = c.distance + c.speed
	}
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	for c.battery > 0 && c.battery >= c.batteryDrain {
		c.Drive()
	}
	return c.distance >= trackDistance
}
