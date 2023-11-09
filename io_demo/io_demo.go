package iodemo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

//io 包 提供了io原语的基本接口。
//io包 包装了一些已经存在的原语实现。
//io 包的接口 不是协程安全的。

type SrcReader struct {
	data  []byte
	index int32
}

func NewSrcReader() *SrcReader {
	r := &SrcReader{
		data:  []byte("this is demo test one"),
		index: 0,
	}
	return r
}
func (s *SrcReader) Read(p []byte) (n int, err error) {
	if s == nil {
		return 0, errors.New("is nil")
	}
	if s.index >= int32(len(s.data)) {
		return 0, io.EOF
	}
	ret := copy(p, s.data)
	s.index += int32(ret)
	return ret, nil
}

type DstWriter struct {
}

func NewDstWriter() *DstWriter {
	ret := &DstWriter{}
	return ret
}
func (o *DstWriter) Write(p []byte) (n int, err error) {
	fmt.Println("write data: ", string(p))
	return len(p), nil
}

func RunIoPackageDemo() {
	fmt.Println("run io package demo....")
	fmt.Println("const value: ", io.EOF)

	fmt.Println("copy: ")
	{
		src := strings.NewReader("this is read data")
		if n, e := io.Copy(os.Stdout, src); e != nil {
			fmt.Println("read data fail, e: ", e)
		} else {
			fmt.Println("read data succ, ret n: ", n)
		}
	}
	{
		if n, e := io.Copy(os.Stdout, NewSrcReader()); e != nil {
			fmt.Println("happen error, e: ", e)
		} else {
			fmt.Println("read data len: ", n)
		}
	}

	{
		if n, e := io.Copy(NewDstWriter(), NewSrcReader()); e != nil {
			fmt.Println("happen error, e: ", e)
		} else {
			fmt.Println("read data len: ", n)
		}
	}

	////: CopyN()
	{
		r := strings.NewReader("this is demo")
		if _, e := io.CopyN(os.Stdout, r, 4); e != nil {
			fmt.Println("fail e: ", e)
		} else {
			fmt.Println("copy n succ.")
		}
	}

	{
		r := strings.NewReader("thi")
		if _, e := io.CopyN(os.Stdout, r, 4); e != nil {
			fmt.Println("fail e: ", e)
		} else {
			fmt.Println("copy n succ.")
		}
	}

	//ReadAll()
	{
		r := strings.NewReader("this io.ReadAll() test..........................")
		dst, e := io.ReadAll(r)
		if e != nil {
			fmt.Println("", dst, ", error: ", e)
		} else {
			fmt.Println("", string(dst))
		}
	}

	{
		fmt.Println("io.ReadFull()")
		r := strings.NewReader("this io.ReadFull()")
		dst := make([]byte, r.Len())
		if _, e := io.ReadFull(r, dst); e != nil {
			errors.Is(e, io.EOF)
			fmt.Println("read fail, e: ", e, ", ", string(dst))
		} else {
			fmt.Println("read succ: ", string(dst))
		}
	}

	{
		r := strings.NewReader("11111111111111111111111isfdasdfa sfasd")
		r.WriteTo(os.Stdout)

		io.WriteString(os.Stdout, "adfadfadfafa1222222222222222222222222")
	}

	///
}

type IOWriterDemo struct{}

func (o *IOWriterDemo) Write(p []byte) (n int, err error) {
	fmt.Println("writ data : ", string(p))
	return len(p), nil
}
func NewIOWriterDemo() *IOWriterDemo {
	return &IOWriterDemo{}
}

func IoTypeDemoRun() {
	fmt.Println("io.Type demo: ") //这些接口主要是用在： 函数定义的参数或者 一些读写的逻辑上，为了考虑通用性，可以定义这些接口的变量。然后对不同句柄进行实例化。
	//
	var ioWriterDemo io.Writer = NewIOWriterDemo()
	ioWriterDemo.Write([]byte("this io.Write() op."))
	//....
	var x1, x2 strings.Builder
	w := io.MultiWriter(&x1, &x2)
	io.Copy(w, strings.NewReader("sdfadfadfad"))
	fmt.Println(x1.String(), x2.String())

	var y1 bytes.Buffer
	y1.String()

	var _ io.ReadWriter
	var _ io.ReadWriteCloser
	switch x := func() int { return 1 }(); x {
	case 1:
	}
	var xxx interface{} = 1
	switch xxx.(type) {
	case int:

	}
	switch v := xxx.(type) {
	case int:
		fmt.Println(v)
	}
	var xyz map[string]any = make(map[string]any)
	xyz["aaa"] = struct{}{}
	xyz["bbb"] = 1231
	for k, v := range xyz {
		fmt.Println(k, v)
	}
}

type MpType[KType comparable, VType any] map[KType]VType // comparable只能作为类型参数的约束，不能作为变量的类型。
