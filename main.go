package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gitzhang10/BFT/config"
	"github.com/gitzhang10/BFT/tusk"
)

var conf *config.Config
var err error

func init() {
	conf, err = config.LoadConfig("", "config")
	if err != nil {
		panic(err)
	}
}

func main() {
	if conf.Protocol == "tusk" {
		startTusk()
	} else {
		panic(errors.New("the protocol is unknown"))
	}
}

func startTusk() {
	node := tusk.NewNode(conf)
	if err = node.StartP2PListen(); err != nil {
		panic(err)
	}
	// wait for each node to start
	time.Sleep(time.Second * 15)
	if err = node.EstablishP2PConns(); err != nil {
		panic(err)
	}
	node.InitRBC(conf)
	fmt.Println("node starts the Tusk!")
	go node.RunLoop()
	go node.HandleMsgLoop()
	node.ConstructedBlockLoop()
}
