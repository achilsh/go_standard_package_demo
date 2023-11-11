package templateDemo

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

//比如想生成指定类型的结构体代码：

//结构体中定义的 fieldName 是 template 中的 .Name； 将struct 中的field对应的value 作为模版的.Name 生成的源码。
//模板中的.Name 值及其操作， 实际上是struct 中 fieldName 对应的value和操作。

type ItemField struct {
	FieldName string
	FieldType string
}
type StructDataType struct {
	PkgName        string
	BusiStructType string
	FieldOneName   string
	FieldTwoName   string
	FieldThreeName string
	//
	MapKeyType   string
	MapValueType string
	FuncNameOne  string
	//
	MemberOneFunc    string
	MemberOneFuncRet string
	MemberTwoFunc    string
	MemberTwoFuncRet string

	Cache  bool
	Fields []ItemField
}

func runCreateStructType() {
	tplFileName := "struct_template.tpl"
	prefixPath, _ := filepath.Abs("./")

	tmpFile, e := os.Open(filepath.Join(prefixPath, "template", tplFileName))
	if e != nil {
		fmt.Println("open file fail, e: ", e)
		return
	}

	buf, e := io.ReadAll(tmpFile)
	if e != nil {
		fmt.Println("read from file fail, e: ", e)
		return
	}
	fmt.Println("read from file succ, ret len: ", len(buf))
	tmpHandle, e := template.New("create_struct_type").Parse(string(buf))
	// template.New("aaaa").ParseFiles("f1.tpl", "f2.tpl")
	if e != nil {
		fmt.Println("parse template fail, e: ", e)
		return
	}
	var structTypeData StructDataType = StructDataType{
		PkgName:        "templateDemo",
		BusiStructType: "StructCreateDemo",
		FieldOneName:   "Name",
		FieldTwoName:   "Age",
		FieldThreeName: "Score",
		MapKeyType:     "int32",
		MapValueType:   "*float32",
		FuncNameOne:    "DemoFuncFun",
		//
		MemberOneFunc:    "CallOne",
		MemberOneFuncRet: "int32",
		MemberTwoFunc:    "CallTwo",
		MemberTwoFuncRet: "float32",
		//
		Cache: true,
		Fields: []ItemField{
			{"ABC", "int32"},
			{"EFG", "float32"},
			{"HIG", "string"},
			{"OPQ", "bool"},
			{"RST", "[]byte"},
		},
	}

	// 将指定的数据传给已解析的模板上。
	e = tmpHandle.Execute(os.Stdout, &structTypeData)
	if e != nil {
		fmt.Println("call template execute fail, e: ", e)
		return
	}

	// var mpData map[int]string = make(map[int]string)
	// mpData[1]= "aaaaa"
	// template.New("create map key").Parse(string(buf))
}

func runStringType() {
	strFile, e := filepath.Abs(filepath.Join("template", "str.tpl"))
	if e != nil {
		fmt.Println("get absolute fail, e: ", e)
		return
	}
	tmpFile, e := os.Open(strFile)
	if e != nil {
		fmt.Println("open file fail, e: ", e)
		return
	}

	buf, e := io.ReadAll(tmpFile)
	if e != nil {
		fmt.Println("read from file fail, e: ", e)
		return
	}
	fmt.Println("read from file succ, ret len: ", len(buf))
	tmpHandle, e := template.New("create_struct_type").Parse(string(buf))
	if e != nil {
		fmt.Println("parse template fail, e: ", e)
		return
	}
	s := "aaaaa"
	e = tmpHandle.Execute(os.Stdout, &s)
	if e != nil {
		fmt.Println("execute fail, e: ", e)
		return
	}
}

func RunGenStructType() {
	runCreateStructType()
	runStringType()
}
