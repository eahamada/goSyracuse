package main

import (
    "fmt"
)

func syracuse(x,altmax,indice int) (bool,int,int) {


    if  x == 1 {

    return true,indice,altmax

   } else if x%2 == 0 {

       indice = indice+1
       if altmax <= x {
       altmax = x
       }
      return syracuse(x/2,altmax,indice)

    } else {

        indice = indice+1
       if altmax <= x {
       altmax = x
       }
       return syracuse(3*x+1,altmax,indice)

    }
}

func main() {
    var altmax,indice int
    fmt.Println(syracuse(15,altmax,indice))
}
