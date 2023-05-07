package main

import (
	"example/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

type Person struct {
	name string;
}

func main() {

	message := "World";
	fmt.Printf("Hello, %v.\n", message);
	fmt.Println(quote.Go());

    log.SetPrefix("greetings: ")
    log.SetFlags(log.Llongfile)

	person1:= Person{"Thomas"};
	person2:= Person{"Torben"};
	person3:= Person{"Janis"};
	person4:= Person{"Alfred"};
	 
	fmt.Println(getGreetingMessage(person1.name));

	names := []string{person2.name, person3.name, person4.name}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	 
	fmt.Println(messages)
}


func getGreetingMessage(name string) (string) {
	message, err := greetings.Hello(name);

    if (err != nil) {
        log.Fatal(err);
    }


	return message;
}