package main

import(
  "fmt"
  "math"
)

func main() {
  var e float64
  var n float64
  var x float64
  fmt.Scanln(&x)
  for n=1;n<=x;n++{
    e=(1+1/n)
    e=math.Pow(e,n)
  }
  fmt.Println(e)
}