package raspi

/*
 * Build this for execution on the Raspberry Pi:
 *
 *    GOARM=6 GOARCH=arm GOOS=linux go install ...
 */
import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func NewRaspi() *gobot.Robot {
	r := raspi.NewAdaptor()

	// Grab the builtin LED for fun
	led := gpio.NewLedDriver(r, "13")
	work := func() {
		gobot.Every(1 * time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)
	return robot
}
