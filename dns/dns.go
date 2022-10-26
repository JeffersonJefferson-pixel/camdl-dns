package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

type dnsDefinition struct {
	Meta  dnsMetaJSON
	Nodes []*enode.Node
}

type dnsMetaJSON struct {
	URL          string    `json:"url,omitempty"`
	Seq          uint      `json:"seq"`
	Sig          string    `json:"signature,omitempty"`
	Links        []string  `json:"links"`
	LastModified time.Time `json:"lastModified"`
}


func toEnr(privKeyHex string, seq uint64, host string, tcp uint16, udp uint16) *enode.Node {

	// key, _ := crypto.LoadECDSA(file)
	key, _ := crypto.HexToECDSA(privKeyHex)

	ip := net.ParseIP(host)

	var r enr.Record
	// might need other attributes
	if (seq > 0) {
		r.SetSeq(seq)
	}
	if len(ip) > 0 {
		r.Set(enr.IP(ip))
	}
	if udp != 0 {
		r.Set(enr.UDP(udp))
	}
	if tcp != 0 {
		r.Set(enr.TCP(tcp))
	}
	enode.SignV4(&r, key)
	n, err := enode.New(enode.ValidSchemes, &r)
	if err != nil {
		panic(err)
	}
	return n
} 

func writeTreeNodes(directory string, def *dnsDefinition) {
	ns := make(nodeSet, len(def.Nodes))
	ns.add(def.Nodes...)
	_, nodesFile := treeDefinitionFiles(directory)
	writeNodesJSON(nodesFile, ns)
}

func treeDefinitionFiles(directory string) (string, string) {
	meta := filepath.Join(directory, "enrtree-info.json")
	nodes := filepath.Join(directory, "nodes.json")
	return meta, nodes
}

func exit(err interface{}) {
	if err == nil {
		os.Exit(0)
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	var def dnsDefinition
	var dataDir = filepath.Join("data")
	var nodes []*enode.Node
	rawNodes := loadRawNodesJSON("data/raw-nodes.json")
	// maybe only need public key
	for _, rawNode := range rawNodes {
		var n = toEnr(rawNode.PrivKey, rawNode.Seq, rawNode.Host, rawNode.Tcp, rawNode.Udp)
		fmt.Println(n.String())
		nodes = append(nodes, n)
	}	

	def.Nodes = nodes
	writeTreeNodes(dataDir, &def)

	os.Exit(0)
}