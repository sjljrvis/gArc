package p2p

import (
	"flag"
	"log"
	"strings"

	"github.com/perlin-network/noise/crypto/ed25519"
	"github.com/perlin-network/noise/network"
	"github.com/perlin-network/noise/network/backoff"
	"github.com/perlin-network/noise/network/discovery"
	"github.com/perlin-network/noise/network/nat"
)

var (
	host     = "localhost"
	protocol = "tcp"
)

//Start is for getting public ip of the instance
func Start() {

	peersFlag := flag.String("peers", "", "peers to connect to")
	portFlag := flag.Int("port", 3000, "port to listen to")

	flag.Parse()

	port := uint16(*portFlag)
	peers := strings.Split(*peersFlag, ",")
	keys := ed25519.RandomKeyPair()
	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(protocol, host, port))

	nat.RegisterPlugin(builder)
	builder.AddPlugin(new(backoff.Plugin))
	builder.AddPlugin(new(discovery.Plugin))

	net, err := builder.Build()
	if err != nil {
		log.Fatal(err)
		return
	}

	go net.Listen()

	if len(peers) > 0 {
		net.Bootstrap(peers...)
	}

	select {}
}
