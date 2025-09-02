package main

import (
	"fmt"
	"sync"
	"time"
)

type Goods struct {
	v map[string]int
	m sync.Mutex
	sync.RWMutex
}

func (g *Goods) Inc(key string, num int) {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("%v库存增加，已加锁\n", g.v[key])
	g.v[key]++
	fmt.Printf("%v库存增加完毕，解锁\n", g.v[key])
}

func (g *Goods) value(key string) int {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("库存上锁\n")
	return g.v[key]
}

func main() {
	mutext := sync.Mutex{}
	g := Goods{
		v: make(map[string]int),
		m: mutext,
	}

	for i := 0; i < 10; i++ {
		go g.Inc("榴莲", i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(g.value("榴莲"))

}
