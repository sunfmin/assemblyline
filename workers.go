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
	gid := fmt.Sprintf("g%d", NewId("WashWaterPot"))
	gg := GGroup("WashWaterPot")
	fmt.Println("Starting WashWaterPot ", gid)
	for {
		var waterPot WaterPot
		select {
		case _ = <-RC.Quits["WashWaterPot"]:
			fmt.Println("Quiting WashWaterPot ", gid)
			goto Exit
		default:
		}
		gg.Status[gid] = "running"
		waterPot = WaterPot{Name: "WaterPot", Id: NewId("WaterPot")}
		ForDemoPause(worktime)
		fmt.Println("washed a water pot")
		SendCommand("Thing.Move", Movement{"", "CleanWaterPots", waterPot})
		gg.Status[gid] = "stopped"
		CleanWaterPots <- waterPot
	}
Exit:
	delete(gg.Status, gid)
}

func BoilWater(worktime time.Duration) {
	gid := fmt.Sprintf("g%d", NewId("BoilWater"))
	gg := GGroup("BoilWater")
	fmt.Println("Starting BoilWater ", gid)
	for {
		select {
		case _ = <-RC.Quits["BoilWater"]:
			fmt.Println("Quiting BoilWater ", gid)
			goto Exit
		default:
		}

		waterPot := <-CleanWaterPots
		gg.Status[gid] = "running"
		ForDemoPause(worktime)
		waterPot.Boiled = true
		fmt.Println("boiled a water pot")
		SendCommand("Thing.Move", Movement{"CleanWaterPots", "BoiledWaterPots", waterPot})
		gg.Status[gid] = "stopped"
		BoiledWaterPots <- waterPot
	}
Exit:
	delete(gg.Status, gid)
	fmt.Println(gg.Status)
}

func WashTeaPot(worktime time.Duration) {
	gid := fmt.Sprintf("g%d", NewId("WashTeaPot"))
	gg := GGroup("WashTeaPot")
	fmt.Println("Starting WashTeaPot ", gid)
	for {
		select {
		case _ = <-RC.Quits["WashTeaPot"]:
			fmt.Println("Quiting WashTeaPot ", gid)
			goto Exit
		default:
		}

		gg.Status[gid] = "running"
		teaPot := TeaPot{Name: "TeaPot", Id: NewId("WaterPot")}

		ForDemoPause(worktime)
		fmt.Println("washed a tea pot")
		SendCommand("Thing.Move", Movement{"", "CleanTeaPots", teaPot})
		gg.Status[gid] = "stopped"

		CleanTeaPots <- teaPot
	}
Exit:
	delete(gg.Status, gid)
}

func PickTea(worktime time.Duration) {
	gid := fmt.Sprintf("g%d", NewId("PickTea"))
	gg := GGroup("PickTea")
	fmt.Println("Starting PickTea ", gid)
	for {
		select {
		case _ = <-RC.Quits["PickTea"]:
			fmt.Println("Quiting PickTea ", gid)
			goto Exit
		default:
		}

		gg.Status[gid] = "running"
		tea := Tea{Name: "Tea", Id: NewId("Tea")}

		ForDemoPause(worktime)
		fmt.Println("picked a tea")
		SendCommand("Thing.Move", Movement{"", "TeaBags", tea})
		gg.Status[gid] = "stopped"

		TeaBags <- tea
	}
Exit:
	delete(gg.Status, gid)
}

func MakePotOfTea(worktime time.Duration) {
	gid := fmt.Sprintf("g%d", NewId("MakePotOfTea"))
	gg := GGroup("MakePotOfTea")
	fmt.Println("Starting MakePotOfTea ", gid)
	for {

		select {
		case _ = <-RC.Quits["MakePotOfTea"]:
			fmt.Println("Quiting MakePotOfTea ", gid)
			goto Exit
		default:

		}

		teaPot := <-CleanTeaPots
		tea := <-TeaBags
		waterPot := <-BoiledWaterPots

		gg.Status[gid] = "running"
		potOfTea := PotOfTea{Tea: tea, WaterPot: waterPot, TeaPot: teaPot, Name: "PotOfTea", Id: NewId("PotOfTea")}
		ForDemoPause(worktime)
		fmt.Println("made pot of tea")
		SendCommand("Thing.Move", Movement{"", "PotsOfTea", potOfTea})
		gg.Status[gid] = "stopped"

		PotsOfTea <- potOfTea
	}
Exit:
	delete(gg.Status, gid)
}

func WashCup(worktime time.Duration) {
	gid := fmt.Sprintf("g%d", NewId("WashCup"))
	gg := GGroup("WashCup")
	fmt.Println("Starting WashCup ", gid)
	for {
		select {
		case _ = <-RC.Quits["WashCup"]:
			fmt.Println("Quiting WashCup ", gid)
			goto Exit
		default:
		}

		gg.Status[gid] = "running"
		cup := Cup{Name: "Cup", Id: NewId("Cup")}
		ForDemoPause(worktime)
		fmt.Println("washed a cup")
		SendCommand("Thing.Move", Movement{"", "CleanCups", cup})
		gg.Status[gid] = "stopped"

		CleanCups <- cup
	}
Exit:
	delete(gg.Status, gid)
}

func MakeCupOfTea(worktime time.Duration) {
	gid := fmt.Sprintf("g%d", NewId("MakeCupOfTea"))
	gg := GGroup("MakeCupOfTea")
	fmt.Println("Starting MakeCupOfTea ", gid)
	for {
		select {
		case _ = <-RC.Quits["MakeCupOfTea"]:
			fmt.Println("Quiting MakeCupOfTea ", gid)
			goto Exit
		default:
			potOfTea := <-PotsOfTea
			cup := <-CleanCups

			gg.Status[gid] = "running"
			cupOfTea := CupOfTea{Cup: cup, PotOfTea: potOfTea, Name: "CupOfTea", Id: NewId("CupOfTea")}
			ForDemoPause(worktime)
			fmt.Println("made cup of tea")
			SendCommand("Thing.Move", Movement{"", "CupsOfTea", cupOfTea})
			gg.Status[gid] = "stopped"

			CupsOfTea <- cupOfTea
		}
	}
Exit:
	delete(gg.Status, gid)
}
