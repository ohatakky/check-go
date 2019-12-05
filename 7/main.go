package main

func f(ns [2]int) {
	println(&ns)
	ns[0] = 100
}

func main() {
	var ns [2]int
	println(&ns)
	f(ns)
	println(ns[0])
}
