package network

import (
	"log"
	"net"
	"strconv"

	"github.com/gogo/protobuf/proto"
	uuid "github.com/satori/go.uuid"
	protos "github.com/sjljrvis/gArch/protos"
	store "github.com/sjljrvis/gArch/store"
	types "github.com/sjljrvis/gArch/types"
)

var (
	peerStore   = store.Init(10)
	peerChannel = make(chan net.Conn)
	activePeers = []*types.Peer{}
	msgChannel  = make(chan []byte)
	activeIPs   = []string{}
	selfAddress = ""
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
	selfAddress = "127.0.0.1:" + strconv.Itoa(port)

	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}

	// go connector(listener)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Some error", err)
			}
			peerChannel <- conn
		}
	}()

	if len(peers) > 0 {
		go func() {
			for _, peer := range peers {
				if len(peer) > 0 {
					clientConn, err := net.Dial("tcp", peer)
					if err != nil {
						log.Println("Peer disconnected 1->", clientConn.RemoteAddr())
						clientConn.Close()
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

			var msgChan = make(chan []byte)
			peerID, _ := uuid.NewV4()
			peer := &types.Peer{Conn: conn, Active: true, Msg: msgChan, ID: peerID}
			activePeers = append(activePeers, peer)

			go func(peer *types.Peer) {
				handshake := initHandshake(selfAddress)
				peer.Msg <- handshake
			}(peer)

			go write(peer)
			go read(peer)

		}
	}

}
