package main

import uuidlib "github.com/satori/go.uuid"

func uuid() string {
	return uuidlib.Must(uuidlib.NewV4(), nil).String()
}
