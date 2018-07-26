package master

import (
	"fmt"
	"os"

	"html"
	"net/http"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/mqtt"
)

func NewDataLogger(port string) *gobot.Robot {
	mqttAdaptor := mqtt.NewAdaptor("tcp://192.168.1.4:1833", "rockpi")
	mqttAdaptor.SetAutoReconnect(true)
	mqttSensors := mqtt.NewDriver(mqttAdaptor, "sensors/rockpi")
	firmataAdaptor := firmata.NewAdaptor(port)

	led := gpio.NewLedDriver(firmataAdaptor, "13")
	button := gpio.NewButtonDriver(firmataAdaptor, "12")
	buzzer := gpio.NewBuzzerDriver(firmataAdaptor, "10")

	pot := aio.NewAnalogSensorDriver(firmataAdaptor, "0")
	snd := aio.NewAnalogSensorDriver(firmataAdaptor, "2")

	last_pot_val := uint16(0)
	last_snd_val := uint16(0)

	work := func() {

		snd.On(aio.Data, func(data interface{}) {
			//fmt.Println("snd:10:", data)
			last_snd_val = uint16(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 256),
			)
		})

		pot.On(aio.Data, func(data interface{}) {
			last_pot_val = uint16(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 256),
			)
		})

		/*
		 * Can't seem to get button to work? :(  configured for input?
		 */
		button.On(gpio.ButtonPush, func(data interface{}) {
		})

		gobot.Every(60*time.Second, func() {
			potstr := "pot:2:" + string(last_pot_val)
			fmt.Println(potstr)
			bytestr := []byte(potstr)
			mqttSensors.Publish(bytestr)

			sndstr := "snd:0:" + string(last_snd_val)
			fmt.Println(sndstr)
			bytestr = []byte(sndstr)
			mqttAdaptor.Publish("sensors/rockpi/snd", bytestr)
			// mqttSensors.Publish("sensors/rockpi/snd", sndstr)
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Toggle()
			buzzer.Tone(gpio.C4, gpio.Quarter)
		})
	}

	datalogger := gobot.NewRobot("datalogger-"+port,
		[]gobot.Connection{firmataAdaptor, mqttAdaptor},
		[]gobot.Device{Led, button, buzzer, pot, snd},
		work,
	)
	return datalogger
}

func NewAPIServer(master *gobot.Master) *api.API {
	server := api.NewAPI(master)
	server.Port = "3000"

	server.AddHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
	})

	master.AddCommand("otto",
		func(params map[string]interface{}) interface{} {
			return "This command is for otto"
		})

	hello := master.AddRobot(gobot.NewRobot("hello"))
	hello.AddCommand("Hiii", func(params map[string]interface{}) interface{} {
		return fmt.Sprintf("This command is attached to robot: %v", hello.Name)
	})
	return server
}

func main() {

	// "/dev/cu.usbmodem1411111",
	serial_ports := []string{
		"/dev/tty.usbmodem145121",
	}

	// Create a master robot controller
	master := gobot.NewMaster()
	for _, port := range serial_ports {
		master.AddRobot(NewDataLogger(port))
	}

	// Create an API server
	server := NewAPIServer(master)
	server.Debug()
	server.Start()

	// Start the Robots
	master.Start()

	os.Exit(0)
}
