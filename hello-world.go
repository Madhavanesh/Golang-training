package main

import (
	"fmt"
)

//	func main() {
//		fmt.Println("Hello World")
//	}
const s string = "constant"

func main() {
	// fmt.Println("go" + "lang")
	// fmt.Println("1+2 : ", 1+2)
	// fmt.Println("7.0/14.0 : ", 7.0/14.0)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!false)
	// var a = "intial"
	// var b, c int = 1, 2
	// fmt.Println(a, "\n", b, c)
	// f := "Wastemad"
	// fmt.Println(f)
	// fmt.Println(s)
	// const n = 4000000000
	// const d = 3e20 / n
	// fmt.Println(d)
	// fmt.Println(int64(d))
	// fmt.Println(math.Sin(d))
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	// if n := 9; n < 9 {
	// 	fmt.Println("l")
	// } else if n < 10 {
	// 	fmt.Println("m")
	// }
	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("noon")
	// default:
	// 	fmt.Println("afternoon")
	// }
	// whatamI := func(i interface{}) {
	// 	switch i.(type) {
	// 	case bool:
	// 		fmt.Println("boolean")
	// 	case int:
	// 		fmt.Println("integer")
	// 	case string:
	// 		fmt.Println("String")
	// 	}
	// }
	// whatamI("Mad")
	// whatamI(true)
	// whatamI(9)
	// var a[5]int
	// fmt.Println(a)
	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(b)
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))
	s = append(s, "d", "e", "f")
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(s[:5])
}
