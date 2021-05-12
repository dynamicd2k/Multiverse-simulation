package main

import (
	"time"

	"github.com/iotaledger/multiverse-simulation/logger"
)

var log = logger.New("Simulation")

const nodesCount = 10000000

func main() {
	log.Info("Start simulation.... [done]")
	defer log.Info("Shutting donw simulation.... [done]")

	testNetwork := network.New(
		network.Nodes(nodesCount, multiverse.NewNode, network.ZipfDistribution(0.9, 10000000)),
		network.Delay(30*time.Millisecond, 250*time.Millisecond),
		network.PacketLoss(0, 0.05),
		network.Topology(network.WattsStrogatz(4, 1)),
	)
}
