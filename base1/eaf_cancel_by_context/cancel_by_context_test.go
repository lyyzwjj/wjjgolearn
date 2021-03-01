package eaf_cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/9/1 12:02 上午
 * @description
 */
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestChannel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled") //所有的协程都收到了 类似于广播机制
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
