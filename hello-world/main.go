package main

import "fmt"

const (
	defaultName   = "world"
	defaultPrefix = "Hello"
)

func Greet(name string, lang string) string {
	prefix := defaultPrefix

	if name == "" {
		name = defaultName
	}

	switch lang {
	case "sp":
		prefix = "Hola"
	case "fr":
		prefix = "Bonjour"
	}

	return fmt.Sprintf("%s, %s", prefix, name)
}
func main() {
	fmt.Println(Greet("pal", ""))
}
