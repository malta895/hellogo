package btrees

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}

		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}

		return true
	}

	return true
}

func runBtreeExercise() {
	fmt.Println("hello, we start Walking the tree")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("start checking if trees are the same")
	areTreesTheSame := Same(tree.New(1), tree.New(1))
	fmt.Println(areTreesTheSame)

	areTreesTheSame = Same(tree.New(1), tree.New(2))
	fmt.Println(areTreesTheSame)

	areTreesTheSame = Same(tree.New(1), tree.New(3))
	fmt.Println(areTreesTheSame)
}
