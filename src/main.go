package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main()  {
	eb := englishBot{
		message: "Hi there",
	}

	sb := spanishBot{
		message: "Hola muchachos",
	}

	printGreetings(eb)
	printGreetings(sb)

	fmt.Println("***Getting some data with a HTTP Get request***")

	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("HTTP Get operation failed with reason:", err)
		os.Exit(1)
	}


	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Println("HTTP response body couldn't be sent to stdout: ", err)
	}
}
