package main

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/ethereum/go-ethereum/p2p/dnsdisc"
	. "github.com/ethereum/go-ethereum/p2p/enode"
)

type mapResolver map[string]string

func TestImportRawNodes(t *testing.T) {

	rawNodes := loadRawNodesJSON("test-data/raw-nodes.json")
	
	var (
		wantRawNodes = rawNodeSet{
			{2, "v4", "127.0.0.1", "1b366c21defe06a7a11a50545e949939f098de4764f51b9f14c87e7cc4fdc007", 30303, 30303},
			{1, "v4", "157.90.215.208", "65d833def30e9701c90256812a4faddb7f40f6a58f17a3a6f55df071cd1a9b60", 30303, 30303},
		}
	)
	if !reflect.DeepEqual(rawNodes, wantRawNodes) {
		t.Errorf("wrong raw node:\nhave %v\nwant %v", spew.Sdump(rawNodes), spew.Sdump(wantRawNodes))
	}
}

func TestResolve(t *testing.T) {
	r := mapResolver{
		"D35R6AYZQQV44SJTN4LL4IDHQQ.data": "enrtree-branch:T6BDSAJ42ULQIFBXCL37CG22PY,W7RFDFIQ7E7UTR7JGZ4PF5TL7A",
    	"FDXN3SN67NA5DKA4J2GOK7BVQI.data": "enrtree-branch:",
    	"T6BDSAJ42ULQIFBXCL37CG22PY.data": "enr:-Iu4QGRZGYS2msgyOA01pYHUxjZlvtx4l1ENcnfFj1lMi-HaE9NRGv86-zH5BpcDZZ92vnbchGJgbgozPLJyWzAMjekCgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQKaSG74cn8__e7KiDhshLx1_e30xc4XnflDPobHiBjb1YN0Y3CCdl-DdWRwgnZf",
    	"W7RFDFIQ7E7UTR7JGZ4PF5TL7A.data": "enr:-Iu4QBtBbjm2NeR-M1ALvFtllP-YdV7tGQh3JNeW1OAMYoFeYDITWGyh4ppblJcsGqx3punmKZrHTsC3rTzMPjWktGcBgmlkgnY0gmlwhJ1a19CJc2VjcDI1NmsxoQPTiisL4vylgrqkZplt-an_D7tUdgujhp5nFRUEPP1kI4N0Y3CCdl-DdWRwgnZf",
    	"data": "enrtree-root:v1 e=D35R6AYZQQV44SJTN4LL4IDHQQ l=FDXN3SN67NA5DKA4J2GOK7BVQI seq=13 sig=VkmSgPRxerjcMAONrvcE0ZB7mqinJSq-cU3HW5bgwUMeflQbOlZF3W4DPKzGSXyNFUYBlgb-40pJn2OdHMF65gE",
	}
	var (
		wantNodes = []*Node{
			toEnr("1b366c21defe06a7a11a50545e949939f098de4764f51b9f14c87e7cc4fdc007", 2, "127.0.0.1", 30303, 30303),
			toEnr("65d833def30e9701c90256812a4faddb7f40f6a58f17a3a6f55df071cd1a9b60", 1, "157.90.215.208", 30303, 30303),
		}
		wantSeq = uint(13)
	)

	c  := NewClient(Config{Resolver: r})
	stree, err := c.SyncTree("enrtree://APOUCYI6V23ZVM3PSRJAFCCR2F7I3XQSDL2WJMG6LDXYGHA2KY2FK@data")

	if err != nil {
		t.Fatal("sync error:", err)
	}
	if !reflect.DeepEqual((sortByID(stree.Nodes())), sortByID(wantNodes)) {
		t.Errorf("wrong nodes in synced tree:\nhave %v\nwant %v", spew.Sdump(stree.Nodes()), spew.Sdump(wantNodes))
	}
	if stree.Seq() != wantSeq {
		t.Errorf("synced tree has wrong seq: %d", stree.Seq())
	} 
}

func sortByID(nodes []*Node) []*Node {
	sort.Slice(nodes, func(i, j int) bool {
		return bytes.Compare(nodes[i].ID().Bytes(), nodes[j].ID().Bytes()) < 0
	})
	return nodes
}


func (mr mapResolver) LookupTXT(ctx context.Context, name string) ([]string, error) {
	if record, ok := mr[name]; ok {
		return []string{record}, nil
	}
	return nil, errors.New("not found")
}