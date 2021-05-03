package _86split_string

import (
	"fmt"
	"os"
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

// 基准测试
// go test -bench=. 运行所有的
// go test -bench=Split   (函数名 Split)
// -benchmem 查看每次操作内存情况 go test -bench=Split -benchmem

// 默认情况下 每个基准测试最少运行1秒.如果在Benchmark函数返回时没有到1秒,则b.N的值就会按1,2,5,10,20,50,...增加,并且函数再次运行
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

/*
wjj@wjjdeMacBook-Pro 086split_string % go test -bench=Split
goos: darwin
goarch: amd64
pkg: wjjgolearn/02liwenzhou/086split_string
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkSplit-8         2978284               397.4 ns/op
PASS
ok      wjjgolearn/02liwenzhou/086split_string  1.604s
*/
/*
BenchmarkSplit-8    8表示GOMAXPROCS 核心数  397.4 ns/op 每次操作耗费秒数
*/

/*
wjj@wjjdeMacBook-Pro 086split_string % go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: wjjgolearn/02liwenzhou/086split_string
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkSplit-8         2634758               430.2 ns/op           240 B/op          4 allocs/op
PASS
ok      wjjgolearn/02liwenzhou/086split_string  1.607s
*/

/*
240 B/op   	每次申请内存大小
4 allocs/op	申请内存次数
*/

/* 内存优化完之后
wjj@wjjdeMacBook-Pro 086split_string % go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: wjjgolearn/02liwenzhou/086split_string
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkSplit-8         7367452               175.0 ns/op            80 B/op          1 allocs/op
PASS
ok      wjjgolearn/02liwenzhou/086split_string  1.466s
*/

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

// go test -bench=Fib2
func BenchmarkFib1(b *testing.B) {
	// benchmarkFib(b, b.N) 绝对不允许 用b.N随便当参数
	benchmarkFib(b, 1)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}
func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}

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

	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作
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

// 重置时间 b.ResetTimer() 有些准备工作或其他的不应该被统计 见 gba_benchmark

// 并行测试  有使用goroutine的情况下测试
// -cpu 指定使用的CPU数量
// go test -bench=. -cpu 1
// go test -bench=Split -cpu 1
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

// Setup与TearDown
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

// 示例函数 以Example为前缀 利用 go dock 工具
// 三大好处
// 1. 示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
// 2. 示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。  go test -run Example
// 3. 示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。下图为strings.ToUpper函数在Playground的示例函数效果。
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
