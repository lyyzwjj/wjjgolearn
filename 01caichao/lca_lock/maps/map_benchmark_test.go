package maps

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/9/13 9:35 下午
 * @description
 */

type Condition struct {
	name        string
	NumOfReader int
	NumOfWriter int
}
type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func initConditionList() []Condition {
	conditionList := []Condition{
		{"读多写少", 10000, 10},
		{"读写一样", 100, 100},
		{"读少写多", 10, 10000}}
	return conditionList
}

func benchmarkMap(b *testing.B, hm Map, NumOfWriter int, NumOfReader int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncmap(b *testing.B) {
	conditionList := initConditionList()
	for _, condition := range conditionList { // 不关心返回值结果 但是也是需要占位
		fmt.Println("在", condition.name, "的情况下: ")
		b.Run("map with RWLock", func(b *testing.B) {
			hm := CreateRWLockMap()
			benchmarkMap(b, hm, condition.NumOfWriter, condition.NumOfReader)
		})

		b.Run("sync.map", func(b *testing.B) {
			hm := CreateSyncMapBenchmarkAdapter()
			benchmarkMap(b, hm, condition.NumOfWriter, condition.NumOfReader)
		})

		b.Run("concurrent map", func(b *testing.B) {
			superman := CreateConcurrentMapBenchmarkAdapter(199)
			benchmarkMap(b, superman, condition.NumOfWriter, condition.NumOfReader)
		})
	}
}
