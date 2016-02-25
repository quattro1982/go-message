package main

import (
    "fmt"
)

func ping(pings chan<- string, pongs <-chan string, done chan string,  msg string){
    pings <- msg
    reply := <- pongs
    fmt.Println(reply)

    done <- ""
}

func pong(pings <-chan string, pongs chan<- string){
    msg:= <-pings
    pongs <- msg

}

func main(){
    done := make (chan string)

    pings:= make (chan string)
    pongs:= make (chan string)
    go ping(pings, pongs, done, "passed message")
    pong(pings, pongs)
    
    <- done

    //fmt.Println(<-pongs)
}