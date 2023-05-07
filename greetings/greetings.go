package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)


func Hello(name string) (string, error) {
	err:= isValidName(name);
	if err != nil {
		return "", err;
	}

	message:= fmt.Sprintf(randomGreeting(), name);
	return message, nil;
}

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

func isValidName(name string) error {
	if(name == "" || len(name) == 0){
		return errors.New("empty name");
	}

	return nil;
}

func init () {
	//rand.Seed(time.Now().UnixNano());
}

func randomGreeting() string {
	greetingList:= []string {
		"Hi %v. Welcome!",
		"Great to see you, %v.",
		"Hail, %v! Well met!",
	}

	return greetingList[rand.Intn(len(greetingList))];
}