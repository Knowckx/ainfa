package main

import (
	"fmt"
	"time"

	"github.com/Knowckx/ainfa/parallel"
)

/* 测试并发 */

// 包装一下需要放到协程的参数 | 用户自己声明
type InputArgs struct {
	*parallel.ProcessPool
	Args string // 带入的参数
}

func main() {
	total := 5                    // 任务总数
	wg := parallel.NewWaitPool(2) // 1 声明同时跑的最大的协程数量
	for i := 0; i < total; i++ {
		wg.Add(1) // 2 协程数量+1, 超过限制会卡住
		in := &InputArgs{
			ProcessPool: parallel.NewProcessPool(wg, i+1, total),
			Args:        fmt.Sprintf("Test Value %d", i+1),
		}
		go Woker(in)
	}
	wg.Wait()
	fmt.Println("All Work Finidshed")
}

func Woker(in *InputArgs) {
	defer in.Done() // 3 释放
	fmt.Printf("Working on the task InputArgs:%s, process %s.\n", in.Args, in.Process())
	time.Sleep(3 * time.Second)
	fmt.Println("Working Finidsh", in.Args)
}
