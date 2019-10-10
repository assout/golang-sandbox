package main
2
​
3
import (
4
    "fmt"
5
    "math"
6
)
7
​
8
func main() {
9
    var x, y int = 3, 4
10
    var f float64 = math.Sqrt(float64(x*x + y*y))
11
    var z uint = uint(f)
12
    fmt.Println(x, y, z)
13
}
14
​