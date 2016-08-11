package main

import "fmt"

type S struct {
    x, y int
}

func main() {
    s := []*S{&S{1,2},&S{3,4}}
    fmt.Println("s", s)
    fmt.Println(*s[0], *s[1])
    c := make(chan []*S)
    f := func() {
        t := <-c
        fmt.Println("t", t)
        fmt.Println(*t[0], *t[1])
        t[0] = new(S)
        fmt.Println("new t", t)
        close(c)
    }
    go f()
    c <- s
    <- c
    fmt.Println("post s", s)
    fmt.Println(*s[0], *s[1])
} 
