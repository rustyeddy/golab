package duino

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"

	"time"
)

func NewDuino(port string) *gobot.Robot {

	// Firmata adapter talk over serial port
	firma := firmata.NewAdaptor(port)

	// The onboad led
	led := gpio.NewLedDriver(firma, "13")

	// Our callback function every 2 seconds to toggle the led.
	work := func() {
		gobot.Every(2 * time.Second, func() {
			led.Toggle()
		})
	}

	// Our very own robot
	robot := gobot.NewRobot("duino+" + port,
		[]gobot.Connection{firma},
		[]gobot.Device{led},
		work,
	)
	return robot
}
