package main

import (
	"bradreed.co.uk/posproxy/printer"
	"bradreed.co.uk/posproxy/server"
	"log"
)

func main() {
	printer, err := printer.GetPrinter()
	if err != nil {
		log.Fatalf("Could not get printer: %v", err)
		return
	}

	defer printer.Close()

	server.StartServer(printer)
}
