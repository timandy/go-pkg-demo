package go_pkg_demo

import "fmt"

func PrintInfo(value interface{}) interface{} {
	fmt.Println(value)
	println(1)
	return value
}
