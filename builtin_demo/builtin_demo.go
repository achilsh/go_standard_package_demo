package builtin_demo

import "fmt"

const (
	_ = iota
	A
	B
	C
	D
)

func BuiltinDemoRun() {
	fmt.Println(A, B, C, D)
}
