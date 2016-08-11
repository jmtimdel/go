package main

//import "code.google.com/p/go-tour/tree"
import "./tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
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
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    fmt.Println(t1, t2)
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for v1 := range(ch1) {
        v2, ok := <-ch2
        if v1 != v2 || !ok {
            return false
        }
    }
    _, ok := <-ch2
    return !ok
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
