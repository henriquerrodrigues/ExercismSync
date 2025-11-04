package airportrobot
import "fmt"
// Write your code here.

type Greeter interface{
    LanguageName() string
    Greet(visitorName string) string
}

type GermanGreeter struct{}

func (g GermanGreeter) LanguageName() string{
    return "German"
}

func (g GermanGreeter) Greet(visitorName string) string{
    return fmt.Sprintf("Hallo %s!", visitorName)
}

type Italian struct{}

func (i Italian) LanguageName() string{
    return "Italian"
}

func (i Italian) Greet(visitorName string) string{
    return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string{
    return "Portuguese"
}

func (p Portuguese) Greet(visitorName string) string{
    return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string{
    language := greeter.LanguageName()
    greeting := greeter.Greet(visitorName)

    return fmt.Sprintf("I can speak %s: %s", language, greeting)
}




// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
