// runtime_pprof/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int // nil
	for {
		select {
		case v := <-c: // 空通道取值 阻塞
			fmt.Printf("recv from chan, value:%v\n", v)
		default: // select 一直在这里调用default转 cpu永远在运行for循环
			time.Sleep(time.Millisecond * 500) // 优化获取不到值 先睡几秒钟
		}
	}
}

// go build main.go
// ./main -cpu=true
// go tool pprof cpu.pprof
// go tool pprof cpu_optimize.pprof

/*
wjj@wjjdeMBP 088pprof_demo % go tool pprof cpu.pprof
Type: cpu
Time: May 3, 2021 at 10:45pm (CST)
Duration: 20.15s, Total samples = 1.89mins (563.56%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 10
Showing nodes accounting for 113.48s, 99.95% of 113.54s total
Dropped 10 nodes (cum <= 0.57s)
      flat  flat%   sum%        cum   cum%
    45.59s 40.15% 40.15%     88.64s 78.07%  runtime.selectnbrecv
    33.96s 29.91% 70.06%     36.99s 32.58%  runtime.chanrecv
    24.84s 21.88% 91.94%    113.51s   100%  main.logicCode
     9.09s  8.01% 99.95%      9.12s  8.03%  runtime.newstack
*/
/*
wjj@wjjdeMBP 088pprof_demo % go tool pprof cpu_optimize.pprof
Type: cpu
Time: May 3, 2021 at 10:45pm (CST)
Duration: 20s, Total samples = 0
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 10
Showing nodes accounting for 0, 0% of 0 total
      flat  flat%   sum%        cum   cum%
*/
/*
top 10 cpu占用前10的函数
flat：当前函数占用CPU的耗时
flat：:当前函数占用CPU的耗时百分比
sun%：函数占用CPU的耗时累计百分比
cum：当前函数加上调用当前函数的函数占用CPU的总耗时
cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
最后一列：函数名称
*/
// list 查看某个方法具体代码执行时间
/*
wjj@wjjdeMBP 088pprof_demo % go tool pprof cpu.pprof
Type: cpu
Time: May 3, 2021 at 10:45pm (CST)
Duration: 20.15s, Total samples = 1.89mins (563.56%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 10
Showing nodes accounting for 113.48s, 99.95% of 113.54s total
Dropped 10 nodes (cum <= 0.57s)
      flat  flat%   sum%        cum   cum%
    45.59s 40.15% 40.15%     88.64s 78.07%  runtime.selectnbrecv
    33.96s 29.91% 70.06%     36.99s 32.58%  runtime.chanrecv
    24.84s 21.88% 91.94%    113.51s   100%  main.logicCode
     9.09s  8.01% 99.95%      9.12s  8.03%  runtime.newstack
(pprof) list logicCode
Total: 1.89mins
ROUTINE ======================== main.logicCode in /Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/088pprof_demo/main.go
    24.84s   1.89mins (flat, cum)   100% of Total
         .          .      8:   "runtime/pprof"
         .          .      9:   "time"
         .          .     10:)
         .          .     11:
         .          .     12:// 一段有问题的代码
         .       10ms     13:func logicCode() {
         .          .     14:   var c chan int // nil
         .          .     15:   for {
         .          .     16:           select {
    24.84s   1.89mins     17:           case v := <-c: // 空通道取值 阻塞
         .          .     18:                   fmt.Printf("recv from chan, value:%v\n", v)
         .          .     19:           default: // select 一直在这里调用default转 cpu永远在运行for循环
         .          .     20:                   time.Sleep(time.Millisecond * 500) // 优化获取不到值 先睡几秒钟
         .          .     21:           }
         .          .     22:   }
*/

// web  web页面 查看

// 可视化  brew install graphviz
// dot -version 查看是否安装
// # 查看内存占用数据
// go tool pprof -inuse_space http://127.0.0.1:8080/debug/pprof/heap
// go tool pprof -inuse_objects http://127.0.0.1:8080/debug/pprof/heap
// # 查看临时内存分配数据
// go tool pprof -alloc_space http://127.0.0.1:8080/debug/pprof/heap
// go tool pprof -alloc_objects http://127.0.0.1:8080/debug/pprof/heap

func main() {
	var isCPUPprof bool // 是否开启CPU profile的标志位
	var isMemPprof bool // 是否开启内存 profile的标志位

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		f1, err := os.Create("./cpu_optimize.pprof") // 在当前路径下创建一个cpu.pprof文件
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(f1) // 往文件中记录CPU profile信息
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}
