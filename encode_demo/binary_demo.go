package encode_demo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

//主要是是一些二进制编码方式的处理，例如网络上的大小端问题等。
//字节序列操作的接口：
// type ByteOrder interface {
// 	Uint16([]byte) uint16 //获取 2个字节
// 	Uint32([]byte) uint32 //获取 4个字节
// 	Uint64([]byte) uint64  // 获取8个字节
// 	PutUint16([]byte, uint16) // 写入2个字节
// 	PutUint32([]byte, uint32) // 写入4个字节
// 	PutUint64([]byte, uint64) // 写入8个字节
// 	String() string
// }
// 实现了 ByteOrder 接口的对象有: binary.LittleEndian binary.BigEndian
// 下面练习用 binary.BigEndian 来写demo.

type BinaryCodeDemo struct {
	BigEndianCode    binary.ByteOrder
	LittleEndianCode binary.ByteOrder
	//
	data  []byte
	index int32 //已读位置
	//
	rwBuf      bytes.Buffer      //可以用bufio，其中io 可以设置网络连接对象，或者其他文件的句柄啥的。
	otherRwBuf *bufio.ReadWriter // 这是定义 由外界输入的io.writer/io.reader的 buf对象。
}

func NewBinCodeDemo() *BinaryCodeDemo {
	r := &BinaryCodeDemo{
		BigEndianCode: &binary.BigEndian,
		// LittleEndianCode: &binary.LittleEndian,
		data:  make([]byte, 1024),
		index: 0,
		//
		otherRwBuf: bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)),
	}
	return r
}
func RunBinaryDemo() {
	demo := NewBinCodeDemo()
	demo.Set2byte()
	demo.Get2byte()
	//
	demo.Set4byte()
	demo.Get4byte()
	//
	demo.Set8byte()
	demo.Get8byte()
	//
	demo.write()
	demo.read()
}

func (p *BinaryCodeDemo) Set2byte() {
	x := uint16(123)
	p.BigEndianCode.PutUint16(p.data, x)
	p.index += 2
	fmt.Println("set uint16: ", x)

}
func (p *BinaryCodeDemo) Set4byte() {
	y := uint32(4567)
	p.BigEndianCode.PutUint32(p.data[p.index:], y)
	p.index += 4
	fmt.Println("set uint32: ", y)

}
func (p *BinaryCodeDemo) Set8byte() {
	z := uint64(89101112)
	p.BigEndianCode.PutUint64(p.data[p.index:], z)
	p.index += 8
	fmt.Println("set uint64: ", z)
}

func (p *BinaryCodeDemo) Get2byte() uint16 {
	x := p.BigEndianCode.Uint16(p.data[:2])
	fmt.Println("get uint16 data: ", x)
	return x
}
func (p *BinaryCodeDemo) Get4byte() uint32 {
	x := p.BigEndianCode.Uint32(p.data[2:6])
	fmt.Println("get uint32 data: ", x)
	return x
}
func (p *BinaryCodeDemo) Get8byte() uint64 {
	x := p.BigEndianCode.Uint64(p.data[6:14])
	fmt.Println("get uint64 data: ", x)
	return x
}

func (p *BinaryCodeDemo) write() {
	//io.Writer 可以是网络连接，这样就可以往连接里写数据。主要是以大端来验证
	//对 data的描述：数据必须是固定大小的值或固定大小的值的切片，或者是指向此类数据的指针。
	// 布尔值编码为一个字节:1表示真，0表示假。写入w的字节使用指定的字节顺序进行编码，
	// 并从数据的连续字段中读取。在编写结构体时，为字段名为空(_)的字段写入零值。
	//

	fmt.Println(" binary code write run.")
	toWriteData := []byte("this write and read demo.")
	e := binary.Write(&p.rwBuf, p.BigEndianCode, toWriteData)
	if e != nil {
		fmt.Println("write data fail, e: ", e, ", to write data: ", string(toWriteData))
	}
	fmt.Println("write to data succ: ", string(toWriteData))
	//
	fmt.Println("----> write data to stdout: ", string(toWriteData))
	binary.Write(p.otherRwBuf, p.BigEndianCode, toWriteData)
}

func (p *BinaryCodeDemo) read() {
	//io.Reader 可以是网络连接，这样就可以从网络连接上读取数据。主要是用大端来验证。
	//data的描述：Data必须是指向固定大小值的指针或固定大小值的切片。
	// 从r读取的字节使用指定的字节顺序解码，并写入数据的连续字段。
	// 在解码布尔值时，零字节将被解码为false，而任何其他非零字节将被解码为true。
	// 当读入结构体时，字段名为空(_)的字段数据将被跳过;也就是说，空白字段名可以用于填充。
	// 当读入结构体时，必须导出所有非空白字段，否则Read可能会panic。

	fmt.Println("binary code read demo run")
	toReadData := make([]byte, p.rwBuf.Len())
	e := binary.Read(&p.rwBuf, p.BigEndianCode, &toReadData)
	if e != nil {
		fmt.Println("read from binary code demo fail, e: ", e)
		return
	}
	fmt.Println("read from binary code succ, data: ", string(toReadData))

	fmt.Println("--->>> begin to read form stdin data, with 10 nums data:")
	toReadBuf := make([]byte, 10)
	binary.Read(p.otherRwBuf, p.BigEndianCode, toReadBuf)
	fmt.Println(" to read data: ", string(toReadBuf))
}
