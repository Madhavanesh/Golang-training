package main

import (
	"fmt"
	//"unicode/utf8"
	//"math"
)

//	func main() {
//		fmt.Println("Hello World")
//	}
//const s string = "constant"

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
	// var s []string
	// fmt.Println(s, s == nil, len(s) == 0)
	// s = make([]string, 3)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, "d", "e", "f")
	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println(c)
	// fmt.Println(s[:5])
	// m := make(map[string]int)
	// m["k1"] = 2
	// m["k2"] = 3
	// fmt.Println(m)
	// delete(m, "k2")
	// _, prs := m["k1"]
	// fmt.Println(prs)
	// n := map[string]int{"f": 1, "c": 2}
	// fmt.Println(n)
	// n := []int{1, 3, 4}
	// sum := 0
	// for _, n := range n {
	// 	sum += n
	// }
	// for i, n := range n {
	// 	if n == 4 {
	// 		fmt.Println(i)
	// 	}
	// }
	// fmt.Println(sum)
	// m := map[string]int{"f": 1, "b": 2}
	// for k, v := range m {
	// 	fmt.Println(k, "->", v)
	// }
	// for k := range m {
	// 	fmt.Println(k)
	// }
	// for i, c := range "go" {
	// 	fmt.Println(i, c)
	// }
	// fmt.Println(plus(1, 4, 5))
	// a, b := vals()
	// fmt.Println(a, b)
	// _, c := vals()
	// fmt.Println(c)
	// sum(1, 2)
	// sum(1, 3, 5)
	// n := []int{1, 2, 3, 4}
	// sum(n...)
	// nextInt := intseq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// nextint := intseq()
	// fmt.Println(nextint())
	// fmt.Println(nextInt())
	// fmt.Println(fact(7))
	// var fib func(n int) int
	// fib = func(n int) int {
	// 	if n < 2 {
	// 		return n
	// 	}
	// 	return fib(n-1) + fib(n-2)
	// }
	// fmt.Println(fib(7))
	// i := 1
	// zeroval(i)
	// fmt.Println(i)
	// zeroptr(&i)
	// fmt.Println(i)
	// fmt.Println(&i)
	// const s = "สวัสดี"
	// fmt.Println(len(s))
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%x", s[i])
	// }
	// fmt.Println(utf8.RuneCountInString(s))
	// for idx, runevalue := range s {
	// 	fmt.Printf("%#U starts at %d\n", runevalue, idx)
	// }
	// fmt.Println("\nUsing DecodeRuneInString")
	// for i, w := 0, 0; i < len(s); i += w {
	// 	runevalue, width := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%#U starts at %d\n", runevalue, i)
	// 	w = width
	// 	examineRune(runevalue)
	// }
	// fmt.Println(person{"Bob", 40})
	// fmt.Println(person{name: "Alice", age: 49})
	// fmt.Println(person{name: "John"})
	// fmt.Println(&person{name: "Ann", age: 40})
	// fmt.Println(newperson("James"))
	// s := person{name: "Reb", age: 47}
	// fmt.Println(s.age)
	// sp := &s
	// fmt.Println(sp.name)
	// dog := struct {
	// 	name   string
	// 	isgood bool
	// }{
	// 	"gold", true,
	// }
	// fmt.Println(dog)
	// r := rect{width: 5, height: 10}
	// c := circle{radius: 7}
	// fmt.Println(r.area())
	// fmt.Println(r.perm())
	// rp := &r
	// fmt.Println(rp.area())
	// fmt.Println(rp.perm())
	// measure(r)
	// measure(c)
	co := container{base: base{n: 1}, str: "som"}
	fmt.Printf("co={num:%v,str:%v}\n", co.n, co.str)
	fmt.Println(co.base.n)
	fmt.Println(co.describe())
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println(d.describe())
}

//	func plus(a int, b int, c int) int {
//		return a + b + c
//	}
//
//	func vals() (int, int) {
//		return 3, 7
//	}
//
//	func sum(n ...int) {
//		fmt.Print(n, " ")
//		total := 0
//		for _, n := range n {
//			total += n
//		}
//		fmt.Println(total)
//	}
//
//	func intseq() func() int {
//		i := 0
//		return func() int {
//			i++
//			return i
//		}
//	}
//
//	func fact(n int) int {
//		if n == 0 {
//			return 1
//		}
//		return n * fact(n-1)
//	}
//
//	func zeroval(ival int) {
//		ival = 0
//	}
//
//	func zeroptr(iptr *int) {
//		*iptr = 0
//	}
//
//	func examineRune(r rune) {
//		if r == 't' {
//			fmt.Println("found tee")
//		} else if r == 'ส' {
//			fmt.Println("found so sua")
//		}
//	}
// type person struct {
// 	name string
// 	age  int
// }

//	func newperson(name string) *person {
//		p := person{name: name}
//		p.age = 42
//		return &p
//	}
// type geomentry interface {
// 	area() float64
// 	perm() float64
// }
// type rect struct {
// 	width, height float64
// }
// type circle struct {
// 	radius float64
// }

//	func (r rect) area() float64 {
//		return r.width * r.height
//	}
//
//	func (r rect) perm() float64 {
//		return 2 * (r.width + r.height)
//	}
//
//	func (c circle) area() float64 {
//		return math.Pi * c.radius * c.radius
//	}
//
//	func (c circle) perm() float64 {
//		return 2 * math.Pi * c.radius
//	}
//
//	func measure(g geomentry) {
//		fmt.Println(g)
//		fmt.Println(g.area())
//		fmt.Println(g.perm())
//	}
type base struct {
	n int
}
type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.n)
}
