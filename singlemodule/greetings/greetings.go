package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)


// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}


func MultipleHellos(names []string) (map[string]string, error){
    //A map associated with messages.
    messages := make(map[string]string)
    //Loop through the received  slice of names, calling the function to get a message for each name
    for _, name := range names {
        message, err := Hello(name)
            if err != nil {
                return nil, err
            }
    // insdie the map associate retrieved message with name
            messages[name] = message
        }
    return messages, nil
}



//randomFormat returns a greeting   
func randomFormat() string{
    //A slice of message formats
    formats := []string{
        "Hi %v.",
        "Nice to see you, %v!",
        "Ciao %v!",
    }

    //return randomly selected message format by specfying random index.
    return formats[rand.Intn(len(formats))]
}


