package helloworld

import "fmt"

const (
	spanish                  = "Spanish"
	french                   = "French"
	emptyStrGreetDestination = "World"
	englishGreetingPrefix    = "Hello, "
	spanishGreetingPrefix    = "Hola, "
	frenchGreetingPrefix     = "Bonjour, "
)

func SayHello(name, language string) string {

	if name == "" {
		name = emptyStrGreetDestination
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := englishGreetingPrefix

	switch language {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	}

	return prefix
}

func main() {
	fmt.Println(SayHello("Ayyash", ""))
}
