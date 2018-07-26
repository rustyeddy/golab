package main

/*
   "gobot.io/x/gobot/drivers/gpio"
   "gobot.io/x/gobot/platforms/firmata"
*/

func main() {
	//fmt.Println("Gotto is the real thing baby!")
	//fmt.Printf("\tos: %s - arch: %s\n", runtime.GOOS, runtime.GOARCH)

	/*
	 * We will make some assumptions for now.  For Darwin & Linux
	 * will be MACRO controller.  Most likely NOT a MACRO controller
	 * and maybe a MASTER controller.
	 *
	 * For GOARCH is arm we'll assume it is a MACRO controller,
	 * probably a MICRO controller (gpio), but probably not a master.
	 */

	/*
	 * Just assume (for now) we are darwin with an arduino attached.
	 * We will also manage two raspberry pi's and a beaglebone black,
	 * the later two also have arduino's.
	 */
	// Create a bot for our duino
	/*
		port := "/dev/tty.usbmodem145121"
		d1 := duino.NewDuino(port)

		port = "/dev/tty.usbmodem1451121"
		d2 := duino.NewDuino(port)

		// Create a bot for the raspi and bbone we are managing
		raspi := raspi.NewRaspi()
		bbone := bbone.NewBbone()

		master := gobot.NewMaster()
		addAPI(master)
		addCommands(master)

		mqtt := mqtt.NewAdaptor("tcp://192.168.1.4:1883", "master")
		err := mqtt.Connect()
		if err != nil {
			fmt.Println("Some error happened MQTT`ing!")
		}

		fmt.Println("Publishing to mqtt")
		mqtt.Publish("sensors/foo/bar", []byte("aser"))

		master.AddRobot(d1)
		master.AddRobot(d2)
		master.AddRobot(raspi)
		master.AddRobot(bbone)

		// Set this up just for the heck of it
		gobot.Every(3*time.Second, func() {
			fmt.Println("Publishing more... sensors")
			mqtt.Publish("sensors/foo", []byte("fufufuf"))
		})

		// This is it
		master.Start()
	*/
}

/*
func addAPI(master *gobot.Master) {
	// Create the master of control
	a := api.NewAPI(master)
	a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Heeeelllloooo, %q\n", html.EscapeString(r.URL.Path))
	})
	a.Debug()
	a.Start()
}

func addCommands(master *gobot.Master) {
	master.AddCommand("cmd_gotto_version",
		func(params map[string]interface{}) interface{} {
			return "This command is attached to master"
		})
}
*/
