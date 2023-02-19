// group cache 中的 防缓存击穿

package singleflight

import "sync"

// call is an in-flight or completed Do call
//  请求体，代表正在进行中，或已经结束的请求。使用 sync.WaitGroup 锁避免重入。
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// OnceFlightGroup represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
// 主体，维护一个call 结构体，管理不同的键
type OnceFlightGroup struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// 多次请求key，只会执行一次第二个方法fb参数，其他方法等待fn执行完，返回fn的返回值，实现防击穿
func (g *OnceFlightGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	//1. 如果该键存在，则取消互斥锁，阻塞线程 等待 请求体执行完 fn方法执行
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	//2.若m中不存在键，则创建一个请求体，并设置等待线程次数，加入到m中
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock() //完成上面操作后，取消互斥锁，其他线程能够走到步骤1中等待

	//3.执行更新缓存fn方法，并取消等待线程
	c.val, c.err = fn()
	c.wg.Done()

	//4.执行完毕，删除该键的请求体，回收内存
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
