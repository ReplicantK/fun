package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	latinHelloPrefix = "Salve, "

	spanish = "Spanish"
	french = "French"
	latin = "Latin"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	} 

	switch lang {
		case spanish:
			return spanishHelloPrefix + name
		case french:
			return frenchHelloPrefix + name
		case latin:
			return latinHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
