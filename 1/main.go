package main

func main() {
	as := []struct{ N int }{{N: 10}, {N: 20}}

	for _, a := range as {
		a.N *= 10
	}

	println(as[0].N, as[1].N)
}
