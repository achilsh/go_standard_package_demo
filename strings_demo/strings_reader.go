package strings_demo

import (
	"fmt"
	"strings"
)

// strings.Reader 是一个包装 string 的对象， 该对象实现了 io.Reader, io.ReaderAt, io.ByteReader,
// io.ByteScanner, io.RuneReader, io.RuneScanner, io.Seeker, and io.WriterTo 接口(从 对象上 读取数据，并更新已读位置)，
// 并且该对象 只有读方法。他其实是一个 从string 读取数据的对象，同时读的方式更加标准化一下。已经满足一些通用接口。
// 该对象可以 创建默认对象； 也可以通过： 一个字符串作为入参创建，
// 使用这个主要想从 string 中 通过 Reader() 接口获取数据。


type StringsReaderDemo struct {
	defstrReader *strings.Reader 
	fromStrStrReader *strings.Reader
}

type createStrReader func(*StringsReaderDemo)
func NewStringReaderDemo(opts ...createStrReader) *StringsReaderDemo {
	r := new(StringsReaderDemo)
	for _, o := range opts {
		o(r)
	}
	return r
}

func newDefStrReader() createStrReader {
	return func(d *StringsReaderDemo) {
		d.defstrReader = new(strings.Reader)
	}
}
func newStringStrReader(s string) createStrReader {
	return func(d *StringsReaderDemo) {
		d.fromStrStrReader = strings.NewReader(s)
	}
}

func RunStringsReader() {
	r := NewStringReaderDemo(newDefStrReader(), newStringStrReader("aaaa12"))
	r.length()

	r.read()
	r.length()

	r = NewStringReaderDemo(newDefStrReader(), newStringStrReader("aaaa12"))
	r.length()
	r.write_to()
	r.length()

}

func (s *StringsReaderDemo) length() {
	fmt.Println("run strings.Reader.Len()")
	if s == nil {
		return 
	}
	if s.defstrReader != nil {
		fmt.Println("default strings.Reader len: ", s.defstrReader.Len())
	}
	if s.fromStrStrReader != nil {
		fmt.Println("from str strings.Reader len: ", s.fromStrStrReader.Len())
	}
}

func(s *StringsReaderDemo) read() { //通过 标准的read()接口， 从strings.Reader对象中读取数据。读完之后，Reader内部的已读位置增加。
	fmt.Println("strings.Reader.Read()")
	if s == nil {
		return 
	}
	//...
	if s.defstrReader != nil {
		dst := make([]byte, s.defstrReader.Len())

		n, e := s.defstrReader.Read(dst) //
		if e != nil {
			fmt.Println("read from default fail, e: ", e)
		} else {
			fmt.Println("read from default Reader succ, data: ", dst, ", len: ", n)
		}
	}
		


	if s.fromStrStrReader != nil {
		dst := make([]byte, s.fromStrStrReader.Len())
		n, e := s.fromStrStrReader.Read(dst)
		if e != nil {
			fmt.Println("read from fail, e : ", e) 
		} else {
			fmt.Println("read from reader succ, data: ", string(dst), ", len: ", n)
		}
	}

	xy := []any {"aa", false, 123, 12.123}
	fmt.Println(xy)
	for _, x := range xy {
		fmt.Println(x)
	}
}

type ff func(data []byte)(n int, err error)
func (f *ff) Write(data []byte)(n int, err error) {
	dst :=make([]byte, 1000)
	n =copy(dst, data)
	fmt.Println("dst: ", string(dst))
	return n, nil
}

func (s *StringsReaderDemo) write_to() {
	fmt.Println("run strings.Reader.WriteTo()") //将strings.Reader对象往实现了 write()接口的对象里写数据。 
	if s == nil {
		return 
	}
	//(p []byte) (n int, err error)
	var f *ff = new(ff)
	if s.defstrReader != nil {
		fmt.Println("w def string reader.")
		s.defstrReader.WriteTo(f)
	}
	if s.fromStrStrReader != nil {
		fmt.Println("w str string reader.")
		s.fromStrStrReader.WriteTo(f)
	}
}