package syncdemo

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

//sync.Pool: 作用类似空闲队列。用于存放空闲的对象。注意：pool 中的空闲对象有可能随时被 回收。所以里面的内存不是可靠的。
// sync.Pool 是可被多协程 安全的使用。

type SyncPoolDemo struct {
	pool sync.Pool 
}
func NewSyncPoolDemo() *SyncPoolDemo {
	vLen := 100
	r := &SyncPoolDemo{
		pool : sync.Pool {
			New: func() any {
				d := make([]byte, vLen)
				return &d
				//return new(bytes.Buffer)
			},
		},
	}
	return r
}

func RunSyncPoolDemo() {
	fmt.Println("run sync.Pool: ")
	r := NewSyncPoolDemo() 
	//多协程下运行：
	var wg sync.WaitGroup 

	for i := 0; i < 10; i++ {
		wg.Add(1) 
		//
		go func() {
			defer wg.Done()
			
			v := r.pool.Get().(*[]byte) //取出来的是地址。
			if v == nil {
				fmt.Println("not get addr.")
				return 
			}
			// 需要初始化这些空间。
		
			// do busi logic...
			fmt.Println("len: ", len(*v))
			copy(*v, []byte("aaaa"))
			time.Sleep(10*time.Millisecond)


			r.pool.Put(v)
		}()
	}
	wg.Wait()
	fmt.Println("done.")
}

///
type BytesBuffPool struct {
	capacity int32 
	pool sync.Pool 
}

func RunBytesBuffPool() {
	fmt.Println("run RunBytesBuffPool: ")
	t := NewBytesBuffPool(256)

	for i := 0; i < 100; i++ {
		v := t.Get()
		///
		fmt.Println("cap: ", v.Cap(), ", len: ", v.Len())

		v.Write([]byte("this is demo"))
		//
		t.Put(v)
	}
	
}
func NewBytesBuffPool(c int32) *BytesBuffPool {
	r := &BytesBuffPool{
		capacity: c,
		pool : sync.Pool {
			New: func() any {
				t:= new(bytes.Buffer)
				t.Grow(int(c))
				return t
			},
		},
	}
	return r
}

func (o *BytesBuffPool) Get() *bytes.Buffer{
	t := o.pool.Get().(*bytes.Buffer)
	////////需要重置存储空间。
	t.Reset()
	return t
}


func (o *BytesBuffPool) Put(b* bytes.Buffer) {
	if b == nil {
		return 
	}
	if b.Cap() <= int(o.capacity) {
		o.pool.Put(b)
	}
}