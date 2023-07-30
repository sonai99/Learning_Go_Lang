package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(s[2])

	s = append(s, "d")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c", c)

	l := s[2:3]
	fmt.Println(l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
