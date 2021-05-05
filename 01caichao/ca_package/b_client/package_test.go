package b_client

import (
	"github.com/wzzst310/wjjgolearn/01caichao/ca_package/a_series"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/30 3:36 下午
 * @description
 */

func TestPackage(t *testing.T) {
	t.Log(a_series.GetFibonacciSerie(5))
	// t.Log(a_series.square(5))  小写方法名无法在包外访问
	// t.Log(a_series.Square(5))  // 改成大写即可
}
