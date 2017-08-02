package main

import (
"fmt"
"os"
)

func main() {
	numberOfLines := 0
	//check for the correct number of arguments
	if len(os.Args) == 1 {
        fmt.Println("Please, pass file name as parameter")
		return
	}
    if len(os.Args) > 2 {
        fmt.Println("too much parameters")
        return
    }
	fileName := os.Args[1]

	file, err := os.Open(fileName)
    if err != nil {
        // handle the error
        fmt.Println("No such file")
        return
    }
    defer file.Close()

    // get the file size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // read the file
    bytes := make([]byte, stat.Size())
    _, err = file.Read(bytes)
    if err != nil {
        return
    }
    fileString := string(bytes)
    
    //iterate over the file for searching '\n'
    for _, c := range fileString {
    	if c == '\n' {
    		numberOfLines++
    	}
    }
    fmt.Println(numberOfLines)
}