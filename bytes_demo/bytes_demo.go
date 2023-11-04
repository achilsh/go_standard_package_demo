package bytes_demo

import (
	"bytes"
	"fmt"
)

//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
//下面方法是操作 ： 定义了一些对 字节分片 的操作
type BytesDemo struct {
	src []byte //主要 对字节分片操作。就是普通的基本结构的分片，需要业务自己去定义。std pkg 只提供操作的方法。
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
	b := []byte("123abc")
	fmt.Println("\n a : b ", bytes.Compare(a, b))
}
func (p *BytesDemo) contains() {
	a := []byte("123abc")
	b := []byte("abc")
	fmt.Println("\n a contains b: ", bytes.Contains(a, b))
}
func (p *BytesDemo) mapCall() {
	//对[]byte中每个rune 进行运行。并返回最终结果。
	s := []byte("abc")
	t := bytes.Map(func(r rune) rune {
		return r + rune(1)
	}, s)
	fmt.Println(string(t))
}

func (p *BytesDemo) repeat() {
	s := []byte("a")
	t := bytes.Repeat(s, 10)
	fmt.Println(string(t))
}
func (p *BytesDemo) split() {
	//通过，把[]byte 拆成 []byte 的 slice.
	t := bytes.Split([]byte("a,b,c"), []byte(","))
	for _, v := range t {
		fmt.Println(string(v))
	}
}
func (p *BytesDemo) join() {
	//通过，把所有的byte slice 连接程一个新的[]byte
	s := [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")}
	t := bytes.Join(s, []byte(","))
	fmt.Println(string(t))
}
func (p *BytesDemo) title() {
	//将word 的开始字母变成大写。
	t := bytes.Title([]byte("this is demo, check it."))
	fmt.Println(string(t))
}
func (p *BytesDemo) tolower() {
	//将byte的所有大写编程小写。
	t := bytes.ToLower([]byte("I am STUDENT. But how for you?"))
	fmt.Println(string(t))
	//将byte的所有变成大写.
	fmt.Println(string(bytes.ToUpper([]byte("i am student, and you ?"))))
	//对[]byte中每个字符，转化为标题大写的。
	fmt.Println(string(bytes.ToTitle([]byte("This is Demo, chack it?"))))
}

func (p *BytesDemo) trim() {
	//必须从头开始匹配，或者从尾部开始匹配字符串，如果字符串在指定内。那么就去除。如果发现没有，就中断匹配。
	//去除也是从头开始，或者从尾部开始去除。
	fmt.Println(string(bytes.Trim([]byte("!a  !abc d!"), "!")))
	//通过切掉cutset中包含的所有前导和尾随utf -8编码的代码点，返回s的子切片
	fmt.Println(string(bytes.Trim([]byte("!!! Achtung! Achtung! !!! "), "! ")))

	//从左边匹配字符串是否在集合集合内，如果存在则去除。如果不匹配就终端匹配。
	fmt.Println(string(bytes.TrimLeft([]byte("123abads2323"), "0123456789")))

	//对字符串， 从左开始匹配对应字符串，完全匹配则删除。不匹配则中断匹配。
	fmt.Println(string(bytes.TrimPrefix([]byte("aa bb cc"), []byte("a"))))
	///对字符串， 从右边开始匹配对应字符串，完全匹配则删除。不匹配则中断匹配。
	fmt.Println(string(bytes.TrimSuffix([]byte("aa bb cc"), []byte("cc"))))

	//对字符串，分别从左和从右开始 过滤空格或者\t。匹配继续过滤，不匹配就终端匹配。就是：去掉字符串前后的空。
	fmt.Println(string(bytes.TrimSpace([]byte("    \t\tb  aa \tcc\t\t\t       "))), "aa")
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
	b.tolower()

	//
	b.trim()
	fmt.Println("\n >>> run bytes demo end <<<<<")

}
