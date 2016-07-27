package main

import (
    "fmt"
)

func syracuse(x,altitude_max,tps int) (bool,int,int) {

    if x > 0 {
    if  x == 1 {

        return true,tps,altitude_max
     } else if x%2 == 0 {

      tps = tps+1
       if altitude_max <= x {
       altitude_max = x
       }
       return syracuse(x/2,altitude_max,tps)

    } else {

        tps = tps+1
       if altitude_max <= x {
       altitude_max = x
       }
       return syracuse(3*x+1,altitude_max,tps)

    }
} else {
     return false,0,0
}

}

func main() {
    var altitude_max,tps int
    fmt.Println(syracuse(127,altitude_max,tps))
}
