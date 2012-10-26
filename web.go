package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Movement struct {
	From  string
	To    string
	Thing interface{}
}

type GoroutineGroup struct {
	Name   string
	Status map[string]string
}

var GoroutineGroups = make(map[string]*GoroutineGroup)

func GGroup(name string) (r *GoroutineGroup) {
	r = GoroutineGroups[name]

	if r != nil {
		return
	}

	r = &GoroutineGroup{
		Name:   name,
		Status: make(map[string]string),
	}
	GoroutineGroups[name] = r
	return
}

type Thing int

type Input struct {
	Workers map[string]int
}

type Output struct {
	MethodName string
	Workers    map[string]*Worktime
}

func (t Thing) UpdateConfig(ri *Input, reply *Output) (err error) {
	reply.MethodName = "Thing.UpdateConfig"
	RC.Restart(ri.Workers, false)
	reply.Workers = CurrentWorkTime
	return
}

func init() {
	rpc.Register(new(Thing))
}

func Server() {
	rand.Seed(int64(time.Now().Nanosecond()))

	http.Handle("/connect", websocket.Handler(Connect))
	err := http.ListenAndServe(":7890", nil)
	if err != nil {
		panic(err)
	}
}

func ForDemoPause(worktime time.Duration) {
	time.Sleep(worktime)
}

var Sockets []*websocket.Conn

type Command struct {
	MethodName string
	Data       interface{}
}

func Connect(ws *websocket.Conn) {
	Sockets = append(Sockets, ws)

	SendCommand("Connection.Ready", time.Now())
	jsonrpc.ServeConn(ws)

	var newsockets []*websocket.Conn
	for _, s := range Sockets {
		if s == ws {
			continue
		}
		newsockets = append(newsockets, s)
	}
	Sockets = newsockets
}

func SendCommand(methodName string, data interface{}) {
	cmd := Command{methodName, data}
	for _, s := range Sockets {
		err := websocket.JSON.Send(s, cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendGoroutineStatus() {
	for {
		SendCommand("Goroutine.Status", GoroutineGroups)
		time.Sleep(time.Millisecond * 500)
	}
}

var ids = make(map[string]chan int64)

func NewId(t string) (r int64) {
	if ids[t] == nil {
		ids[t] = make(chan int64)

		go func() {
			var i int64
			for {
				i = i + 1
				ids[t] <- i
				if i > 10000 {
					i = 0
				}
			}
		}()
	}
	r = <-ids[t]
	return
}
