package main

func f(ns []int) {
	println(&ns)
	ns[0] = 100
}

func main() {
	var ns []int
	println(&ns)
	f(ns)
	println(ns[0])
}
