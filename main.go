package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Version 1
	// // 'make' takes a type of a slice and number of elements we want it to be initialised with
	// byteSlice := make([]byte, 99999) // "give me an empty byte slice with space for 99999 elements"

	// // read the body of the response into the byte slice
	// resp.Body.Read(byteSlice)

	// // print out the byte slice which contains all the html in the body of the response
	// fmt.Println(string(byteSlice))

	// Version 2
	// condensed version of the above code ^^
	io.Copy(os.Stdout, resp.Body)

}
