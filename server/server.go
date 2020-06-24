package server

import (
	"log"
	"net/http"

	"bradreed.co.uk/posproxy/printer"
	"github.com/gorilla/websocket"
)

type Server struct {
	Printer printer.Printer
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		s.Printer.Write(p)

		// echo the message back
		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}

func StartServer(p printer.Printer) {
	server := &Server{
		Printer: p,
	}
	http.HandleFunc("/", server.handler)

	log.Fatal(http.ListenAndServe(":6969", nil))
}
