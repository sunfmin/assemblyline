package main

import (
	"fmt"
	"time"
)

var CleanWaterPots = make(chan WaterPot)
var BoiledWaterPots = make(chan WaterPot)
var CleanTeaPots = make(chan TeaPot)
var CleanCups = make(chan Cup)
var TeaBags = make(chan Tea)
var PotsOfTea = make(chan PotOfTea)
var CupsOfTea = make(chan CupOfTea)

func WashWaterPot(worktime time.Duration) {
	for {
		var waterPot WaterPot
		select {
		case _ = <-RC.Quits["WashWaterPot"]:
			break
		default:
		}

		waterPot = WaterPot{Name: "WaterPot", Id: NewId("WaterPot")}
		ForDemoPause(worktime)
		fmt.Println("washed a water pot")
		SendCommand("Thing.Move", Movement{"", "CleanWaterPots", waterPot})
		CleanWaterPots <- waterPot
	}
}

func BoilWater(worktime time.Duration) {
	for {
		select {
		case _ = <-RC.Quits["BoilWater"]:
			break
		default:
		}

		waterPot := <-CleanWaterPots

		ForDemoPause(worktime)
		waterPot.Boiled = true
		fmt.Println("boiled a water pot")
		SendCommand("Thing.Move", Movement{"CleanWaterPots", "BoiledWaterPots", waterPot})
		BoiledWaterPots <- waterPot
	}
}

func WashTeaPot(worktime time.Duration) {
	for {
		select {
		case _ = <-RC.Quits["WashTeaPot"]:
			break
		default:
		}

		teaPot := TeaPot{Name: "TeaPot", Id: NewId("WaterPot")}

		ForDemoPause(worktime)
		fmt.Println("washed a tea pot")
		SendCommand("Thing.Move", Movement{"", "CleanTeaPots", teaPot})
		CleanTeaPots <- teaPot
	}
}

func PickTea(worktime time.Duration) {
	for {
		select {
		case _ = <-RC.Quits["PickTea"]:
			break
		default:
		}

		tea := Tea{Name: "Tea", Id: NewId("Tea")}

		ForDemoPause(worktime)
		fmt.Println("picked a tea")
		SendCommand("Thing.Move", Movement{"", "TeaBags", tea})
		TeaBags <- tea
	}
}

func MakePotOfTea(worktime time.Duration) {
	for {

		select {
		case _ = <-RC.Quits["MakePotOfTea"]:
			break
		default:

		}

		teaPot := <-CleanTeaPots
		tea := <-TeaBags
		waterPot := <-BoiledWaterPots

		potOfTea := PotOfTea{Tea: tea, WaterPot: waterPot, TeaPot: teaPot, Name: "PotOfTea", Id: NewId("PotOfTea")}
		ForDemoPause(worktime)
		fmt.Println("made pot of tea")
		SendCommand("Thing.Move", Movement{"", "PotsOfTea", potOfTea})
		PotsOfTea <- potOfTea
	}
}

func WashCup(worktime time.Duration) {
	for {
		select {
		case _ = <-RC.Quits["WashCup"]:
			break
		default:
		}
		cup := Cup{Name: "Cup", Id: NewId("Cup")}

		ForDemoPause(worktime)
		fmt.Println("washed a cup")
		SendCommand("Thing.Move", Movement{"", "CleanCups", cup})
		CleanCups <- cup
	}
}

func MakeCupOfTea(worktime time.Duration) {
	for {
		select {
		case _ = <-RC.Quits["MakeCupOfTea"]:
			break
		default:
			potOfTea := <-PotsOfTea
			cup := <-CleanCups

			cupOfTea := CupOfTea{Cup: cup, PotOfTea: potOfTea, Name: "CupOfTea", Id: NewId("CupOfTea")}
			ForDemoPause(worktime)
			fmt.Println("made cup of tea")
			SendCommand("Thing.Move", Movement{"", "CupsOfTea", cupOfTea})
			CupsOfTea <- cupOfTea
		}
	}
}
