package main

import (
	"time"
)

var CurrentWorkTime = map[string]Worktime{
	"WashWaterPot": {1 * time.Second, 1, WashWaterPot},
	"BoilWater":    {6 * time.Second, 1, BoilWater},
	"WashTeaPot":   {2 * time.Second, 1, WashTeaPot},
	"PickTea":      {1 * time.Second, 1, PickTea},
	"MakePotOfTea": {5 * time.Second, 1, MakePotOfTea},
	"WashCup":      {2 * time.Second, 1, WashCup},
	"MakeCupOfTea": {0 * time.Second, 1, MakeCupOfTea},
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
	if !first {
		for name, wt := range CurrentWorkTime {
			for i := 0; i < wt.Workers; i++ {
				rc.Quits[name] <- 1
			}
		}
	}

	for name, workers := range newconfig {
		wt := CurrentWorkTime[name]
		wt.Workers = workers
		for i := 0; i < workers; i++ {
			go CurrentWorkTime[name].Worker(CurrentWorkTime[name].Time)
		}
	}
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

	for {
		<-CupsOfTea
	}
}
