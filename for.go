package main

import "fmt"

func main()  {
    i:= 1

    for i<=4 {
        fmt.Println(i)
        i=i+1
    }
    
    for j:=7; j<10; j++ {
        fmt.Println(j)
    }

    for n:=2 ;n<=10; n++ {
        if n%2==0 {
            fmt.Println(n)
        }
    }
}

