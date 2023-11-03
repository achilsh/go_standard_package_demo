package strings_demo

import (
	"fmt"
	"strings"
)

// strings是实现了 操作utf-8字符串 的函数
// strings是实现了 操作utf-8字符串 的函数
// strings是实现了 操作utf-8字符串 的函数
// strings是实现了 操作utf-8字符串 的函数
// strings是实现了 操作utf-8字符串 的函数
// strings是实现了 操作utf-8字符串 的函数

type StringsDemo struct {
}

func NewStringsDemo() *StringsDemo {
	r := &StringsDemo{}
	return r
}

func StringsDemoRun() {
	sd := NewStringsDemo()
	sd.clone()
	sd.compare()
	sd.contains()
	sd.containsAny()
	sd.containsRune()
	sd.count()
	sd.cut()
	sd.cutPrefix()
	sd.equalFlod()
	sd.fields()
	sd.fieldFunc()
	sd.hasPrefix()
	sd.hasSuffix()
	sd.index()
	sd.indexAny()
	sd.indexFunc()

}

func (s *StringsDemo) clone() {
	fmt.Println("run strings.Clone().")
	s1 := "this is demo clone"
	s2 := strings.Clone(s1) //将原来的字符串 新分配内存并给新内存copy 了字符串内容
	fmt.Println("origin str: ", s1, ", clone str: ", s2)
}
func (s *StringsDemo) compare() {
	fmt.Println("run strings.Compare()")
	s1 := "this test one"
	s2 := "this test two"
	fmt.Println("compare: ", strings.Compare(s1, s2)) //如果直接使用字符串 比较: ==, < ,> 效率会更高。
}

func (s *StringsDemo) contains() {
	fmt.Println("run strings.Contains()")
	s1 := "th is is demo"
	s2 := "demo"
	s3 := "ok"

	fmt.Println(s1, ", is contains, ", s2, ", : ", strings.Contains(s1, s2)) //检查字符串是否包含子串。
	fmt.Println(s1, ", is contains, ", s3, ", : ", strings.Contains(s1, s3)) //检查字符串是否包含子串。
	fmt.Println(s1, ", is contains: ", "", ", : ", strings.Contains(s1, "")) //检查字符串是否包含子串。任何字符串都包含空子串。
}

func (s *StringsDemo) containsAny() {
	fmt.Println("strings.ContainsAny")
	s1 := "this is dmeo"
	s2 := "abc"
	s3 := "ai"
	s4 := "sabc"
	s5 := ""
	fmt.Println(strings.ContainsAny(s1, s2)) //判断子串的任何字符是否包含在源字符串内。
	fmt.Println(strings.ContainsAny(s1, s3))
	fmt.Println(strings.ContainsAny(s1, s4))
	fmt.Println(strings.ContainsAny(s1, s5)) //其中子字符串为空时，那么内部就没有字符串，就不会包含在在源字符串内。
}

func (s *StringsDemo) containsRune() {
	fmt.Println("strings.ContainRune().")
	s1 := "this is demo"
	s2 := 'd'
	s3 := 'a'
	fmt.Println(strings.ContainsRune(s1, s2)) // 其中Rune 是一个4字节的码点。用 单引号标识，或者一个整形数字标识。
	fmt.Println(strings.ContainsRune(s1, s3))
}

func (s *StringsDemo) count() {
	fmt.Println("strings.Count().")
	s1 := "this is demo yes!, i am."
	s2 := "i"

	fmt.Println(strings.Count(s1, s2)) //在源字符串中统计子串出现个数。
	fmt.Println(strings.Count(s1, "")) //如果子串是空的，结果返回源字符串长度 + 1
}

func (s *StringsDemo) cut() {
	fmt.Println("strings.Cut().")
	s1 := "th is is demo, he is testing."
	s2 := "is"
	b, a, f := strings.Cut(s1, s2) //用子串去切分字符串，只在子串第一次出现的位置上切分，并返回切分后的 前部分，后部分，和 是否可以切分标识。
	fmt.Println("found before: ", b, ", found after: ", a, ", is found: ", f)
}

func (s *StringsDemo) cutPrefix() {
	fmt.Println("strings.CutPrefix()") //如果子串是字符串的最左边子串，那么就从源字符串中切掉
	s1 := "this is demo!"
	s2 := "this"
	s3 := "this ok"

	fmt.Println(strings.CutPrefix(s1, s2)) //s2是 s1的最左端 子串（前缀），可以被切除，返回剩下的子串
	fmt.Println(strings.CutPrefix(s1, s3)) //s3 不是 s1的最左端子串，不可以被切除。返回源字符串。
}

func (s *StringsDemo) cutSuffix() {
	fmt.Println("strings.CutSuffix()") //如果子串是源字符串的最右端子串（后缀），那么就从源串中切除子串。否则不切除。
	s1 := "this is demo, ok"
	s2 := "ok"
	s3 := "abc"
	fmt.Println(strings.CutSuffix(s1, s2))
	fmt.Println(strings.CutSuffix(s1, s3))
}

func (s *StringsDemo) equalFlod() {
	fmt.Println("strings.EqualFold()") //忽略大小写下比较两字符串是否相等。
	s1 := "this is demo"
	s2 := "This is Demo"
	s3 := "absAdfadf"
	fmt.Println(strings.EqualFold(s1, s2))
	fmt.Println(strings.EqualFold(s1, s3))
}

func (s *StringsDemo) fields() { //默认按空格来切分 字符串，返回的一些列 子串 分片。
	fmt.Println("strings.Fields()")
	s1 := "this    is    demo test!"

	//
	r := strings.Fields(s1) //默认情况下是按空格来切分字符串。返回是字符串分片。
	for _, v := range r {
		fmt.Printf("1:%v:1\n", v)
	}
}

func (s *StringsDemo) fieldFunc() { //主要是用于按某些字符来切分源字符串，得到不包含特定字符的字符串 分片。
	fmt.Println("strings.FieldsFunc()")
	s1 := "b11\nbb cc\ndd"

	r := strings.FieldsFunc(s1, func(c rune) bool {
		if c == '\n' {
			return true
		}
		return false
	}) // 对源字符串的每个字符，如果满足内部函数（内部函数返回为true），则从此字符开始切分源字符串，返回不包含该字符的子字符串。

	for _, v := range r {
		fmt.Printf("2:%v:2\n", v)
	}
}

func (s *StringsDemo) hasPrefix() {
	fmt.Println("strings.HasPrefix()")
	s1 := "this is demo"
	s2 := "this"
	s3 := "is"

	fmt.Println(strings.HasPrefix(s1, s2)) //查找子串是否是源字符串的最左子串。
	fmt.Println(strings.HasPrefix(s1, s3)) //查找子串是否是源字符串的最左子串。
}

func (s *StringsDemo) hasSuffix() {
	fmt.Println("strings.HasSuffix()")
	s1 := "this is demo"
	s2 := "demo"
	s3 := "is"
	fmt.Println(strings.HasSuffix(s1, s2)) //查找子串是否是 源字符的 最右子串。
	fmt.Println(strings.HasSuffix(s1, s3))
}

func (s *StringsDemo) index() {
	fmt.Println("strings.Index()")
	s1 := "this is demo"
	s2 := "demo"
	s3 := "ok"
	fmt.Println(strings.Index(s1, s2)) //查找子串第一次出现在源串中的位置。
	fmt.Println(strings.Index(s1, s3))
}

func (s *StringsDemo) indexAny() {
	fmt.Println("strings.IndexAny()")
	s1 := "this is demo"
	s2 := "ieo"
	fmt.Println(strings.IndexAny(s1, s2)) //返回 子串中任何一字符 第一次出现在源字符串中的位置。
}

func (s *StringsDemo) indexFunc() {
	fmt.Println("strings.IndexFunc()")
	s1 := "this is demo"
	f := func(r rune) bool {
		if r >= 'a' + 10  && r < 'a' + 15 { // rune 是一个字码，是4个字节大小的，如果字符标识，用单引号即可。
			return true
		}
		return false
	}
	fmt.Println(strings.IndexFunc(s1, f)) //查找源字符串中 第一次满足特定条件下的字符位置。
}