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

	peers := *peersFlag
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
		go func() {
			resolved, err := net.ResolveTCPAddr("tcp", peers)
			if err != nil {
				log.Panic(err)
			}
			// localAddr, _ := net.ResolveTCPAddr("tcp", externalIP.String()+":"+strconv.Itoa(externalPort))
			log.Println(">>>>>>>", "tcp", strconv.Itoa(externalPort))
			conn, _ := net.DialTCP("tcp", nil, resolved)
			fmt.Println("Hello", conn, resolved)
		}()
	}

	resolveAddr, err := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Println("Please check Error in TCP Address", err)
	}
	fmt.Println("...........", resolveAddr)
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
