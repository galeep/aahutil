package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "runtime"
)

func spew(stuff string, rounds int) { 
   for i := 0; i < rounds; i++ {
        fmt.Printf("%s", stuff)
    }
}

func tu(msg string) { 
    fmt.Println(strings.ToUpper(msg))
}

func main() {

    aaaPtr := flag.String("a", "A", "a string")
    hhhPtr := flag.String("h", "H", "a string")
    aNumPtr := flag.Int("an", 42, "an int")
    hNumPtr := flag.Int("hn", 3, "an int")
    xtraPtr := flag.String("x", "", "extra string")
    xNumPtr := flag.Int("xn", 1, "extra int")
    nlPtr := flag.Bool("n", true, "a bool")
    
    // This doesn't seem to work.
    // This should be incompatible with most other flags
    // and just toupper the string
    tuPtr := flag.String("u", "", "extra string")
    
    flag.Parse()
    
    justcase := len(strings.TrimSpace(*tuPtr))
    if justcase > 1 {
        tu(*tuPtr)
	// OK, yes, this is bad, 
	// but it is all so bad
        os.Exit(1)
    } 

    spew(*aaaPtr, *aNumPtr)
    spew(*hhhPtr, *hNumPtr)
    
    // Still quirky
    wat := len(strings.TrimSpace(*xtraPtr))    

    //if len(strings.TrimSpace(*xtraPtr)) < 1 {
    if wat > 1 {
        spew(*xtraPtr, *xNumPtr)
    }

    if *nlPtr == true {
        fmt.Println("") 
    }
}
