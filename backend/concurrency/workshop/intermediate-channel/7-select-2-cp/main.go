package main

import (
	"fmt"
	"time"
)

type character struct {
	name            string
	activity        string
	defaultActivity string
}

func (c *character) awake(movementInput, attackInput chan string) {
	for {
		//sama seperti checkpoint sebelumnya
		select {
		//ketika menerima dari attackInput maka jalankan
		case c.activity = <-attackInput:
			fmt.Printf("%s melakukan serangan %s\n", c.name, c.activity)
		//ketika menerima dari movementInput maka jalankan
		case c.activity = <-movementInput:
			fmt.Printf("%s bergerak ke %s\n", c.name, c.activity)
		//tetapi sekarang menambahkan default activity yang menjalankan
		default:
			fmt.Printf("%s %s\n", c.name, c.defaultActivity)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
