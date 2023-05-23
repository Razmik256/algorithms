package main

import "fmt"

type dNode struct {
	name     string
	conNodes map[*dNode]int
}

func find_djikstra_first_iter(start *dNode, fin *dNode, from *dNode) int {
	arr := []int{}
	for n, v := range start.conNodes {
		arr = append(arr, v+find_djikstra(n, fin, start))
	}
	smallest := 99999
	for _, a := range arr {
		if a < smallest {
			smallest = a
		}
	}
	return smallest
}

func find_djikstra(start *dNode, fin *dNode, from *dNode) int {
	smallest_path := new(dNode)
	smallest := 99999
	for n, v := range start.conNodes {
		if n == from {
			continue
		}
		if v < smallest {
			smallest = v
			smallest_path = n
		}
	}
	if smallest_path == fin {
		return smallest
	}
	return smallest + find_djikstra(smallest_path, fin, start)
}

func new_con(n1, n2 *dNode, val int) {
	n1.conNodes[n2] = val
	n2.conNodes[n1] = val
}

func find_djikstra_init() {
	s := dNode{"s", make(map[*dNode]int)}
	a := dNode{"a", make(map[*dNode]int)}
	b := dNode{"b", make(map[*dNode]int)}
	c := dNode{"c", make(map[*dNode]int)}
	d := dNode{"d", make(map[*dNode]int)}
	g := dNode{"g", make(map[*dNode]int)}
	f := dNode{"f", make(map[*dNode]int)}

	new_con(&s, &a, 2)
	new_con(&s, &b, 1)
	new_con(&a, &g, 3)
	new_con(&b, &g, 5)
	new_con(&a, &b, 4)
	new_con(&b, &c, 1)
	new_con(&c, &d, 1)
	new_con(&d, &f, 2)
	new_con(&g, &f, 1)
	fmt.Println(find_djikstra_first_iter(&s, &f, nil))
}
