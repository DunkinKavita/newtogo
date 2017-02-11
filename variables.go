package main

import "fmt"

func main(){

    var a string = "This is a string"
    fmt.Println(a)

    var b,c int = 1,2
    fmt.Println("b+c:",b+c)

    var d = (true&&false)
    fmt.Println(d)

    var e int
    fmt.Println("let's see what the default int value is:", e)

    f:= "short"
    fmt.Println("this was not explicitly declared as string:", f)

}
