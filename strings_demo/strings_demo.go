package strings_demo

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	sd.indexLast()
	sd.lastIndexFunc()
	sd.join()
	sd.mapCall()
	sd.repeat()
	sd.replace()
	sd.replaceAll()
	sd.split()
	sd.splitN()

	sd.splitAfter()
	sd.splitAfterN()

	sd.title()
	sd.toTitle()
	sd.toLower()
	sd.toUpper()

	sd.trimCall()
	sd.trimFunc()

	sd.trimPrefix()
	sd.trimSuffix()
	sd.trimSpace()
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

func (s *StringsDemo) indexlast() {

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
		if r >= 'a'+10 && r < 'a'+15 { // rune 是一个字码，是4个字节大小的，如果字符标识，用单引号即可。
			return true
		}
		return false
	}
	fmt.Println(strings.IndexFunc(s1, f)) //查找源字符串中 第一次满足特定条件下的字符位置。
}

func (s *StringsDemo) indexLast() {
	fmt.Println("run strings.LastIndex()")
	s1 := "this is demo ok"
	s2 := "demo"

	fmt.Println(strings.LastIndex(s1, s2)) //在子串在源字符串 最后出现的位置。
}
func (s *StringsDemo) lastIndexFunc() {
	fmt.Println("run strings.LastIndexFunc()")
	s1 := "abc 123abc"
	fmt.Println(strings.LastIndexFunc(s1, func(d rune) bool { //判断字符串中字符从后往前，第一次满足函数条件时，字符位置。
		if d == 'a' {
			return true
		}
		return false
	}))
}

func (s *StringsDemo) join() {
	fmt.Println("run strings.Join()")
	x := []string{"aaa", "bbb", "ccc"}
	t := strings.Join(x, "|") // 将字符串分片用 特定的字符串分割符 来拼接。
	fmt.Println(t)
}

func (s *StringsDemo) mapCall() {
	fmt.Println("run strings.Map()")
	//对字符串副本中每个字符使用统一处理函数。
	s1 := "abcefa"
	s2 := strings.Map(func(d rune) rune {
		return d + 1
	}, s1)
	fmt.Println(s2)
}

func (s *StringsDemo) repeat() {
	fmt.Println("run strings.Repeat()")
	s1 := "is important! "

	fmt.Println(strings.Repeat(s1, 3))
}

func (s *StringsDemo) replace() {
	fmt.Println("run strings.Replace()")
	s1 := "no no no, can no do!"
	old := "no"
	newStr := "ok"

	ret := strings.Replace(s1, old, newStr, 0) //对某个字符串， 使用新的子串替代 内部子串 n 次。并返回替换后的新字符串。 如果替换次数 < 0, 那就全部替换字符串的子串。
	fmt.Println(ret)
}

func (s *StringsDemo) replaceAll() {
	fmt.Println("run strings.ReplaceAll()")
	s1 := "ok, ok, ok, ye..."
	old := "ok"
	newStr := "no"

	ret := strings.ReplaceAll(s1, old, newStr)
	fmt.Println(ret)
}

func (s *StringsDemo) split() {
	fmt.Println("run strings.Split()")
	s1 := "a,b,c,d , e"
	r1 := strings.Split(s1, ",") //按分割符将字符串进行分割，分割后的字符串列表 不包含分隔符。
	fmt.Printf("%q\n", r1)
	r2 := strings.Split(s1, "") //如果分隔符是空字符串，那么分割就是将源字符串按每个字符进行分割。返回字符串列表。
	fmt.Printf("%q\n", r2)
}
func (s *StringsDemo) splitN() {
	fmt.Println("strings.SplitN()")
	s1 := "a,b,c,d, e,  f"
	r1 := strings.SplitN(s1, ",", 2) //指定返回的子串个数。 n == 1， 返回列表长度为1， 内容为源字符串，
	// n > 1 时，返回子串个数为n, 返回 列表中最后一个子串是： 源串被分割成 n-1端后的最后 一段。
	fmt.Printf("%q\n", r1)

}
func (s *StringsDemo) splitAfter() {
	fmt.Println("run string.SplitAfter()")
	s1 := "a, b,c  , d"
	r1 := strings.SplitAfter(s1, ",") //使用分割符对字符串进行划分，每个子串都包含分隔符， 返回子串列表，
	fmt.Printf("%q\n", r1)
}

func (s *StringsDemo) splitAfterN() {
	fmt.Println("run strings.SplitAfterN()")
	s1 := "a,b,c, d"
	//分割符包含在 返回的子串中。
	r1 := strings.SplitAfterN(s1, ",", 5) //其中n 指定返回的子串列表中的子串个数。如果是1，返回列表长度为1，内容为源字符串的列表。
	// 如果n 为 > 1  表示， 返回多个子串，其中最后一个子串为： 源字符串被分割n-1次后的最后一段。
	fmt.Printf("%q\n", r1)
}

func (s *StringsDemo) title() {
	fmt.Println("run strings.Title()") //是将字符串的首字母变为大写
	s1 := "this is demo, check it ok!"
	//golang.org/x/text/cases
	//cases.Title(language.English).String(s1)
	r1 := strings.Title(s1)
	fmt.Println(r1)
	fmt.Println(cases.Title(language.English).String(s1)) ////是将字符串中 单词的首字母变为大写
}

func (s *StringsDemo) toTitle() {
	fmt.Println("run strings.ToTitle()") //将字符串中的所有的 字码编程 大写字码。
	t1 := "this is demo, check it"
	fmt.Printf("%q\n", strings.ToTitle(t1))
}

func (s *StringsDemo) toLower() {
	fmt.Println("run strings.ToLower()") //将字符串 所有字母 转化为 小写。
	s1 := "This Is DemoOK, YYYYYEEEEE"
	r1 := strings.ToLower(s1)
	fmt.Printf("%q\n", r1)
}

func (s *StringsDemo) toUpper() {
	fmt.Println("run strings.ToUpper()") //将字符中 所有字母 转为大写
	s1 := "this is demo,, yee!"
	r1 := strings.ToUpper(s1)
	fmt.Println(r1)
}

func (s *StringsDemo) trimCall() {
	fmt.Println("run strings.Trim()")
	s1 := "a!! and ! no body 99"
	r1 := strings.Trim(s1, "!9") //从字符串的最左边和最右边 分别切 出现在 切除字符集中 的字符。切掉一个后再继续判断是否满足条件，直到不满足为止。如果最左边、最右字符不存在切除字符集则停止迭代。
	fmt.Printf("%q\n", r1)
}

func (s *StringsDemo) trimFunc() {
	fmt.Println("Run string.TrimFunc()")
	s1 := "12abcd2321"
	r1 := strings.TrimFunc(s1, func(d rune) bool { //分别从字符串的最左边和最右边，切除满足函数条件（返回值为true)的字符。切除一个再继续判断是否满足，直到不满足条件位置。
		if d >= '0' && d <= '9' {
			return true
		}
		return false
	})
	fmt.Printf("%q\n", r1)
}

//有切除最左边的字符，这些字符是在 切除字符集内。
//有切除最右边的字符，切除的字符是在 切除字符集内。

// ...
func (s *StringsDemo) trimPrefix() {
	fmt.Println("run strings.TrimPrefix()") //从字符串最左边匹配 切除子串，匹配则切除最左边的子串。
	s1 := "this is demo, ok yes"
	r1 := strings.TrimPrefix(s1, "this i")
	fmt.Printf("%q\n", r1)
}

func (s *StringsDemo) trimSuffix() {
	fmt.Println("run strings.TrimSuffix()") //从字符串最右边匹配 待切除的子串，匹配则切除最右匹配的子串
	s1 := "this is demo,yyy!"
	r1 := strings.TrimSuffix(s1, "demo,yyy!")
	fmt.Printf("%q\n", r1)
}

func (s *StringsDemo) trimSpace() {
	fmt.Println("run strings.TrimSpace") // 分别从最左边和最右边的字符上切除空格，如果能被切除则继续切除； 直到不是空格字符 结束切除。
	s1 := "    s1ssf   idfadf        "
	r1 := strings.TrimSpace(s1)
	fmt.Printf("%q\n", r1)
}
