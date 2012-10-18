package main

import (
	"fmt"
	"time"
)

type Computer struct {
	C *CPU
	S *Screen
	M *Memory
	K *Keyboard
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

var ComputersMissingCPU = make(chan Computer)
var ComputersMissingScreen = make(chan Computer)
var ComputersMissingMemory = make(chan Computer)
var ComputersMissingKeyboard = make(chan Computer)

var ComputerAssemblyLine = make(chan Computer)

var ComputersCompleted = make(chan Computer)

func ComputerSkeletonCreator() {
	for {
		fmt.Println("making a skeleton")
		ComputerAssemblyLine <- Computer{}
		time.Sleep(1 * time.Second)
	}
}

func MissingPartDetector() {
	for {
		computerNeedAssembly := <-ComputerAssemblyLine
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
			ComputersCompleted <- computerNeedAssembly
		}()

	}
}

func main() {
	go ComputerSkeletonCreator()
	go MissingPartDetector()
	go CPUMaker()
	go CPUAssember()
	go KeyboardMaker()
	go KeyboardAssember()
	go ScreenMaker()
	go ScreenAssember()
	go MemoryMaker()
	go MemoryAssember()

	for {
		computer := <-ComputersCompleted
		fmt.Println("■■■■■ Completed: ", computer)
	}
}
