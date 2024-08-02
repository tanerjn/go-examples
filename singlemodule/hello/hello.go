package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("Received: ")
    log.SetFlags(0)
    
    // Slice of names
    names :=  []string{"Alice", "Bob", "Gladys"}

    // Request a greeting message.
    messages, err := greetings.MultipleHellos(names)

    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(messages)
}
