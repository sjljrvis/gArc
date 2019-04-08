package network

import (
	"log"
	"net"
	"strconv"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/peerfind/protos"
)

// Peer Struct
type Peer struct {
	conn net.Conn
	msg  chan []byte
	IP   string
}

var (
	peerChannel = make(chan net.Conn)
	activePeers = make(map[net.Conn]bool)
	msgChannel  = make(chan []byte)
	activeIPs   = []string{}
)

func connector(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		peerChannel <- conn
	}
}

func initHandshake(ip string) []byte {
	msg := &protos.Arc{
		Type: "handshake",
		Data: []byte(ip),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return data
}

// Init connections and discovery here
func Init(port int, peers []string) {

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	selfAddress := "127.0.0.1:" + strconv.Itoa(port)

	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}

	go connector(listener)

	if len(peers) > 0 {
		go func() {
			for _, peer := range peers {
				if len(peer) > 0 {
					clientConn, err := net.Dial("tcp", peer)
					if err != nil {
						log.Println("Peer disconnected ->", clientConn.RemoteAddr())
						clientConn.Close()
						activePeers[clientConn] = false
						return
					}
					peerChannel <- clientConn
				}
			}
		}()
	}

	for {
		select {
		case conn := <-peerChannel:
			activePeers[conn] = true
			peer := &Peer{conn: conn}
			go peer.read(msgChannel)
			go peer.write(msgChannel)
			handshake := initHandshake(selfAddress)
			msgChannel <- handshake
		}
	}

}
