package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
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
    
    // This doesn't seem to work... 
    // This should be incompatible with most other flags
    // and just toupper the string
    tuPtr := flag.String("u", "", "extra string")
    //var svar string
    //flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    //fmt.Println("aaa:", *aaaPtr)
    //fmt.Println("hhh:", *hhhPtr)
    //fmt.Println("aNum:", *aNumPtr)
    //fmt.Println("hNum:", *hNumPtr)
    //fmt.Println("tail:", flag.Args())

    //fmt.Println("svar:", svar)
    //fmt.Println("tail:", flag.Args())

    //for i := 0; i < *numbPtr; i++ { 
     //   fmt.Printf("%s", svar)
    // }
    // for i := 0; i < *numb2Ptr; i++ {
    //    fmt.Printf("%s", *wordPtr)
    // }
    justcase := len(strings.TrimSpace(*tuPtr))
    if justcase > 1 {
        tu(*tuPtr)
	// OK, yes, this is bad, 
	// but it is all so bad
        os.Exit(1)
    } 

    spew(*aaaPtr, *aNumPtr)
    spew(*hhhPtr, *hNumPtr)
    
    // de bug 
    dafuq := len(strings.TrimSpace(*xtraPtr))    

    //if len(strings.TrimSpace(*xtraPtr)) < 1 {
    if dafuq > 1 {
        spew(*xtraPtr, *xNumPtr)
    }

    if *nlPtr == true {
        fmt.Println("") 
    }
}
