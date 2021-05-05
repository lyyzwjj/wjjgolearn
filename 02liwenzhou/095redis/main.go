package main

// golang 操作redis
// go get -u github.com/go-redis/redis
import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "wjjzst.com:6379",
		Password: "wzzst310",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功")
	// zset
	key := "rank"
	items := []redis.Z{
		{Score: 99, Member: "PHP"},
		{Score: 96, Member: "Golang"},
		{Score: 97, Member: "Python"},
		{Score: 99, Member: "Java"},
	}
	// 把元素都追加到key
	redisdb.ZAdd(key, items...)
	// 给Golang + 10 分
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
	// 取分数最高的3个
	// zrevrange rank 0 2 withscores
	ret, err := redisdb.ZRevRangeWithScores(key, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的
	// ZREVRANGEBYSCORE key max min [WITHSCORES]
	// zrevrangebyscore rank 100 95 withscores
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(key, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
