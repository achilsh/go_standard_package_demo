package genericdemo

import "fmt"

//本章节主要 介绍 go的模板使用。

type ValueConstraint interface { 
	// 申明模板类型限制的类型。这样在写模板时模板的形参类型可以用该声明的类型来限制。
	// 这里单独申明的作用，主要是为了多次多地方复用 类型形参限制。
	string|int|float32
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
}