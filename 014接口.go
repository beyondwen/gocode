package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	//t := T850{Name: "wenhao"}
	r := R2D2{Broken: false}
	err := Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("robot is on")
	}
	err = Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("机器人运行了")
	}
}
