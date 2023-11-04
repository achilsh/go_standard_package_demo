package bytes_demo

import (
	"bytes"
	"fmt"
)

//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
//下面操作是对 bytes.Buffer： 可变大小，带有读写方法的 字节buf。字节buf对象 可以默认创建，可以通过其他 字节分片创建， 可以通过 字符串来创建。
type BytesBuffer struct {
	bf *bytes.Buffer // is a variable-sized buffer of bytes with Read and Write methods
	//可变大小，可读写 的字节buf, 。
}

func BytesBufferDemo() {
	fmt.Println(".... run bytes buffer demo....")
	b1 := NewBytesBuffer(newStrBuffer("this is demo"))
	b1.bytes()
	///
	b1.cap()
	b2 := NewBytesBuffer(newBytesBuffer())
	b2.cap()
	//
	b1.grow()
	b1.cap()
	//
	b2.grow()
	b2.cap()
	//
	b1.length()
	b2.length()
	//
	b1.next()
	fmt.Println("next test 2....")
	b2.next()
	//
	b1.length()
	///
	///
	b3 := NewBytesBuffer(newStrBuffer("this is demo"))
	b3.length()
	b3.read()
	b3.bytes()
	//
	b4 := NewBytesBuffer(newStrBuffer("aaaaaa\nbbbbb\ncccc"))
	b4.readbytes()
	b3.bytes()
	//
	fmt.Println("read from: to do")
	b5 := NewBytesBuffer(newBytesBuffer())
	b5.grow()
	b5.readfrom()
	fmt.Println("after read, bytes.Buffer unread data: ")
	b5.bytes()
	//
	b6 := NewBytesBuffer(newStrBuffer("cccc\ndddd\neeee\n"))
	b6.readstring()
	b6.bytes()
	b6.reset()
	b6.bytes()
	b6.length()
	//
	b7 :=  NewBytesBuffer(newStrBuffer("aaccbberfdfadfadfadf123123"))
	b7.str()
	//
	b8 := NewBytesBuffer(newBytesBuffer())
	b8.write() 
	b8.bytes()
	b8.length()
	//
	b9 := NewBytesBuffer(newStrBuffer("this is demo write to...."))
	b9.bytes()
	b9.writeto()
	b9.bytes()
	b9.length()
}

// //
type createBytesBuffer func(*BytesBuffer)

func NewBytesBuffer(c ...createBytesBuffer) *BytesBuffer {
	b := &BytesBuffer{}
	for _, cc := range c {
		cc(b)
	}
	return b
}

// bytes.Buffer 通过 其他分片创建。
func newBytesBuffer() createBytesBuffer {
	d := make([]byte, 0, 1024)
	return func(b *BytesBuffer) {
		b.bf = bytes.NewBuffer(d) //原来的分片 d 不能再使用，因为已经转移到 bytes.Buffer对象内了。
	}
}

// bytes.Buffer 通过 其他字符串创建。
func newStrBuffer(data string) createBytesBuffer {
	return func(b *BytesBuffer) {
		b.bf = bytes.NewBufferString(data) // creates and initializes a new Buffer using string s as its initial contents.
		//然后从这个 string上的数据进行读写。
	}
}
