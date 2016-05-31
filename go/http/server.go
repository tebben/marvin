package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tebben/marvin/go/marvin/models"
	"log"
	netHTTP "net/http"
	"strconv"
)

// MarvinServer is the type that contains all of the relevant information to set
// up the Marvin HTTP Server
type MarvinServer struct {
	marvin *models.Marvin
	host string // Hostname for example "localhost" or "192.168.1.14"
	port int    // Portnumber where you want to run your http server on
	endpoints []models.Endpoint // Configured endpoints for Marvin HTTP
}

// CreateServer initialises a new Marvin HTTPServer based on the given parameters
func CreateServer(marvin *models.Marvin, host string, port int, endpoints []models.Endpoint) models.HTTPServer {
	return &MarvinServer{
		marvin: marvin,
		host: host,
		port: port,
		endpoints: endpoints,
	}
}

// Start command to start the GOST HTTPServer
func (ms *MarvinServer) Start() {
	go startWebsockets()

	log.Printf("Started Marvin HTTP Server on %v:%v", ms.host, ms.port)
	router := createRouter(ms)
	httpError := netHTTP.ListenAndServe(ms.host+":"+strconv.Itoa(ms.port), router)

	if httpError != nil {
		log.Fatal(httpError)
		return
	}
}

func startWebsockets(){
	go hub.run()
	netHTTP.HandleFunc("/ws", serveWs)

	err := netHTTP.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Stop command to stop the Marvin HTTP server, currently not supported
func (ms *MarvinServer) Stop() {

}

func createRouter(ms *MarvinServer) *httprouter.Router {
	router := httprouter.New()
	for _, endpoint := range ms.endpoints{
		ep := endpoint
		for _, op := range ep.GetOperations() {
			operation := op
			if operation.Handler == nil {
				continue
			}

			switch operation.OperationType {
				case models.HTTPOperationGet : {
					router.GET(operation.Path, func(w netHTTP.ResponseWriter, r *netHTTP.Request, p httprouter.Params) { operation.Handler(w, r, ms.marvin)})
				}
			}
		}
	}

	router.ServeFiles("/app/*filepath", netHTTP.Dir("client/app"))
	return router
}

/*
func actionSocketHandler(ws *websocket.Conn) {
	var oMsg models.ActionMessage

	for {
		if err := websocket.JSON.Receive(ws, &oMsg); err != nil {
			log.Printf("Error receiving socket message: %v", err)
			break
		}
		log.Printf("Error receiving socket message: %v", oMsg.Action)
		//system.Marvin.Trigger(oMsg.Action, oMsg.Payload)
	}
}*/
