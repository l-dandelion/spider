package pool

import (
	"sync"
)

type Pool struct {
	pool  chan bool      //控制数量
	group sync.WaitGroup //用于等待所有爬虫执行完毕
}

//增加爬虫
func (pool *Pool) Add() {
	pool.pool <- true
	pool.group.Add(1)
}

//爬虫结束
func (pool *Pool) Done() {
	<-pool.pool
	pool.group.Done()
}

//等待所有爬虫结束
func (pool *Pool) Wait() {
	pool.group.Wait()
}

func (pool *Pool) Init(max_num int) {
	pool.pool = make(chan bool, max_num)
}

//进程池中已分配的数量
func (pool *Pool) Num() int {
	return len(pool.pool)
}
