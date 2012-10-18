package main

import (
	"fmt"
	"time"
)

type Keyboard struct {
	SerialId string
}

type CPU struct {
	HZ string
}

type Screen struct {
	Size string
}

type Memory struct {
	Size string
}

var KeyboardLine = make(chan *Keyboard)
var CPULine = make(chan *CPU)
var ScreenLine = make(chan *Screen)
var MemoryLine = make(chan *Memory)

func CPUMaker() {
	for {
		fmt.Println("making a cpu")
		CPULine <- &CPU{"1800HZ"}
		time.Sleep(3 * time.Second)
	}
}

func KeyboardMaker() {
	for {
		fmt.Println("making a keyboard")
		KeyboardLine <- &Keyboard{"US Keyboard"}
		time.Sleep(1 * time.Second)
	}
}

func ScreenMaker() {
	for {
		fmt.Println("making a screen")
		ScreenLine <- &Screen{"1080P"}
		time.Sleep(2 * time.Second)
	}
}

func MemoryMaker() {
	for {
		fmt.Println("makding a memory")
		MemoryLine <- &Memory{"16G"}
		time.Sleep(2 * time.Second)
	}
}
