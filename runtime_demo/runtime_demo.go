package runtimedemo

import (
	"fmt"
	"runtime"
)

// 本章节主要说明 go runtime 相关的知识：

type RunTimeDemo struct {

}

func NewRunTimeDemo() *RunTimeDemo {
	return &RunTimeDemo{}
}
func RunDemo() {
	fmt.Println("call runtime funcs:")
	r := NewRunTimeDemo()
	//
	r.getNumCpu()
	r.GOMAXPROCS()
	r.print_stack()
}

func(o *RunTimeDemo) getNumCpu() {
	fmt.Println("cpu nums: ", runtime.NumCPU()) //获取当前进程能使用的cpu数目。
}
func (o* RunTimeDemo) GOMAXPROCS() {
	fmt.Println("set max cpu nums for using: ", runtime.GOMAXPROCS(0)) //设置能并行运行的最大 cpu 数。
	fmt.Println("set max cpu nums for using: ", runtime.GOMAXPROCS(4)) //如果参数值 < 1， 不改变当前设置，并返回已经设置的值。
	fmt.Println("set max cpu nums for using: ", runtime.GOMAXPROCS(0))
}

func (o *RunTimeDemo) print_stack() {
	fmt.Println("run print caller stack: ")
	var ch chan struct{} = make(chan struct{})
	go func() {
		<-ch
	}()

	data := make([]byte, 1024*4)
	n := runtime.Stack(data, true)  //获取当前栈空间，参数如果true,那么其他协程的栈空间也会被加到后面
	fmt.Println("stack len: ", n, ", buf: ", string(data))
	
	fmt.Println("only current goroutine stack: ")
	{
		data := make([]byte, 1024*4)
		n := runtime.Stack(data, false)  //获取当前栈空间，参数如果true,那么其他协程的栈空间也会被加到后面
		fmt.Println("stack len: ", n, ", buf: ", string(data))
	}

	close(ch)

}