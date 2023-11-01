package bufio_demo

import "fmt"

func (p *BufioDemo) buffered() {
	if p == nil || p.BfReader == nil {
		fmt.Println("buf is nil")
		return
	}
	fmt.Printf("buffered size: %d\n", p.BfReader.Buffered())

}
func (p *BufioDemo) discard() {

}
func (p *BufioDemo) peek() {
	//返回下n个字节，但是不挪动buf中未读的位置。类似stack中的top()接口那样。
	data, e := p.BfReader.Peek(1)
	fmt.Println("peek ret data: ", string(data), ", e: ", e)

}
func (p *BufioDemo) readByte() {
	data, e := p.BfReader.ReadByte()
	fmt.Println("read byte: ", string(data), ", e: ", e)

}
func (p *BufioDemo) readBytes() {
	data, e := p.BfReader.ReadBytes(byte('\n')) // 's'
	if e != nil {
		fmt.Println("read bytes fail, e: ", e, ", data: ", string(data))
		return
	}
	fmt.Println("read bytes, ret data: ", string(data))

}
func (p *BufioDemo) read() {
	data := make([]byte, 1000)
	n, e := p.BfReader.Read(data)
	if e != nil {
		fmt.Println("read data fail, e: ", e)
		return
	}
	fmt.Println("data ret len: ", n)
	fmt.Println("data: ", string(data))
}
func (p *BufioDemo) readLine() {
	var lda, prefix, e = p.BfReader.ReadLine()
	if e != nil {
		//
	}
	fmt.Println("read line fail, e: ", e, ", data: ", string(lda), ", prefix: ", prefix)

}
func (p *BufioDemo) readRun() {

}
func (p *BufioDemo) readSlice() {

}
func (p *BufioDemo) readString() {
	data, e := p.BfReader.ReadString(byte('\n'))
	fmt.Println("read string, data: ", data, ", e: ", e)

}
func (p *BufioDemo) reset() {

}
func (p *BufioDemo) size() {

}
func (p *BufioDemo) unreadbyte() {
	//就是将已经的读过的前一个byte，作为未读字段，等到下一次读时可以被读出来。
	e := p.BfReader.UnreadByte()
	fmt.Println("unread byte e: ", e)

}
func (p *BufioDemo) unreadrun() {

}

// 定义 scanner 的方法
func (p *BufioDemo) scan_call() {
	i := 5
	for p.BfScaner.Scan() {
		i--
		fmt.Println("scan data: ", p.BfScaner.Text())
		if i <= 0 {
			fmt.Println("stop by manual scan")
			break
		}
	}
	if p.BfScaner.Err() != nil {
		fmt.Println("scan happen err: ", p.BfScaner.Err().Error())
	}
}