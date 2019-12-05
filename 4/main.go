package main

func f(n int) { n  = 100}

func main() {
	var n int
	f(n)
	println(n)
}
