package main

import (
	"fmt"
)

func CPUAssember() {
	for {
		cpu := <-CPULine
		computer := <-ComputersMissingCPU
		computer.C = cpu
		SendCommand("Computer.Move", Movement{"ComputerAssemblyLine", "ComputersMissingCPU", computer})

		ForDemoPause(CPUAssemberWorkTime)
		fmt.Println("assembled a cpu")
		ComputerAssemblyLine <- computer
		SendCommand("Computer.Move", Movement{"ComputersMissingCPU", "ComputerAssemblyLine", computer})
	}
}

func MemoryAssember() {
	for {
		memory := <-MemoryLine
		computer := <-ComputersMissingMemory
		computer.M = memory
		SendCommand("Computer.Move", Movement{"ComputerAssemblyLine", "ComputersMissingMemory", computer})
		ForDemoPause(MemoryAssemberWorkTime)
		fmt.Println("assembled a memory")
		ComputerAssemblyLine <- computer
		SendCommand("Computer.Move", Movement{"ComputersMissingMemory", "ComputerAssemblyLine", computer})
	}
}

func KeyboardAssember() {
	for {
		keyboard := <-KeyboardLine
		computer := <-ComputersMissingKeyboard
		computer.K = keyboard
		SendCommand("Computer.Move", Movement{"ComputerAssemblyLine", "ComputersMissingKeyboard", computer})
		ForDemoPause(KeyboardAssemberWorkTime)
		fmt.Println("assembled a keyboard")
		ComputerAssemblyLine <- computer
		SendCommand("Computer.Move", Movement{"ComputersMissingKeyboard", "ComputerAssemblyLine", computer})
	}
}

func ScreenAssember() {
	for {
		screen := <-ScreenLine
		computer := <-ComputersMissingScreen
		computer.S = screen
		SendCommand("Computer.Move", Movement{"ComputerAssemblyLine", "ComputersMissingScreen", computer})
		ForDemoPause(ScreenAssemberWorkTime)
		fmt.Println("assembled a screen")
		ComputerAssemblyLine <- computer
		SendCommand("Computer.Move", Movement{"ComputersMissingScreen", "ComputerAssemblyLine", computer})
	}
}
