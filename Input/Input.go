package input

import "errors"

type Input interface {
	CheckInput() error
}

func CheckInput(s []string) error {
	if len(s) == 0 {
		err := errors.New("not enough arguments\nuse the following format: ./alarmclock HH:MM")
		return err
	} else if len(s) > 2 {
		err := errors.New("too many arguments\nuse the following format: ./alarmclock HH:MM")
		return err
	} else if len(s) == 2 {
		//check for am or pm

	}
	return nil
}
