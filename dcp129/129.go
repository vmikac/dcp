package main

import "fmt"

func main() {
   var n float64 = 1e6
   var low float64 = 0
   high := n
   guess := (high+low)/2
   for i:=0; i<64 ; i++ {
     if guess * guess > n {
        high = guess
     } else {
        low = guess
     }
     guess = (high+low)/2
   }
   fmt.Println(guess)
}

