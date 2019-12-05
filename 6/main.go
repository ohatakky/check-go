package main

type A struct{ N int }

func f(a A) { a.N = 100 }

func main() {
	var a A
	f(a)
	println(a.N)
}
