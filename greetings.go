package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v, welcome!", name)
	return message
}
