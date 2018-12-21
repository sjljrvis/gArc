package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gortc/stun"
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
	log.Println("Fetch Peers from server")
}

func init() {
	initLogger()
	initDB()
	fetchPeers()
}

func main() {
	var port int
	c, err := stun.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		panic(err)
	}
	log.Println(stun.TransactionID, stun.BindingRequest)
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN server, waiting for response message.
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		port = xorAddr.Port
		fmt.Println("your IP is", xorAddr.IP, xorAddr.Port)
		go func() {
			p := make([]byte, 2048)
			addr := net.UDPAddr{
				Port: port,
				IP:   net.ParseIP("127.0.0.1"),
			}
			ser, err := net.ListenUDP("udp", &addr)
			if err != nil {
				fmt.Printf("Some error %v\n", err)
				return
			}
			for {
				_, remoteaddr, err := ser.ReadFromUDP(p)
				fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
				if err != nil {
					fmt.Printf("Some error  %v", err)
					continue
				}
			}
		}()
	}); err != nil {
		panic(err)
	}

	mac := helpers.GetMacAddress()
	log.Println("MAC Address ->", mac)
	go p2p.Start()
	routines.DirWatcher()
}
