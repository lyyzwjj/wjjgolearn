package _86split_string

import (
	"reflect"
	"testing"
)

// go test 测试组
// 自测试 单独跑某个测试用例
// go test -run=TestSplit/case_3
// 测试覆盖率 go test -cover
// 测试函数覆盖率 100%
// 测试代码覆盖率 60%
// go 还提供了一个额外的 -coverprofile参数 用来将覆盖率相关的记录信息输出到一个文件
// go test -cover -coverprofile=c.out
// go tool cover -html=c.out  // go 的工具 go tool cover -help
// go test -cover -coverprofile=cover.out
// go tool cover -html=cover.out
// 绿色 覆盖到的 红色 未覆盖到的

func TestSplit(t *testing.T) {
	// 利用结构体包装测试参数
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	//testGroup := []testCase{
	//	testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
	//	testCase{"a:b:c", ":", []string{"a", "b", "c"}},
	//	testCase{"abcef", "bc", []string{"a", "ef"}},
	//	testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	//}
	/*testGroup := []testCase{
		{"babcbef", "b", []string{"", "a", "c", "ef"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcef", "bc", []string{"a", "ef"}},
		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got , tc.want){
			t.Fatalf("want:%v but got:%v", tc.want, got)
		}
	}*/

	testGroup := map[string]testCase{
		"case_1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": {"abcef", "bc", []string{"a", "ef"}},
		"case_4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but got:%#v\n", tc.want, got)
			}
		})
	}
}

/*func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败
		t.Errorf("want:%v but got:%v", want, ret)

	}
	fmt.Printf("%#v\n", ret)
}

func Test2Split(t *testing.T) {
	ret := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败
		t.Errorf("want:%v but got:%v", want, ret)
	}
	fmt.Printf("%#v\n", ret)
}

func Test3Split(t *testing.T) {
	ret := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败
		t.Fatalf("want:%v but got:%v", want, ret)
	}
	fmt.Printf("%#v\n", ret)
}
*/
