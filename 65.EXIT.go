The Go PlaygroundRun  Format   Imports  Share About
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25

// Use `os.Exit` to immediately exit with a given
// status.

package main

import "fmt"
import "os"

func main() {

    // `defer`s will _not_ be run when using `os.Exit`, so
    // this `fmt.Println` will never be called.
    defer fmt.Println("!")

    // Exit with status 3.
    os.Exit(3)
}

// Note that unlike e.g. C, Go does not use an integer
// return value from `main` to indicate exit status. If
// you'd like to exit with a non-zero status you should
// use `os.Exit`.

// $ go run exit.go
// exit status 3

// $ go build exit.go
// $ ./exit
// $ echo $?
// 3