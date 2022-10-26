package main

import "github.com/ethereum/go-ethereum/common"

// rawNodeSet is the raw-nodes.json file format. It holds a set of node records
// as a JSON object.
type rawNodeSet []rawNodeJSON

type rawNodeJSON struct {
	Seq 	uint64	`json:"seq"`
	Id  	string	`json:"id"`
	Host  	string	`json:"host"`
	PrivKey	string	`json:"privKey"`
	Udp  	uint16 	`json:"udp"`
	Tcp  	uint16 	`json:"tcp"`
}

func loadRawNodesJSON(file string) rawNodeSet {
	var rawNodes rawNodeSet
	if err := common.LoadJSON(file, &rawNodes); err != nil {
		exit(err)
	}
	return rawNodes
}
