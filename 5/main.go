package main

func main() {
	var a struct{ N int }
	b := a
	a.N = 100
	println(b.N)
}
