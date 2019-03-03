package main

import (
	"errors"
	"strconv"
)

func mustBeNumber(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("invalid number")
	}
	return nil
}

func mustBeString(input string) error {
	if input == "" {
		return errors.New("string must be entered")
	}
	return nil
}
