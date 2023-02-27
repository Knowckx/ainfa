package parallel

import (
	"fmt"
	"sync"
)

// wg 加上限制数量
type WaitPool struct {
	wg         sync.WaitGroup
	sem        chan int // 目前运作的协程
	MaxLimit   int      // 协程的最大限制数
	sync.Mutex          // 用户的数据可能需要锁
}

func NewWaitPool(max int) *WaitPool {
	out := &WaitPool{}
	out.sem = make(chan int, max)
	out.MaxLimit = max
	return out
}

func (t *WaitPool) Add(in int) {
	t.wg.Add(in)
	t.sem <- 1
	return
}

func (t *WaitPool) Done() {
	<-t.sem
	t.wg.Done()
	return
}

func (t *WaitPool) Wait() {
	t.wg.Wait()
	return
}

// 带入到worker里的这三个参数通常是必要的
type ProcessPool struct {
	*WaitPool
	Total int
	Index int
}

func GetPercentage(v1, v2 int) float32 {
	res := float32(v1) / float32(v2) * 100
	return res
}

func (t *ProcessPool) Process() string {
	out := fmt.Sprintf("%d/%d %.2f%%", t.Index, t.Total, GetPercentage(t.Index, t.Total))
	return out
}

func NewProcessPool(po *WaitPool, idx, total int) *ProcessPool {
	out := &ProcessPool{}
	out.Index = idx
	out.Total = total
	out.WaitPool = po
	return out
}
