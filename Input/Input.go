package input

import "errors"

func CheckInput(input []string) error {
	if len(input) == 0 {
		err := errors.New("not enough arguments\nuse the following format: ./alarmclock HH:MM")
		return err
	} else if len(input) > 2 {
		err := errors.New("too many arguments\nuse the following format: ./alarmclock HH:MM")
		return err
	} else if len(input) == 2 {
		//check for am or pm

	}
	return nil
}
