package hello

import "fmt"

const spanish = "Spanish"
const french = "French"

const frenchPrefix = "Bonjour"
const englishPrefix = "Hello"
const spanishPrefix = "Hola"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return spanishPrefix + ", " + name + "!"
	}

	if lang == frenchPrefix {
		return frenchPrefix + ", " + name + "!"
	}

	return fmt.Sprintf("%v, %v!", englishPrefix, name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
func HelloSwitch(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%v, %v!", greetingPrefix(lang), name)
}

func main() {
	fmt.Println(Hello("World", "English"))
}
