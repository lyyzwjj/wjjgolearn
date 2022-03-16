package parser

import (
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/config"
	"github.com/lyyzwjj/wjjgolearn/04learngo/crawler/engine"
	"regexp"
)

var carDetailRe = regexp.MustCompile(`<a href="(/m\d+/)" target="_blank"`)

func ParseCarModel(
	contents []byte, _ string) engine.ParseResult {
	matches := carDetailRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: host + string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCarDetail, config.ParseCarDetail),
			})
	}

	return result
}
