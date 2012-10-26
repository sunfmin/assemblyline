package main

import (
	"log"
	"time"
)

var CurrentWorkTime = map[string]*Worktime{
	"WashWaterPot": &Worktime{1 * time.Second, WashWaterPot},
	"BoilWater":    &Worktime{15 * time.Second, BoilWater},
	"WashTeaPot":   &Worktime{2 * time.Second, WashTeaPot},
	"PickTea":      &Worktime{1 * time.Second, PickTea},
	"MakePotOfTea": &Worktime{5 * time.Second, MakePotOfTea},
	"WashCup":      &Worktime{2 * time.Second, WashCup},
	"MakeCupOfTea": &Worktime{500 * time.Millisecond, MakeCupOfTea},
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
	Time   time.Duration
	Worker func(worktime time.Duration) `json:"-"`
}

var RC = NewRestartCommand()

func (rc RestartCommand) Restart(newconfig map[string]int, first bool) {

	log.Println("Restarting: ", newconfig)
	var finshed = make(chan int)
	for name, wt := range CurrentWorkTime {
		go func(name string, wt *Worktime) {
			newworkers := newconfig[name]
			if newworkers == 0 {
				newworkers = 1
			}

			if newworkers > 60 {
				newworkers = 60
			}
			gg := GGroup(name)
			runningWorkers := len(gg.Status)

			if newworkers == runningWorkers {
				finshed <- 1
				return
			}

			if newworkers > runningWorkers {
				added := newworkers - runningWorkers
				for i := 0; i < added; i++ {
					go wt.Worker(wt.Time)
				}
				finshed <- 1
				return
			}

			killed := (runningWorkers - newworkers)
			for i := 0; i < killed; i++ {
				go func() {
					rc.Quits[name] <- 1
				}()
			}
			finshed <- 1
			return
		}(name, wt)
	}
	i := 0

	for {
		_ = <-finshed
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
