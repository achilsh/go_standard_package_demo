package bytes_demo

import (
	"bytes"
	"fmt"
	"io"
)

//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。
//下面操作是针对 bytes.Reader：创建对象时从byte slice获取数据，该对象实现 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner 方法从他对象读取数据。他只有读能力。

type BytesReader struct {
	br *bytes.Reader
	//它是只有 读能力，而且对象创建时数据源从 字节 分片里获取。支持seek查找。
}

func NewByteReader(d []byte) *BytesReader {
	if d == nil || len(d) <= 0 {
		panic("is nil")
		return nil
	}
	r := &BytesReader{
		br: bytes.NewReader(d),
	}
	return r
}
func BytesReaderRun() {
	d := []byte("this is bytes.Reader demot.")
	br := NewByteReader(d)
	br.length()
	//
	br.read()
	br.length()
	//

	br = NewByteReader(d)
	br.readat()
	br.length()
	//
	br.readbyte()
	br.length()

	//
	br.reset()
	br.length()
	br.read()
	//
	br.seek()
	br.size_call()

	//
	br.reset()
	br.write_to()

}

func (p *BytesReader) length() {
	fmt.Println("run bytes.Reader.Len")
	if p == nil || p.br == nil {
		return
	}
	fmt.Println("bytes.Reader len: ", p.br.Len())
}

func (p *BytesReader) read() {
	fmt.Println("run bytes.Reader.Read")
	if p == nil || p.br == nil {
		panic("is nil")
	}
	for {
		dst := make([]byte, 2)
		n, e := p.br.Read(dst) //实现了io.Reader 接口, 从底层的buf中读取数据。已经位置会增加。
		if e != nil && e == io.EOF {
			fmt.Println("has read all data")
			return
		}
		if e != nil {
			fmt.Println("read happen err: ", e)
			panic(e)
		}
		//
		fmt.Println("read from data: ", string(dst), ", ret len: ", n)
	}
}

func (p *BytesReader) readat() {
	fmt.Println("run bytes.Reader.ReadAt.")
	if p == nil || p.br == nil {
		panic("is nil")
	}

	off := int64(0)
	for {
		d := make([]byte, 2)
		n, e := p.br.ReadAt(d, off) //按指定偏移位置去读数据。但是已读的位置保持不变
		if e != nil && e == io.EOF {
			fmt.Println("read all data done.")
			return
		}
		if e != nil {
			fmt.Println("read happen e: ", e)
			panic(e)
		}
		off += int64(n)
		fmt.Println("read data: ", string(d))
	}
}

func (p *BytesReader) readbyte() {
	fmt.Println("run bytes.Reader ReadByte .")
	if p == nil || p.br == nil {
		panic("is nil")
	}

	for {
		r, e := p.br.ReadByte() //从底层buf中读取一个字节，同时已读位置增加。
		if e != nil && e == io.EOF {
			fmt.Println("read done.")
			return
		}
		if e != nil {
			fmt.Println("read e: ", e)
			return
		}

		fmt.Println("readbyte ret: ", string(r))
	}
}

func (p *BytesReader) reset() {
	fmt.Println("run bytes.Reader Reset .")
	if p == nil || p.br == nil {
		panic("is nil")
	}
	//

	p.br.Reset([]byte("reset data")) //将底层的buf 重新置位 新的内容，已读位置置空。重新开始。
}

func (p *BytesReader) seek() {
	fmt.Println("run bytes.Reader Seek .")
	if p == nil || p.br == nil {
		panic("is nil")
	}

	n, e := p.br.Seek(2, io.SeekCurrent) //根据入参类型，返回底层存储位置+最新offset后的绝对位置。底层已读位置更新最新值。
	if e != nil {
		fmt.Println("seek fail, e: ", e)
		return 
	}
	fmt.Println("position: ", n)
}

func (p *BytesReader) size_call() {
	fmt.Println("run bytes.Reader Size .")
	if p == nil || p.br == nil {
		panic("is nil")
	}
	fmt.Println("buf size: ", p.br.Size()) //返回值不受除 Reset()函数的影响。

	p.br.UnreadByte() //将已读的位置向前减一。就是将最新已读的字节变为未读字节。
}

type OutIOWriter struct {
	src []byte
}
func NewOoutIoWriter() *OutIOWriter{
	r := &OutIOWriter{}
	r.src = make([]byte, 1024)
	return r
}
func (o *OutIOWriter) Write(p []byte) (n int, err error) {
	if p == nil || len(p) <= 0 {
		return 0, nil
	}
	return copy(o.src, p),nil 
}
func (p *BytesReader) write_to() {
	fmt.Println("run bytes.Reader WriteTo .")
	if p == nil || p.br == nil {
		panic("is nil")
	}
	//直接意思就是往当前 Bytes.Reader中写数据。
	w := NewOoutIoWriter()
	n, e := p.br.WriteTo(w) //将当前p.br中的buf内容写到 入参对象内，内部调用入参的方法Write(); 当前p.br对象底层已读位置增加。
	if e != nil {
		fmt.Println("write to fail, e: ", e)
		return 
	}
	fmt.Println("write succ data len: ", n)
	
	fmt.Println("dst write data: ", string(w.src))

}
