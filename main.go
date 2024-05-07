package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mariuspod/node_health_check/node_health_check"
)

func main() {
	healthy := healthCheck()
	if healthy {
		log.Println("node is healthy")
		os.Exit(0)
	} else {
		log.Println("node is unhealthy")
		os.Exit(1)
	}
}

func healthCheck() bool {
	simulateHealth, ok := os.LookupEnv("HC_SIMULATE_HEALTH")
	if ok {
		log.Printf("Simulating node health with %s\n", simulateHealth)
		if simulateHealth == "true" {
			return true
		} else {
			return false
		}
	}

	maxBlockLagEnv, ok := os.LookupEnv("HC_MAX_BLOCK_LAG")
	if !ok {
		log.Fatal("Please provide $HC_MAX_BLOCK_LAG")
	}

	maxBlockLag, err := strconv.Atoi(maxBlockLagEnv)
	if err != nil {
		log.Fatal(err.Error())
	}
	selfUrl, ok := os.LookupEnv("HC_SELF_URL")
	if !ok {
		log.Fatal("Please provide $HC_SELF_URL")
	}

	syncing := node_health_check.GetSyncingState(selfUrl)
	if syncing {
		log.Println("node is still syncing")
		return false
	}

	peerUrlsEnv, ok := os.LookupEnv("HC_PEER_URLS")
	if !ok {
		log.Fatal("Please provide $HC_PEER_URLS")
	}
	peerUrls := strings.Split(peerUrlsEnv, ",")

	highestBlock := 0
	for _, u := range peerUrls {
		block := node_health_check.GetBlockHeight(u)
		if block > highestBlock {
			highestBlock = block
		}
	}
	currentBlock := node_health_check.GetBlockHeight(selfUrl)
	if (highestBlock - currentBlock) > maxBlockLag {
		log.Printf("node lags behind more than %d blocks (highest: %d, current: %d)\n", maxBlockLag, highestBlock, currentBlock)
		return false
	}

	return true
}
