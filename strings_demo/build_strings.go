package strings_demo

import (
	"fmt"
	"strings"
)

//这是 string的构造器， 使用构造器的write()方法来构造 string.
//这是 string的构造器， 使用构造器的write()方法来构造 string.
//这是 string的构造器， 使用构造器的write()方法来构造 string.
//这是 string的构造器， 使用构造器的write()方法来构造 string.
//这是 string的构造器， 使用构造器的write()方法来构造 string.
//这是 string的构造器， 使用构造器的write()方法来构造 string.
type BuildStringsDemo struct {
	strDataBuild strings.Builder //这是 string的构造器， 使用构造器的write()方法来构造 string.
}

func NewBuildStringsDemo() *BuildStringsDemo {
	return &BuildStringsDemo{}
}
/////

func RunStringBuilderDemo() {
	r := NewBuildStringsDemo()
	r.cap_call()

	r.grow()
	r.cap_call()

	r.length()

	r.cap_call()
	r.reset()
	r.cap_call()

	r.writeBytes()
	r.length()
	r.cap_call()
	r.string_call()

	r.reset()
	r.string_call()

	r.write_string()
	r.string_call()
}

func (p *BuildStringsDemo) cap_call() {
	fmt.Println("call strings.Builder.Cap()")
	fmt.Println(p.strDataBuild.Cap())
}
func (p *BuildStringsDemo) grow() {
	fmt.Println("call string.Builder.Grow()")
	p.strDataBuild.Grow(100)
	fmt.Println(p.strDataBuild.Len())
}
func (p *BuildStringsDemo) length() {
	fmt.Println("strings.Builder.Len()")
	fmt.Println(p.strDataBuild.Len())
}
func (p* BuildStringsDemo) reset() {
	fmt.Println("strings.Builder.Reset()")
	p.strDataBuild.Reset()
	//
}

func (p *BuildStringsDemo) writeBytes() {
	fmt.Println("strings.Builder.Write()")
	n, e := p.strDataBuild.Write([]byte("this is demo"))
	if e != nil {
		fmt.Println("write data to strings.Builder fail, e: ", e)
		return 
	}
	fmt.Println("strings.Builder.Write() succ, ret n: ", n)
}

func (p *BuildStringsDemo) string_call() {
	fmt.Println("run strings.Builder.String()")
	data := p.strDataBuild.String()
	fmt.Println("Builder string: ", data)
}

func (p *BuildStringsDemo) write_string() {
	fmt.Println("strings.Builder.WriteString()")
	p.strDataBuild.WriteString("\n this is again write.")
}