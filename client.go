package main
import (
    "fmt"
    "io/ioutil"
    "os"
)

    
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    //readfile("/Users/jlonan/Pictures/Scan.jpeg")
    readfile("/tmp/testfile")
}

func readfile(filename string){
    //Read file and load to memory
    dat, err := ioutil.ReadFile(filename)
    check(err)
    //fmt.Print(string(dat))

    f, err := os.Open(filename)
    check(err)

    b1 := make([]byte,1024)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))


}