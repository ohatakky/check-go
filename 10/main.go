package main

func main() {
	a := []int{10, 20, 30}
	println(&a)
	println(cap(a))
	b := append(a, 40, 50)
	println(&b)
	a[0] = 100
	c := append(b, 60, 70)
	b[1] = 200
	println(c[0] + c[1])

	a = append(a, 2)
	println(cap(a))
	println(&a)
}
