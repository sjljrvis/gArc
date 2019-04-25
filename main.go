package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	files "github.com/sjljrvis/gArch/files"
	helpers "github.com/sjljrvis/gArch/helpers"
	network "github.com/sjljrvis/gArch/network"
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
	// initDB()
	// fetchPeers()
}

func main() {
	mac := helpers.GetMacAddress()
	log.Println("MAC Address ->", mac)
	portFlag := flag.Int("port", 3000, "Port to connect")
	peersFlag := flag.String("peers", "", "list of peers")
	flag.Parse()

	port := *portFlag
	peers := strings.Split(*peersFlag, ",")

	log.Println("Peers can connect to address -> ", "127.0.0.1:", port)

	go network.Init(port, peers)
	files.DirWatcher()
}
