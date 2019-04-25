package network

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/gArch/protos"
	types "github.com/sjljrvis/gArch/types"
)

func read(peer *types.Peer) {
	defer peer.Conn.Close()

	data := make([]byte, 1024)
	for {
		len, err := peer.Conn.Read(data)
		if err != nil {
			log.Println("Peer disconnected 2->", peer.IP)
			peer.Conn.Close()
			return
		}
		msg := &protos.Arc{}
		err = proto.Unmarshal(data[0:len], msg)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		switch msg.GetType() {

		case "handshake":
			peer.IP = (string(msg.GetData()))
			log.Println("Connected to Peer :->", string(msg.GetData()))
			data, err := peerStore.Set(peer.IP, peer)
			fmt.Sprint(data, err)
			// log.Println("Active Peers :->")

			// for i := range activePeers {
			// 	log.Print((*activePeers[i]).IP)
			// }
			announce()

		case "gossip":
			gossipIP := strings.Split(string(msg.GetData()), ",")
			fmt.Println("Check", gossipIP)
			for i := range gossipIP {
				if selfAddress != gossipIP[i] {
					connect(gossipIP[i])
				}
			}

		default:
			log.Println("Default message")

		}
	}
}

func announce() {

	var announceIP = []string{}

	peers := peerStore.All()

	for i := range peers {
		announceIP = append(announceIP, i)
	}

	_msg := strings.Join(announceIP, ",")

	if len(peers) > 1 {
		msg := &protos.Arc{
			Type: "gossip",
			Data: []byte(_msg),
		}
		data, err := proto.Marshal(msg)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		for i := range peers {
			peer, _ := peerStore.Get(i)
			peer.Conn.Write(data)
		}
	}

	// if len(activePeers) > 1 {

	// 	for i := range activePeers {
	// 		announceIP = append(announceIP, activePeers[i].IP)
	// 	}

	// 	// _buffer := &bytes.Buffer{}
	// 	// gob.NewEncoder(_buffer).Encode(announceIP)
	// 	// announceIPBytes := _buffer.Bytes()

	// 	_msg := strings.Join(announceIP, ",")
	// 	msg := &protos.Arc{
	// 		Type: "gossip",
	// 		Data: []byte(_msg),
	// 	}
	// 	data, err := proto.Marshal(msg)
	// 	if err != nil {
	// 		log.Fatal("marshaling error: ", err)
	// 	}
	// 	for i := range activePeers {
	// 		activePeers[i].Conn.Write(data)
	// 	}
	// }
}

func connect(IP string) {

	peer, err := peerStore.Get(IP)
	if err != nil {
		clientConn, err := net.Dial("tcp", IP)
		if err != nil {
			log.Println("Peer disconnected 3->", clientConn.RemoteAddr())
			clientConn.Close()
			return
		}
		fmt.Sprint(peer)
		peerChannel <- clientConn

	} else {
		log.Printf("Peer already dialed", len(IP), "<><><><>", IP, err)
	}

	// for i := range activePeers {
	// 	if activePeers[i].IP != IP {
	// 		clientConn, err := net.Dial("tcp", IP)
	// 		if err != nil {
	// 			log.Println("Peer disconnected 3->", clientConn.RemoteAddr())
	// 			clientConn.Close()
	// 			return
	// 		}
	// 		peerChannel <- clientConn
	// 	}
	// }
}
