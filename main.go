package main

import (
	"fmt"
)

type Computer struct {
	Id int64
	C  *CPU
	S  *Screen
	M  *Memory
	K  *Keyboard
}

func (c Computer) String() (r string) {
	r = "Computer <"
	if c.C != nil {
		r = r + "=[CPU]"
	}
	if c.S != nil {
		r = r + "=[Screen]"
	}
	if c.M != nil {
		r = r + "=[Memory]"
	}
	if c.K != nil {
		r = r + "=[Keyboard]"
	}
	r = r + "=>"
	return
}

var ComputersMissingCPU = make(chan Computer, 10)
var ComputersMissingScreen = make(chan Computer, 10)
var ComputersMissingMemory = make(chan Computer, 10)
var ComputersMissingKeyboard = make(chan Computer, 10)

var ComputerAssemblyLine = make(chan Computer)

// var SkeletonLine = make(chan Computer)

var ComputersCompleted = make(chan Computer)

func ComputerSkeletonCreator() {
	for {
		fmt.Println("making a skeleton")
		computer := Computer{Id: NewId()}
		ComputerAssemblyLine <- computer
		SendCommand("ComputerAssemblyLine.Push", computer)
		ForDemoPause(SkeletonCreationTime)
	}
}

func MissingPartDetector() {
	for {
		computerNeedAssembly := <-ComputerAssemblyLine
		// ForDemoPause(100)

		fmt.Println("detecting computer: ", computerNeedAssembly)
		if computerNeedAssembly.C == nil {
			go func() {
				ComputersMissingCPU <- computerNeedAssembly
			}()
			continue
		}

		if computerNeedAssembly.S == nil {
			go func() {
				ComputersMissingScreen <- computerNeedAssembly
			}()
			continue
		}

		if computerNeedAssembly.M == nil {
			go func() {
				ComputersMissingMemory <- computerNeedAssembly

			}()
			continue
		}

		if computerNeedAssembly.K == nil {
			go func() {
				ComputersMissingKeyboard <- computerNeedAssembly
			}()

			continue
		}

		go func() {
			SendCommand("Computer.Move", Movement{"ComputerAssemblyLine", "ComputersCompleted", computerNeedAssembly})
			ComputersCompleted <- computerNeedAssembly
		}()

	}
}

func main() {
	go MemoryMaker()
	go CPUMaker()
	go KeyboardMaker()
	go ScreenMaker()
	go ComputerSkeletonCreator()
	go MissingPartDetector()

	for i := 0; i < CPUAssemberWorkers; i++ {
		go CPUAssember()
	}
	for i := 0; i < MemoryAssemberWorkers; i++ {
		go MemoryAssember()
	}
	for i := 0; i < KeyboardAssemberWorkers; i++ {
		go KeyboardAssember()
	}
	for i := 0; i < ScreenAssemberWorkers; i++ {
		go ScreenAssember()
	}

	go Server()

	for {
		computer := <-ComputersCompleted
		fmt.Println("■■■■■ Completed: ", computer)
	}
}
