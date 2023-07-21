package main

import (
	"sync"
	"time"
)

/*
*
1.读写锁(sync.RWMutex)
想象一下这种场景，当你在银行存钱或取钱时，对账户余额的修改是需要加锁的，因为这个时候，可能有人汇款到你的账户，如果对金额的修改不加锁，
很可能导致最后的金额发生错误。读取账户余额也需要等待修改操作结束，才能读取到正确的余额。大部分情况下，读取余额的操作会更频繁，如果能
保证读取余额的操作能并发执行，程序效率会得到很大地提高。
保证读操作的安全，那只要保证并发读时没有写操作在进行就行。在这种场景下我们需要一种特殊类型的锁，其允许多个只读操作并行执行，但写操作
会完全互斥。这种锁称之为 多读单写锁 (multiple readers, single writer lock)，简称读写锁，读写锁分为读锁和写锁，读锁是允许同时
执行的，但写锁是互斥的。一般来说，有如下几种情况:
1.读锁之间不互斥，没有写锁的情况下，读锁是无阻塞的，多个协程可以同时获得读锁。
2.写锁之间是互斥的，存在写锁，其他写锁阻塞。
3.写锁与读锁是互斥的，如果存在读锁，写锁阻塞，如果存在写锁，读锁阻塞。

Go 标准库中提供了 sync.RWMutex 互斥锁类型及其四个方法：
Lock 加写锁
Unlock 释放写锁
RLock 加读锁
RUnlock 释放读锁
*/
const cost = time.Microsecond

type RW interface {
	Write()
	Read()
}

// RWLock 读写锁
type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mu.RUnlock()
}

// Lock 互斥锁
type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}
