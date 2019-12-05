package main

func main() {
	as := []struct{ N int }{{N: 10}, {N: 20}, {N: 30}}
	fs := make([]func(), len(as))

	for i := range as {
		println(&i)
		fs[i] = func() {
			println(&i)
			as[i].N *= 10
		}
	}

	for _, f := range fs {
		f()
	}
	println(as[0].N, as[1].N, as[2].N)
}
