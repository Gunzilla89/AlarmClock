package main

import (
	"fmt"
	"os"

	"github.com/Gunzilla89/AlarmClock/alarmclock"
	"github.com/Gunzilla89/AlarmClock/input"
)

func main() {

	args := os.Args[1:]

	if err := input.CheckInput(args); err != nil {
		fmt.Printf("%s", err)
	} else {
		// set alarm
		alarmclock.SetAlarm(args[0])

	}
}
