package main

import (
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/p2p/enode"
)

const jsonIndent = "    "

// nodeSet is the nodes.json file format. It holds a set of node records
// as a JSON object.
type nodeSet map[enode.ID]nodeJSON

type nodeJSON struct {
	Seq uint64      `json:"seq"`
	N   *enode.Node `json:"record"`
}

func writeNodesJSON(file string, nodes nodeSet) {
	nodesJSON, err := json.MarshalIndent(nodes, "", jsonIndent)
	if err != nil {
		exit(err)
	}
	if file == "-" {
		os.Stdout.Write(nodesJSON)
		return
	}
	if err := os.WriteFile(file, nodesJSON, 0644); err != nil {
		exit(err)
	}
}

// add ensures the given nodes are present in the set.
func (ns nodeSet) add(nodes ...*enode.Node) {
	for _, n := range nodes {
		v := ns[n.ID()]
		v.N = n
		v.Seq = n.Seq()
		ns[n.ID()] = v
	}
}
