package greetings

import "fmt"
import gd "github.com/vvdvortsova/greetingsdeps"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. %v!", name, gd.Greeting())
    return message
}
