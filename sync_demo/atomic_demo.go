package syncdemo

import (
	"fmt"
	"sync/atomic"
)

//
//
// 这个包主要介绍的是： 原子操作，主要用于多协程间 变量的操作。
// 对某些基本类型数字： 在原来的值上加一个数； 保存一个值到数中； 读取字段的值。
type AtomicDemo struct {

}
func NewAtomicDemo() *AtomicDemo {
	return &AtomicDemo{}
}
func (o *AtomicDemo) runCall() {
	fmt.Println("run atomic func: ")
	var x int32 = 0
	atomic.StoreInt32(&x, 100)
	fmt.Println("load: ", atomic.LoadInt32(&x))
	//
	atomic.AddInt32(&x, 20)
	fmt.Println("load: ", atomic.LoadInt32(&x))
}
func RunAtomicDemo() {
	r := NewAtomicDemo()
	//
	r.runCall()
}