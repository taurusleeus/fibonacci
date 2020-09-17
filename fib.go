package main

import "fmt"

func fib(n int) []int{
  f := make([]int,n)
  f[0] = 0
  f[1] = 1
  
  for i := 2; i < n; i++{
    f[i] = f[i-1] + f[i-2]
  }
  
  return f
  
}

func main(){
  
  n := 0
  
  fmt.Print("Enter the number of terms, start from 2 : ")
  fmt.Scanln(&n)
  
  fmt.Print("Fibonacci series : ")
  fmt.Println(fib(n))

}
