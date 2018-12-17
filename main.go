package main

import (
	"log"
	"time"

	helpers "github.com/sjljrvis/gArch/helpers"
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
	log.Println("Initializing node")
	log.Println("Creating fileDB ")
	log.Println("Sync nodes here")
}

func fetchPeers() {
	log.Println("Fetch Peers from server")
}

func init() {
	initLogger()
	initDB()
	fetchPeers()
}

func main() {
	mac := helpers.GetMacAddress()
	log.Println("MAC Address ->", mac)
	routines.DirWatcher()
}
