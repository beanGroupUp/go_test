package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享数据结构
type SafeCounter struct {
	mu      sync.RWMutex
	counter int
}

/**
在 sync.RWMutex 中，Lock() 和 RLock() 是两种不同类型的锁，它们的主要区别如下：

Lock() - 写锁（独占锁）
独占性：一次只允许一个 goroutine 获取写锁

阻塞效果：

当有 goroutine 持有写锁时，其他所有尝试获取读锁或写锁的 goroutine 都会被阻塞

当有 goroutine 持有读锁时，尝试获取写锁的 goroutine 会被阻塞，直到所有读锁释放

使用场景：用于写操作，确保数据修改的原子性和一致性

RLock() - 读锁（共享锁）
共享性：允许多个 goroutine 同时获取读锁

阻塞效果：

当有 goroutine 持有读锁时，其他 goroutine 仍然可以获取读锁

但当有 goroutine 持有读锁时，尝试获取写锁的 goroutine 会被阻塞

当有 goroutine 持有写锁时，尝试获取读锁的 goroutine 会被阻塞

使用场景：用于只读操作，允许多个读取者同时访问资源
*/

// 写操作-需要独占锁
func (sc *SafeCounter) Inc() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counter++
	fmt.Printf("写入：计数器增加为%d\n", sc.counter)
}

// 读操作 - 可以使用共享锁
func (sc *SafeCounter) GetValue() int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	value := sc.counter
	fmt.Printf("读取：当前值为%d\n", value)
	return value
}

func main() {
	counter := SafeCounter{counter: 0}

	//启动对个读goroutine
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				counter.GetValue()
				time.Sleep(10 * time.Second)
			}
		}(i)
	}

	//启动多个写goroutine
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				counter.Inc()
				time.Sleep(10 * time.Second)
			}
		}(i)
	}

	//等待所有goroutine完成
	time.Sleep(1 * time.Second)
	fmt.Printf("最终结果：%d\n", counter.GetValue())
}
