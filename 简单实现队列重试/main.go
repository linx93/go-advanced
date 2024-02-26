package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// Task 表示队列中的一个任务
type Task struct {
	RetryCount int    // 重试次数
	Data       string // 任务数据
}

// Queue 是具有重试功能的任务队列
type Queue struct {
	tasks *list.List // 使用 list.List 作为队列
}

// NewQueue 创建一个新的 Queue 实例
func NewQueue() *Queue {
	return &Queue{
		tasks: list.New(),
	}
}

// Enqueue 向队列中添加一个任务
func (q *Queue) Enqueue(task *Task) {
	q.tasks.PushBack(task)
}

// Dequeue 从队列中取出一个任务
func (q *Queue) Dequeue() *Task {
	if q.tasks.Len() == 0 {
		return nil
	}
	e := q.tasks.Front()
	if e != nil {
		q.tasks.Remove(e)
	}

	return e.Value.(*Task)
}

// RetryQueue 尝试执行队列中的所有任务，如果失败则重新放入队列进行重试
func (q *Queue) RetryQueue(maxRetries int) {
	for q.tasks.Len() > 0 {
		task := q.Dequeue()
		if task == nil {
			continue
		}
		success := q.processTask(task, maxRetries)
		if !success {
			// 如果任务失败，增加重试计数并重新加入队列
			task.RetryCount++
			fmt.Printf("Task failed, retrying (attempt %d): %s\n", task.RetryCount, task.Data)
			q.Enqueue(task)
		}
	}
}

// processTask 尝试执行任务，并返回是否成功
func (q *Queue) processTask(task *Task, maxRetries int) bool {
	// 模拟任务执行，有一定的失败概率
	if rand.Intn(10) > 5 {
		return true // 成功
	}
	if task.RetryCount >= maxRetries {
		fmt.Printf("Task failed after %d retries: %s\n", task.RetryCount, task.Data)
		return false // 失败次数过多，不再重试
	}
	return false // 失败，需要重试
}

func main() {
	// 创建一个带有重试功能的队列
	queue := NewQueue()

	// 向队列中添加一些任务
	for i := 0; i < 10; i++ {
		task := &Task{
			RetryCount: 0,
			Data:       fmt.Sprintf("Task %d", i),
		}
		queue.Enqueue(task)
	}

	// 尝试执行队列中的所有任务，并处理失败的重试
	queue.RetryQueue(3)

	// 等待一段时间，让队列处理完成
	time.Sleep(1 * time.Second)

	fmt.Println("All tasks processed.")
}
