package bbone

/*
 * Build this for execution on the Beagle board:
 *
 *    GOARM=7 GOARCH=arm GOOS=linux go install ...
 */
import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/beaglebone"
)

// This code gets copied to the beagle board
func NewBbone() *gobot.Robot {

	ba := beaglebone.NewAdaptor()
	led := gpio.NewLedDriver(ba, "P9_12")
	work := func() {

		// Print out we are alive and toggle led every 5 seconds
		gobot.Every(5*time.Second, func() {
			led.Toggle()
			fmt.Println("bbone is alive x 2!")
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{ba},
		[]gobot.Device{led},
		work,
	)
	return robot
}
