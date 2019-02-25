package main

func main()  {
	eb := englishBot{
		message: "Hi there",
	}

	sb := spanishBot{
		message: "Hola muchachos",
	}

	printGreetings(eb)
	printGreetings(sb)
}