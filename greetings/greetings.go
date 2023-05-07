package greetings

import (
	"errors"
	"fmt"
)


func Hello(name string) (string, error) {
	err:= isValidName(name);
	if err != nil {
		return "", err;
	}

	message:= fmt.Sprintf("Hi, %v. Welcome!", name);
	return message, nil;
}

func Bye(name string) (string, error) {
	err:= isValidName(name);
	if err != nil {
		return "", err;
	}

	message:= fmt.Sprintf("Goodbye, %v.", name);
	return message, nil;
}

func isValidName(name string) error {
	if(name == "" || len(name) == 0){
		return errors.New("empty name");
	}

	return nil;
}