package main

import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
)

var arrWs [2]websocket.Conn

func Echo(ws *websocket.Conn) {
    var err error
    for {
        var reply string

        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive")
            break
        }

        fmt.Println("Received back from client: " + reply)

        var msg string = "Ask me something!!"
	if reply == "where are you from!!" {
		msg = "That is not important"
	}
	if reply == "How old are you" {
		msg = "I am 29 years old"
	}

        if arrWs[1] == range websocket {
		arrWs[1] = ws
	}
        //fmt.Println("Sending to client: " + msg)

        if err = websocket.Message.Send(ws, msg); err != nil {
            fmt.Println("Can't send")
            break
        }

        if err = websocket.Message.Send(ws[1], msg); err != nil {
            fmt.Println("Can't send")
            break
        }
    }
}

func main() {
    http.Handle("/sock", websocket.Handler(Echo))

    if err := http.ListenAndServe(":1588", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

