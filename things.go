package main

type WaterPot struct {
	Id     int64
	Name   string
	Boiled bool
}

type TeaPot struct {
	Id   int64
	Name string
}

type Tea struct {
	Id   int64
	Name string
}

type PotOfTea struct {
	Id       int64
	Name     string
	Tea      Tea
	WaterPot WaterPot
	TeaPot   TeaPot
}

type Cup struct {
	Id   int64
	Name string
}

type CupOfTea struct {
	Id       int64
	Name     string
	Cup      Cup
	PotOfTea PotOfTea
}
