package main

import (
	"fmt"
)

func CPUAssember() {
	for {
		cpu := <-CPULine
		computer := <-ComputersMissingCPU
		computer.C = cpu
		fmt.Println("assembled a cpu")
		ComputerAssemblyLine <- computer
	}
}

func MemoryAssember() {
	for {
		memory := <-MemoryLine
		computer := <-ComputersMissingMemory
		computer.M = memory
		fmt.Println("assembled a memory")
		ComputerAssemblyLine <- computer
	}
}

func KeyboardAssember() {
	for {
		keyboard := <-KeyboardLine
		computer := <-ComputersMissingKeyboard
		computer.K = keyboard
		fmt.Println("assembled a keyboard")
		ComputerAssemblyLine <- computer
	}
}

func ScreenAssember() {
	for {
		screen := <-ScreenLine
		computer := <-ComputersMissingScreen
		computer.S = screen
		fmt.Println("assembled a screen")
		ComputerAssemblyLine <- computer
	}
}
