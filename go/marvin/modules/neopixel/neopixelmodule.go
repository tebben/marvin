package neopixel

import (
	"github.com/tebben/marvin/go/marvin/connectors/serial"
	"github.com/tebben/marvin/go/marvin/models"
	"log"
)

const (
	moduleName        = "Neopixel"
	moduleDescription = "Control your neopixels on the Arduino"
)

type NeopixelModule struct {
	models.Module
	settings NeopixelSettings
	Port     *serial.Port
}

func (nm *NeopixelModule) Setup() {
	nm.Name = moduleName
	nm.Description = moduleDescription
	nm.Actions = []models.MarvinAction{
		CreateNeopixelOnAction(nm),
		CreateNeopixelOffAction(nm),
	}
	nm.settings = NeopixelSettings{
		PinNumber: 31,
		NrOfLeds:  3,
		Baud:      115200,
		ComPort:   "COM3",
	}

	go nm.StartSerial()
}

func (nm *NeopixelModule) StartSerial() {
	c := &serial.Config{Name: nm.settings.ComPort, Baud: nm.settings.Baud}
	var err error
	nm.Port, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan int, 1)
	go func() {
		buf := make([]byte, 128)
		var readCount int
		for {
			n, err := nm.Port.Read(buf)
			if err != nil {
				log.Print("ERROR")
			}
			readCount++
			log.Printf("Read %v %v bytes: % 02x %s\n", readCount, n, buf[:n], buf[:n])
			select {
			case <-ch:
				ch <- readCount
				close(ch)
			default:
			}
		}
	}()
}

func (nm *NeopixelModule) StopSerial() {
	nm.Port.Flush()
	nm.Port.Close()
}

func (nm *NeopixelModule) GetSettings() interface{} {
	return nm.settings
}

func (nm *NeopixelModule) SettingsChanged(newSettings interface{}) {
	log.Printf("SETTINGS CHANGED: %v\n", newSettings)
}
