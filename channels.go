package main

import (
    "fmt"
    "time"
)

func worker(done chan bool){
    fmt.Println("Working")
    time.Sleep(2 * time.Second)
    fmt.Println("Done!")

    done <- true

}

func main () {

    //Create new channel called "done"
    done := make(chan bool,1)

    //Run go routine, giving them the channel
    go worker(done)

    //Block Channel
    <-done

}