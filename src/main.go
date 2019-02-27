package main

import (
	"fmt"
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
		fmt.Println("Get operation failed with reason:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

