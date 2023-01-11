package main

import (
	"fmt"
	"github.com/wubba-com/algorithms_and_structures/utils"
)

//var (
//	_ = f("w", x)
//	x = f("x", z)
//	y = f("y", x)
//	z = f("z")
//)
//
//func f(s string, deps ...int) int {
//	print(s)
//	return 0
//}

func main() {

	//f("/n")
	var s1, s2, s3 int
	_, err := fmt.Scanln(&s1, &s2, &s3)
	if err != nil {
		fmt.Println(err)
		return
	}

	var t int
	for v := range utils.Download(s1, s2, s3) {
		t += v
	}
}
