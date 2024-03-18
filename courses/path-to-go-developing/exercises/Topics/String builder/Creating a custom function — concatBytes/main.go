package main

import (
    "fmt"
    "strings" 
) 

// Modify the concat function template below to properly implement the concatBytes function
func concat?(bytes ...byte) string {
    var builder strings.Builder
    for _, x := range bytes {
        builder.Write?(x)
    }
    return builder.String()
}

func main() {
    fmt.Println(concatBytes('J', 'B', 'A', 'c', 'a', 'd', 'e', 'm', 'y'))
}