package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})

	go func() {
		//工作，可理解为这里处理业务
		fmt.Println("工作中。。。")
		time.Sleep(3 * time.Second)

		//关闭退出信号，可理解为处理完业务之后退出
		close(quit)
	}()

	//这里阻塞读取quit通道，直到收到退出信号
	<-quit
	//执行退出处理
	fmt.Println("收到退出信号，推出中。。。")
	//在go的context包中，Context接口中Done()方法返回的就是一个通道信号，用于通知相关操作完成状态，其返回的就是一个空结构体,源码如下:

	type Context interface {
		// Deadline returns the time when work done on behalf of this context
		// should be canceled. Deadline returns ok==false when no deadline is
		// set. Successive calls to Deadline return the same results.
		Deadline() (deadline time.Time, ok bool)

		// Done returns a channel that's closed when work done on behalf of this
		// context should be canceled. Done may return nil if this context can
		// never be canceled. Successive calls to Done return the same value.
		// The close of the Done channel may happen asynchronously,
		// after the cancel function returns.
		//
		// WithCancel arranges for Done to be closed when cancel is called;
		// WithDeadline arranges for Done to be closed when the deadline
		// expires; WithTimeout arranges for Done to be closed when the timeout
		// elapses.
		//
		// Done is provided for use in select statements:
		//
		//  // Stream generates values with DoSomething and sends them to out
		//  // until DoSomething returns an error or ctx.Done is closed.
		//  func Stream(ctx context.Context, out chan<- Value) error {
		//  	for {
		//  		v, err := DoSomething(ctx)
		//  		if err != nil {
		//  			return err
		//  		}
		//  		select {
		//  		case <-ctx.Done():
		//  			return ctx.Err()
		//  		case out <- v:
		//  		}
		//  	}
		//  }
		//
		// See https://blog.golang.org/pipelines for more examples of how to use
		// a Done channel for cancellation.
		Done() <-chan struct{}

		// If Done is not yet closed, Err returns nil.
		// If Done is closed, Err returns a non-nil error explaining why:
		// Canceled if the context was canceled
		// or DeadlineExceeded if the context's deadline passed.
		// After Err returns a non-nil error, successive calls to Err return the same error.
		Err() error

		// Value returns the value associated with this context for key, or nil
		// if no value is associated with key. Successive calls to Value with
		// the same key returns the same result.
		//
		// Use context values only for request-scoped data that transits
		// processes and API boundaries, not for passing optional parameters to
		// functions.
		//
		// A key identifies a specific value in a Context. Functions that wish
		// to store values in Context typically allocate a key in a global
		// variable then use that key as the argument to context.WithValue and
		// Context.Value. A key can be any type that supports equality;
		// packages should define keys as an unexported type to avoid
		// collisions.
		//
		// Packages that define a Context key should provide type-safe accessors
		// for the values stored using that key:
		//
		// 	// Package user defines a User type that's stored in Contexts.
		// 	package user
		//
		// 	import "context"
		//
		// 	// User is the type of value stored in the Contexts.
		// 	type User struct {...}
		//
		// 	// key is an unexported type for keys defined in this package.
		// 	// This prevents collisions with keys defined in other packages.
		// 	type key int
		//
		// 	// userKey is the key for user.User values in Contexts. It is
		// 	// unexported; clients use user.NewContext and user.FromContext
		// 	// instead of using this key directly.
		// 	var userKey key
		//
		// 	// NewContext returns a new Context that carries value u.
		// 	func NewContext(ctx context.Context, u *User) context.Context {
		// 		return context.WithValue(ctx, userKey, u)
		// 	}
		//
		// 	// FromContext returns the User value stored in ctx, if any.
		// 	func FromContext(ctx context.Context) (*User, bool) {
		// 		u, ok := ctx.Value(userKey).(*User)
		// 		return u, ok
		// 	}
		Value(key any) any
	}
}
