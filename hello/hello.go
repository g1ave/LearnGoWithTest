package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Halo, "
const frenchHelloPrefix = "Bonjour, "

const english = "English"
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
