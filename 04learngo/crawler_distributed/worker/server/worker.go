package main

import (
	"fmt"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/fetcher"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler_distributed/rpcsupport"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler_distributed/worker"

	"log"

	"flag"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
