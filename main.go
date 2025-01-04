package main

import (
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
	"log"
)

func main() {
	svr := miner_core.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
