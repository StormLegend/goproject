package main
import "fmt"

const helloPrefix = "Hello,"
const spanish = "Spanish"
const spanishHelloPrefix = "Hola,"
const french ="French"
const frenchHelloPrefix = "Bonjour,"

func Hello(name string, lang string) string{
    prefix := helloPrefix	
    if name == "" {
        name = "World"
    }
    /*
    if lang == spanish {
        return spanishHelloPrefix+name
    }
    if lang == french {
        return frenchHelloPrefix+name
    }
    return helloPrefix+name
    */
    switch lang {
        case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
    }
    return prefix + name
}


func main(){
    fmt.Println(Hello("world",""))

}

