package main

func f(ns []int) {
	ns = append(ns, 100)
}

func main() {
	var ns []int
	f(ns)
	println(len(ns))
}
