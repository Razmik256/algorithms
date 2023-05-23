package main

import "fmt"

type mNode struct {
	value int
	next  []*mNode
}

func init_mnode(val int) mNode {
	return mNode{val, nil}
}

func find_smallest_path(start *mNode) int {
	smallest_val := 999999999
	for _, p := range start.next {
		if p.value == 0 {
			smallest_val = 0
			break
		}
		f := find_smallest_path(p)
		if f < smallest_val {
			smallest_val = f
		}
	}
	return start.value + smallest_val
}

func find_smallest_path_init() {
	start := init_mnode(0)
	n1 := init_mnode(1)
	n2 := init_mnode(5)
	n3 := init_mnode(3)
	n4 := init_mnode(2)
	n5 := init_mnode(1)
	n6 := init_mnode(3)
	n7 := init_mnode(1)
	n8 := init_mnode(1)
	fin := init_mnode(0)
	start.next = []*mNode{&n1}
	n1.next = []*mNode{&n2}
	n2.next = []*mNode{&n3, &n4}
	n3.next = []*mNode{&n7}
	n7.next = []*mNode{&n8}
	n8.next = []*mNode{&fin}
	n4.next = []*mNode{&n5}
	n5.next = []*mNode{&n6}
	n6.next = []*mNode{&fin}
	fmt.Println(find_smallest_path(&start))

}
