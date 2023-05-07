package main

import (
	"example/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {

	message := "World";
	fmt.Printf("Hello, %v.\n", message);
	fmt.Println(quote.Go());

	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(log.Llongfile)
	 
	fmt.Println(getGreetingMessage("Thomas"));
}


func getGreetingMessage(name string) (string) {
	message, err := greetings.Hello(name);

    if (err != nil) {
        log.Fatal(err);
    }


	return message;
}