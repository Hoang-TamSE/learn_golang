package main

import "fmt"

func DoSomeThing(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThing: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}

	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}
