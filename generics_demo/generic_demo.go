package genericdemo

import "fmt"

//本章节主要 介绍 go的模板使用。

//类型参数的限制：用interface 表示。
type ValueConstraint interface { 
	// 申明模板类型限制的类型。这样在写模板时模板的形参类型可以用该声明的类型来限制。
	// 这里单独申明的作用，主要是为了多次多地方复用 类型形参限制。
	~string|~int|~float32  // 接口定义时，可以包含一些类型。~标识只要底层是这些类型都可以，不管是 type OtherType int 后的 OtherType 还是 int 都是在类型限制范围之内。
}

type ValueConstraintOne interface {
	F1()
}

func CallMaps[TypeK comparable, TypeV ValueConstraint](vv map[TypeK]TypeV) TypeV {
	var ret TypeV
	for _, v := range vv {
		ret += v
	}
	return ret
}



func RunGenerics() {
	vv := map[int]string {
		1:"sddd",
		2: "bbbb",
	}
	v := CallMaps(vv) //函数调用时模板类型可以根据参数类型自动推导。而不需要在调用时指定模板类型参数。 
	fmt.Println(v)

	vvv := map[int]int {
		1:1,
		2:12,
	}
	v2 := CallMaps[int, int](vvv)
	fmt.Println(v2)


	CallInterface()
}

// 非泛型 类型中，其方法 不支持 泛型
// type AA struct {}
// func(a* AA)[T any|int](v int) {
// }
// 要实现自定义类型方法支持泛型，那么自定义的类型必须是泛型，比如：
type GenericType[T any | int] struct{}
func (c *GenericType[T]) call(v T) {
	//
}

// 不能定义匿名函数的泛型函数，比如：
// func  demoTest {
//  f := func[T any](v T) {} //定义匿名的泛型函数。
//  f(100)
// }
// 但是可以在匿名函数中调用 泛型函数，比如:
// func() {
//    callGenericFunc[int, float32](1, 123.12) //可以直接调用匿名的泛型函数。
// }()




// 泛型类型的嵌套， 泛型的类型有多种； 另外一种依赖另外一种。比如：T, []T, map[comparable]T, 
// type DemoCallNested[T any, S []T] struct {
	// Data T 
	// Arr S
// }
//实例化实例： var x DemoCallNested[int, []int] 

//给interface 定义泛型：
type DemoInterface[T any] interface {
	Call(data T)
}

type BusiOne[T any] struct {
	Data T
}
func (s* BusiOne[T]) Call(data T) {
	s.Data = data
	fmt.Println(s.Data)
}

func CallInterface() {
	var d1 DemoInterface[int] = &BusiOne[int] {
		Data: 100,
	}
	d1.Call(200)
}
