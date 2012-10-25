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

type Thing int

type Input struct {
	Workers map[string]int
}

type Output struct {
	Workers map[string]Worktime
}

func (t Thing) UpdateConfig(ri *Input, reply *Output) (err error) {
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
