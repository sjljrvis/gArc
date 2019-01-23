package main

import (
	"log"
	"os"
	"time"

	helpers "github.com/sjljrvis/gArch/helpers"
	p2p "github.com/sjljrvis/gArch/p2p"
	routines "github.com/sjljrvis/gArch/routines"
)

const (
	//TimeFormat *generic time stamp foramt used in logger
	TimeFormat = "Jan 2, 2006 at 3:04pm (MST)"
)

func initLogger() {
	log.SetFlags(0)
	log.SetPrefix("app | " + time.Now().Format(TimeFormat) + "| INFO : ")
}

func initDB() {
	log.Printf(os.Getenv("client"))
	log.Print(os.Getenv("destination"))
	log.Println("Initializing node")
	log.Println("Creating fileDB ")
	log.Println("Sync nodes here")
}

func fetchPeers() {
	log.Println("Fetch Peers from DHT")
}

func init() {
	initLogger()
	initDB()
	fetchPeers()
}

func main() {
	mac := helpers.GetMacAddress()
	log.Println("MAC Address ->", mac)
	go p2p.Start()
	routines.DirWatcher()
}
