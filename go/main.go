package main

import (
	"github.com/tebben/marvin/go/marvin/system"
	"github.com/tebben/marvin/go/marvin/modules"
	"github.com/tebben/marvin/go/marvin/modules/hue"
	"github.com/tebben/marvin/go/http"
	"github.com/tebben/marvin/go/marvin/models"
)

var marvin models.Marvin

func main() {
	marvin = system.CreateMarvin()
	marvin.AddModule(&huemodule.HueModule{})
	marvin.AddModule(&modules.PrintModule{})
	marvin.Start()

	marvinServer := http.CreateServer(&marvin, "localhost", 8080, marvin.GetEndpoints())
	marvinServer.Start()
}

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
