package p2p

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	nat "github.com/fd/go-nat"
)

var (
	internalPort = 3000
)

// Forward Port
func Forward() {

	peersFlag := flag.String("peers", "", "peers to connect to")
	portFlag := flag.Int("port", 3000, "port to listen to")

	flag.Parse()

	peers := strings.Split(*peersFlag, ",")
	port := (*portFlag)

	gateway, err := nat.DiscoverGateway()
	if err != nil {
		log.Println("unable to discover gateway")
		return
	}

	internalIP, err := gateway.GetInternalAddress()
	if err != nil {
		log.Println("unable to fetch internal IP")
		return
	}

	externalIP, err := gateway.GetExternalAddress()
	if err != nil {
		log.Println("unable to fetch external IP")
		return
	}

	log.Println("protocol", gateway.Type(), "discovered gateway")

	log.Println("internal_ip", internalIP.String(), "external_ip", externalIP.String())

	externalPort, err := gateway.AddPortMapping("tcp", port, "noise", 1*time.Second)

	if err != nil {
		log.Println("cannot setup port mapping")
		return
	}

	log.Println("internal_port", port, "external_port", externalPort, "external port now forwards to your local port")

	if len(peers) > 0 {
		for _, peer := range peers {
			if len(peer) > 0 {
				clientConn, err := net.Dial("tcp", peer)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(clientConn)
				_, err = clientConn.Write([]byte("hello"))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	resolveAddr, err := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Println("Please check Error in TCP Address", err)
	}
	ln, err := net.ListenTCP("tcp", resolveAddr)
	conn, _ := ln.Accept()

	fmt.Println(conn.RemoteAddr())
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}

}
