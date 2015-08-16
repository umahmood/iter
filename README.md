# iter

A Go library which allows you to iterate over a sequence integers, in a simple 
and memory efficient way.

# Installation

> go get github.com/umahmood/iter

# Usage

When you need to loop over a sequence of integers, you may write:

```for i := 0; i < 10; i++ {
    // ...
}```

With iter you can simply write:

```for i := range iter.N(10) {
    fmt.Printf("%d ", i)
}```

Output:

```
0 1 2 3 4 5 6 7 8 9
```

A nifty thing is that *iter.N(10)* and *iter.N(100)* consume the same amount of 
storage. This is possible because in Go, the empty *struct* consumes no storage.

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

