package main

import (
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/config"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/engine"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/persist"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/scheduler"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
