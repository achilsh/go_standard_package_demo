package templateDemo

import (
	"fmt"
	"os"
	"text/template"
)

//该包主要是根据一些模板配置文件，生成 go 代码。比较适合于低代码的组件和结构的编写。
//实现 了 用于生成文本输出的 数据驱动模板 。 就是把模板应用于 数据结构。
//模板中的注释 引用 数据结构中的元素（struct的字段或者map中的key) 是为了控制执行并派生要显示的值。
//当执行模板时（调用模板的Execute接口时，）游标（其实就是模板的.）值被 设置 为结构体当前位置的值。
// actions 标识的是数据的评估或者控制结构体， 他们是被 {{}} 来分割的，
// {{  }} 外面的字符不会被处理而是被直接copy到输出。

func TemplateDemoRun() {
	{
		// demo1:
		//  {{.}} 中间的点，代表传入模板的数据。可以代表 go 语言中的任何类型，如结构体、哈希, 字符串，等
		//  传给模板的数据可以通过点（.）来访问。如果是复合类型的数据，则可以通过{{.FieldName}}来访问它的字段 FieldName

		//其中 {{ 后面紧跟-和空格，那么在生成模版源码时，会将-和空格都去掉。同理 }} 前面紧贴空格和-，那么生成模版源码时也会吧-和空格去掉。

		//{{ 这里是可打印的字符符串 }} 内加上一些可打印的字符串， 那么在生成源码中是存在这些字符串的。

		// 在 {{ }} 内部加注释， 那么在生成的源码中是没有这些注释的，注意注释书写格式，参考下面。
		//下面是实例： 

		demo1Template := `{{12   -}} < {{-     45}} {{ "文本字符串，在生成的源是有的。" }} 
		{{/** 模版中注释要紧挨着左右大括号， 生成源码时不会被输出到源码中 ***/}} {{- /** 模版中的注释和- 保持有空格就可。减掉空格，生成的源码将不会有这些注释  **/ -}}
		 大括号外面的注释 在生成的源码保持存在。 
		 下面 if pipelie 为空， 不执行 fmt.Println()
		{{if false}} fmt.Println("is false, not call, Dot is unaffected") {{end}}
		{{if 0}} fmt.Println("is 0, not call, Dot is unaffected) {{end}}
		{{if nil}} fmt.Println("is nil, not call, Dot is unaffected") {{end}}
		{{if ""}} fmt.Println("is empty str, not call, Dot is unaffected") {{end}}

		下面if pipeline 不为空，执行 fmt.Println("not empty, call it"), 否则执行  fmt.Println("is empty. call")
		{{if 1}} fmt.Println("is true, call(), Dot is unaffected ")  {{else}}  fmt.Println("is empty, call(), Dot is unaffected") {{end}}
		`


		t1, e := template.New("demo1").Parse(demo1Template)
		if e != nil {
			fmt.Println("parse fail, e : ", e)
		} else {
			t1.Execute(os.Stdout, nil) //输出 12<45
		}

		//{{ pipeline }} 中： The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero
	}

	{
		//{{ 和 }} 包裹的内容统称为 action，分为两种类型：
		//数据求值（data evaluations）
		//控制结构（control structures）
	}

}
