package main

import (
	"log"
	"time"
)

var CurrentWorkTime = []*WorkTime{
	{"WashWaterPot", 1 * time.Second, WashWaterPot},
	{"BoilWater", 15 * time.Second, BoilWater},
	{"WashTeaPot", 2 * time.Second, WashTeaPot},
	{"PickTea", 1 * time.Second, PickTea},
	{"MakePotOfTea", 5 * time.Second, MakePotOfTea},
	{"WashCup", 2 * time.Second, WashCup},
	{"MakeCupOfTea", 500 * time.Millisecond, MakeCupOfTea},
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

type WorkTime struct {
	Name   string
	Time   time.Duration
	Worker func(workTime time.Duration) `json:"-"`
}

var RC = NewRestartCommand()

func (rc RestartCommand) Restart(config map[string]int, first bool) {

	log.Println("Restarting: ", config)
	var finished = make(chan int)
	for _, wt := range CurrentWorkTime {
		go func(name string, wt *WorkTime) {
			workers := config[name]
			if workers == 0 {
				workers = 1
			}

			if workers > 60 {
				workers = 60
			}
			gg := GGroup(name)
			runningWorkers := len(gg.Status)

			if workers == runningWorkers {
				finished <- 1
				return
			}

			if workers > runningWorkers {
				added := workers - runningWorkers
				for i := 0; i < added; i++ {
					go wt.Worker(wt.Time)
				}
				finished <- 1
				return
			}

			killed := runningWorkers - workers
			for i := 0; i < killed; i++ {
				go func() {
					rc.Quits[name] <- 1
				}()
			}
			finished <- 1
			return
		}(wt.Name, wt)
	}
	i := 0

	for {
		_ = <-finished
		if i >= 6 {
			log.Println("Restarted All Finished")
			SendCommand("Restarted", CurrentWorkTime)
			return
		}
		i = i + 1
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
	go SendGoroutineStatus()

	for {
		<-CupsOfTea
	}
}
