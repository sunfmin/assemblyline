package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"math/rand"
	"net/http"
	"net/rpc/jsonrpc"
	"time"
)

type Movement struct {
	From     string
	To       string
	Computer Computer
}

func Server() {
	rand.Seed(int64(time.Now().Nanosecond()))

	http.Handle("/connect", websocket.Handler(Connect))
	err := http.ListenAndServe(":7890", nil)
	if err != nil {
		panic(err)
	}
}

func ForDemoPause(millseconds int64) {
	time.Sleep(time.Duration(millseconds) * time.Millisecond)
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

var ids chan int64

func NewId() (r int64) {
	if ids == nil {
		ids = make(chan int64)

		go func() {
			var i int64
			for {
				i = i + 1
				ids <- i
				if i > 10000 {
					i = 0
				}
			}
		}()
	}
	r = <-ids
	return
}
