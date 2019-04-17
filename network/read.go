package network

import (
	"log"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/gArch/protos"
)

func read(peer *Peer) {
	defer peer.conn.Close()

	data := make([]byte, 1024)

	for {
		len, err := peer.conn.Read(data)
		if err != nil {
			log.Println("Peer disconnected ->", peer.IP)
			peer.conn.Close()
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
			log.Println("Active Peers :->")
			for i := range activePeers {
				log.Print((*activePeers[i]).IP)
			}

		default:
			log.Println("Default message")

		}
	}
}
