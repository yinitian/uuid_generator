package main

import (
	"flag"
	"github.com/imkuqin-zw/id_generator/config"
	"github.com/imkuqin-zw/id_generator/rpc"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	rpc.Run()
}