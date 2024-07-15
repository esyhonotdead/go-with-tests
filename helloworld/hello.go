package helloworld

const (
	// languages
	spanish = "Spanish"
	english = "English"
	french  = "French"

	// prefixes
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" && language == "" {
		name = "World!"
		language = english
	}

	prefix := englishHelloPrefix

	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}
