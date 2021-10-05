package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Gunzilla89/AlarmClock/Input/"
)

func main() {

	args := os.Args[1:]

	if err := Input.CheckInput(args); err != nil {
		fmt.Printf("%s", err)
	} else {
		// set alarm
		setAlarm(args[0])

	}
}

func setAlarm(input string) {
	fmt.Println("Setting alarm...")
	currentTime := time.Now()
	fmt.Printf("The current time is %s\n", currentTime.Format("03:04 PM"))
	fmt.Printf("Alarm is set for, %s\n", input)
}
