package main

import "fmt"
import "time"

func main(){

  for{
    fmt.Println("hello world")
    a := 1
    b := 2
    if(a==b){
      break
    }
  }
  time.Sleep(1 * time.Second)
  
}
