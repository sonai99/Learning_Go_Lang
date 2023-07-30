package main

import "fmt"

func main() {
	// map[key-type]val-type
	m := make(map[string]int)

	m["k1"] = 8
	m["k2"] = 9
	fmt.Println("Map: ", m)
	fmt.Println(m["k1"])
	v := m["k1"]
	fmt.Println(v)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs :", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
