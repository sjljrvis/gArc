package main

import (
	"log"
	"os"
	"time"

	files "github.com/sjljrvis/gArch/files"
	helpers "github.com/sjljrvis/gArch/helpers"
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
	// go p2p.Forward()
	files.DirWatcher()
}
