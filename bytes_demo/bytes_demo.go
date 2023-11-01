package bytes_demo

import (
	"bytes"
	"fmt"
)

type BytesDemo struct {
	src []byte
	//the manipulation of byte slices. (字节分片). It is analogous to the facilities of the strings package
}

func NewByteDemo() *BytesDemo {
	r := &BytesDemo{}
	r.src = make([]byte, 0, 1024)
	r.src = append(r.src, []byte{1, 2, 3, 4, 6}...) //分片可以转化为可变参数， 加上语法糖 ... 即可。
	//
	return r
}

func (p *BytesDemo) copy() {
	d := bytes.Clone(p.src)
	fmt.Print(d)
	d[0] = 100
	fmt.Println(p.src)
	fmt.Print(d)
}
func (p *BytesDemo) compare() {
	a := []byte("abc")
	b :=[]byte("123abc")
	fmt.Println("\n a : b ", bytes.Compare(a, b))
}
func (p *BytesDemo) contains() {
	a := []byte("123abc")
	b := []byte("abc")
	fmt.Println("\n a contains b: ", bytes.Contains(a,b))
}
func (p *BytesDemo)mapCall() {
	s := []byte("abc")
	t := bytes.Map(func(r rune) rune{
			return r+ rune(1)
	}, s)
	fmt.Println(string(t))
}
func (p *BytesDemo) repeat() {
	s := []byte("a")
	t := bytes.Repeat(s,10)
	fmt.Println(string(t))
}
func (p *BytesDemo) split() {
	t := bytes.Split([]byte("a,b,c"), []byte(","))
	for _, v := range t {
		fmt.Println(string(v))
	}
}
func (p *BytesDemo) join() {
	s := [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")}
	t := bytes.Join(s,[]byte(",") )
	fmt.Println(string(t))
}
func (p *BytesDemo)title() {
	t := bytes.Title([]byte("this is demo, check it."))
	fmt.Println(string(t))
}
func BytesDemoRun() {
	fmt.Println("\n >>> run bytes demo begin <<<<<")
	b := NewByteDemo()
	b.copy()
	//
	b.compare()
	b.contains()

	b.mapCall()
	b.repeat()
	//
	b.split()
	b.join()
	//
	b.title()
	fmt.Println("\n >>> run bytes demo end <<<<<")

}
