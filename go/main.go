package main

import (
	"github.com/tebben/marvin/go/marvin"
	"github.com/tebben/marvin/go/modules"
	//"github.com/tebben/marvin/go/serial"
	"golang.org/x/net/websocket"
	"net/http"
	"github.com/tebben/marvin/go/models"
	"log"
)

var Marvin marvin.Marvin

func main() {
	//Setup marvin
	Marvin = marvin.Marvin{}
	Marvin.AddModule(&modules.HueModule{})
	Marvin.AddModule(&modules.PrintModule{})
	Marvin.Start()

/*
	c := &serial.Config{Name: "COM3", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
*/

	http.Handle("/action", websocket.Handler(actionSocketHandler))
	http.Handle("/", http.FileServer(http.Dir("./client")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func actionSocketHandler(ws *websocket.Conn) {
	var oMsg models.ActionMessage

	for {
		if err := websocket.JSON.Receive(ws, &oMsg); err != nil{
			log.Printf("Error receiving socket message: %v", err)
			break
		}

		Marvin.Trigger(oMsg.Action, oMsg.Payload)
	}
}
