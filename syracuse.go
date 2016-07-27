package main

import (
    "fmt"
    "math"
)


func check(nbr float64) bool {
    if nbr%2 ==0 {
        return true
    }else {
        return false
    }
}

func syracuse(x int) (bool,int,int) {

    if x == 4 || x == 2 || x == 1 {
       return true
    } else if x%2 == 0 {
        return check(x/2)
    } else {
        return check(3*x+1)
    }

}
func main() {
    fmt.Println(8%2)
}
