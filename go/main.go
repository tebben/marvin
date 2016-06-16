package main

import (
	"github.com/tebben/marvin/go/http"
	"github.com/tebben/marvin/go/marvin/models"
	"github.com/tebben/marvin/go/marvin/modules/hue"
	"github.com/tebben/marvin/go/marvin/modules/logger"
	"github.com/tebben/marvin/go/marvin/modules/neopixel"
	"github.com/tebben/marvin/go/marvin/system"
)

var marvin models.Marvin

func main() {
	marvin = system.CreateMarvin()
	marvin.AddModule(&huemodule.HueModule{})
	marvin.AddModule(&logmodule.LogModule{})
	marvin.AddModule(&neopixel.NeopixelModule{})
	marvin.Start()

	marvinServer := http.CreateServer(&marvin, "localhost", 8080, marvin.GetEndpoints())
	marvinServer.Start()
}
