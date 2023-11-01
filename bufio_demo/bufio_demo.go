package bufio_demo

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。
//涉及到 io 读写时，下面的对象（带缓存的读写对象）包括：读写缓存 和 读写操作对象。后者是由外界实现io.读写接口的对象传递过来。

type BufioDemo struct {
	BfReader *bufio.Reader //buffering for an io.Reader object
	// bfreaderSz *bufio.Reader
	BfScaner *bufio.Scanner // 用于读取以一定分割符和字符组成的文件，字符串等。
	//
	BfWriter *bufio.Writer // buffering for an io.Writer object
}
type createBufioReader func(d *BufioDemo)

// 这个io reader 是从 string 里面去读。
func newStringReader(data string) createBufioReader {
	return func(d *BufioDemo) {
		// 这个io reader 是从 string 里面去读。
		d.BfReader = bufio.NewReader(strings.NewReader(data))
		//
		d.BfScaner = bufio.NewScanner(strings.NewReader(data))
	}
}

func newOsStdinReader() createBufioReader {
	return func(d *BufioDemo) {
		d.BfReader = bufio.NewReader(os.Stdin)
		d.BfScaner = bufio.NewScanner(os.Stdin)
	}
}
func newOutStdoutWriter() createBufioReader {
	return func(d *BufioDemo) {
		d.BfWriter = bufio.NewWriter(os.Stdout)
	}
}
func newOutTyteBuffer(b *bytes.Buffer) createBufioReader {
	return func(d *BufioDemo) {
		d.BfWriter = bufio.NewWriterSize(b, 10000)
	}
}

func newFileReader(f *os.File) createBufioReader {
	return func(d *BufioDemo) {
		d.BfReader = bufio.NewReader(f)
	}
	// ned to close f . f.Close()
}

func newFileWriter(f *os.File) createBufioReader {
	return func(d *BufioDemo) {
		d.BfWriter = bufio.NewWriter(f)
	}
}

//任何实现 Read([]byte)(n int, e error) 接口的对象都可以作为io.Reader。作为bufio.NewReader()的入参。

// scanner create
// bufio.Scanner 特别适合于读取 由换行符分割的文本行组成的 文件、
// 其中默认分割函数是将输入分割成行，并去掉行终止符。
func newSplitScan() createBufioReader {
	return func(d *BufioDemo) {
		d.BfScaner.Split(bufio.ScanWords)
		// c存在多种split的函数情况。
		// func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
		// func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
		// func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
		// func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	}
}

func NewBufioDemo(ways ...createBufioReader) *BufioDemo {
	ret := &BufioDemo{}
	for _, w := range ways {
		w(ret)
	}
	return ret
}

func BufioStringRun() {
	data := "this is string buf"
	strBuf := NewBufioDemo(newStringReader(data))
	//
	strBuf.scan_call()
	//
	strBuf.readByte()
	// strBuf.peek()
	strBuf.unreadbyte()

	//
	strBuf.readString()
	//
	strBuf.readLine()
	//
	strBuf.read()
	strBuf.readBytes()
	strBuf.buffered()
}

func BufioOsStdinRun() {
	stdinBuf := NewBufioDemo(newOsStdinReader())
	stdinBuf.scan_call()
	//
	stdinBuf.readLine()
	stdinBuf.readBytes()
}
func BufioFileRun() {
	fileName := "x.log"
	f, e := os.Open(fileName)
	if e != nil {
		return
	}
	fileBuf := NewBufioDemo(newFileReader(f))
	fileBuf.readString()
	f.Close()
	//
	f, e = os.OpenFile("x1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		fmt.Println("open x1.log fail, e: ", e)
		panic(e)
	}
	fileWriteBuf := NewBufioDemo(newFileWriter(f))
	fmt.Println("write to file: x1.log--------------------------")
	fileWriteBuf.write()       //
	fileWriteBuf.writeString() //
	f.Close()

}



func BufioOutStdoutRun() {
	stdoutBuf := NewBufioDemo(newOutStdoutWriter())
	stdoutBuf.write()
	stdoutBuf.writeString() //
}
func BufioOutByteBufferRun() {
	b := bytes.Buffer{}
	byteBuffer := NewBufioDemo(newOutTyteBuffer(&b))
	byteBuffer.write()
	//
	byteBuffer.writeString()
	fmt.Println("4) buf data: ", string(b.Bytes()))
}


func Bufio_demo_run() {
	fmt.Println(".... test 1 bufio read from string")
	BufioStringRun()
	fmt.Println("...... test 2 bufio read from os.Stdin")
	BufioOsStdinRun()
	fmt.Println("..... test 3 bufio read from file")
	BufioFileRun()

	fmt.Println("..... write test 1. stdout.....")
	BufioOutStdoutRun()
	fmt.Println("\n..... write 2 byte buffer.")
	BufioOutByteBufferRun()
}