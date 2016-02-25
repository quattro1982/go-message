package main

import (
    "fmt"
    "time"
    "strconv"
)

func clientA(clientA chan string, clientB chan string, done chan string ){
    for i:=0; i <=10; i++ {
        msg:= "msg #"+strconv.Itoa(i)
        time.Sleep(3 * time.Second)
        clientA <-msg
        time.Sleep(3 * time.Second)
        fmt.Println(<-clientB)
        
    }
   // msg := "Client A sends"
   
}

func clientB(clientA chan string, clientB chan string, done chan string){
    msg:= <-clientA
    fmt.Println(msg)
    time.Sleep(2 * time.Second)
    reply:= "In reply to "+ msg
    clientB <- reply
    
}

func main(){
    done := make (chan string)
    chA:= make(chan string,1)
    chB:= make(chan string,1)
    go clientA(chA, chB, done)
    go clientB(chA, chB, done)

    <-done


}