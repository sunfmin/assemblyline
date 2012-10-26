package main

import (
	"fmt"
	"time"
)

var CurrentWorkTime = map[string]*Worktime{
	"WashWaterPot": &Worktime{1 * time.Second, 1, WashWaterPot},
	"BoilWater":    &Worktime{15 * time.Second, 1, BoilWater},
	"WashTeaPot":   &Worktime{2 * time.Second, 1, WashTeaPot},
	"PickTea":      &Worktime{1 * time.Second, 1, PickTea},
	"MakePotOfTea": &Worktime{5 * time.Second, 1, MakePotOfTea},
	"WashCup":      &Worktime{2 * time.Second, 1, WashCup},
	"MakeCupOfTea": &Worktime{0 * time.Second, 1, MakeCupOfTea},
}

var WorkerNames = []string{
	"WashWaterPot",
	"BoilWater",
	"WashTeaPot",
	"PickTea",
	"MakePotOfTea",
	"WashCup",
	"MakeCupOfTea",
}

type RestartCommand struct {
	Quits map[string]chan int
}

func NewRestartCommand() (r RestartCommand) {
	r = RestartCommand{}
	r.Quits = make(map[string]chan int)
	for _, name := range WorkerNames {
		r.Quits[name] = make(chan int)
	}
	return
}

type Worktime struct {
	Time    time.Duration
	Workers int
	Worker  func(worktime time.Duration) `json:"-"`
}

var RC = NewRestartCommand()

func (rc RestartCommand) Restart(newconfig map[string]int, first bool) {

	fmt.Println("Restarting: ", newconfig)
	for name, wt := range CurrentWorkTime {
		go func(name string, wt *Worktime) {
			if !first {
				for i := 0; i < wt.Workers; i++ {
					rc.Quits[name] <- 1
				}
			}
			wt.Workers = newconfig[name]
			if wt.Workers == 0 {
				wt.Workers = 1
			}
			for i := 0; i < wt.Workers; i++ {
				go wt.Worker(wt.Time)
			}
		}(name, wt)
	}

	SendCommand("Restarted", CurrentWorkTime)
}

func main() {
	RC.Restart(map[string]int{
		"WashWaterPot": 1,
		"BoilWater":    1,
		"WashTeaPot":   1,
		"PickTea":      1,
		"MakePotOfTea": 1,
		"WashCup":      1,
		"MakeCupOfTea": 1,
	}, true)

	go Server()
	go SendGoroutineStatus()

	for {
		<-CupsOfTea
	}
}
