package main

import (
	"sync"
	"testing"
)

// 对比读多写少
func BenchmarkReadMore(b *testing.B)   { benchmark(b, &Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 9, 1) }

// 对比读少写多
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }

// 对比读写均衡
func BenchmarkEqual(b *testing.B)   { benchmark(b, &Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B) { benchmark(b, &RWLock{}, 5, 5) }

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
