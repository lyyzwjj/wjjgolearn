package fda_pool

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/9/1 1:50 上午
 * @description
 */
func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	/*if err := pool.ReleaseObj(&ReusableObj{}); err != nil { // 尝试放置超出线程池
		t.Error(err)
	}*/
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			// 不加上注解 等到取第10个的时候就1s钟超时了
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
