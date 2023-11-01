package bytes_demo

import (
	"fmt"
	"io"
)

func (p *BytesBuffer) available() {
	if p == nil {
		return
	}
	//fmt.Println("available len: ", p.bf.Available())
}
func (p *BytesBuffer) bytes() {
	if p == nil {
		fmt.Println("is nil for bytes buffer.")
		return
	}
	fmt.Println("-----")
	fmt.Println("unread buf data: ", string(p.bf.Bytes())) //返回bytes.Bufferd的未读数据。

}

func (p *BytesBuffer) cap() {
	if p == nil || p.bf == nil {
		return
	}
	fmt.Println("byte.Buffer cap: ", p.bf.Cap()) //获取bytes.Buffer 中的容量。
}

func (p *BytesBuffer) grow() {
	if p == nil || p.bf == nil {
		return
	}

	p.bf.Grow(1000) // 增加bytes.Buffer的容量。
}

func (p *BytesBuffer) length() {
	if p == nil || p.bf == nil {
		return
	}
	fmt.Println("unread buf len: ", p.bf.Len()) //返回缓冲区未读部分的字节数;就是未读的数据的长度。
}

func (p *BytesBuffer) next() {
	n := 2
	if p == nil || p.bf == nil {
		return
	}

	//
	for {
		data := p.bf.Next(n) //返回byte.Buffer 中的下n个字节。同时更新已读的游标。类似已经发生读操作。
		if len(data) < n {
			fmt.Println("end to next buf data: ", string(data))
			return
		}
		fmt.Println("next buf data: ", string(data))
	}
}

func (p *BytesBuffer) read() {
	if p == nil || p.bf == nil {
		return
	}
	//
	dstData := make([]byte, 1024)
	n, e := p.bf.Read(dstData) //从bytes.Buffer 中读取数据并存到 dstData中。此时已读游标会增加。
	if e != nil {
		fmt.Println("read data from buf fail, e: ", e)
		return
	}
	fmt.Println("read from bytes.Buffer, data: ", string(dstData), ", len: ", n)
}

func (p *BytesBuffer) readbytes() {
	if p == nil || p.bf == nil {
		return
	}

	data, e := p.bf.ReadBytes(byte('\n')) //从bytes.Buffer中读取一序列字节，读取的结束标志位是指定的字节。
	if e != nil {
		fmt.Println("read bytes fail, e: ", e)
		return
	}
	fmt.Println("read bytes: ", string(data)) //如果读取到，包含结束字节。
}

type ReadFromSrc struct {
	src []byte
}

func (r *ReadFromSrc) Read(p []byte) (n int, err error) {
	if r == nil || r.src == nil || len(r.src) <= 0 {
		return 0, nil
	}
	return copy(p, r.src), io.EOF
}
func NewReadFromSrc() *ReadFromSrc {
	r := &ReadFromSrc{
		src: make([]byte, 1024),
	}
	copy(r.src, []byte("this is demo........"))
	return r
}
func (p *BytesBuffer) readfrom() {
	//从其他地方读取数据并存到 bytes.Buffer中。
	if p == nil || p.bf == nil {
		return
	}
	src := NewReadFromSrc()
	p.bf.ReadFrom(src)
}

func (p *BytesBuffer) readstring() {
	fmt.Println("read string from bytes.Buffer.")
	if p == nil || p.bf == nil {
		return
	}

	data, e := p.bf.ReadString(byte('\n')) //从bytes.Buffer中读取一序列字节，读取的结束标志位是指定的字节。只遍历一次。
	if e != nil {
		fmt.Println("read bytes fail, e: ", e)
		return
	}
	fmt.Println("read bytes: ", string(data)) //如果读取到，包含结束字节。
}

func (p *BytesBuffer) reset() {
	fmt.Println("reset from bytes.Buffer.")
	if p == nil || p.bf == nil {
		return
	}

	p.bf.Reset() //将内部buf空间清除，已读或者未读的位置 都置位 0.
}

func( p *BytesBuffer) str() {
	fmt.Println("string from bytes.Buffer.")
	if p == nil || p.bf == nil {
		return
	}
	s := p.bf.String()
	p.bf.Truncate(10) //把当前的缓存区，设置为已读和n个未读长度。或者说只保留n个未读字节。
	p.bf.UnreadByte() //将已读的位置向前-1，那么最后一个已读的字节变为了未读。

	fmt.Println("unread buf data: ", s) //未读的数据，以string 返回
}
