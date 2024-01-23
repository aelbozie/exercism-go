package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type Italian struct{}

func (l Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (l Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct{}

func (l Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func (l Portuguese) LanguageName() string {
	return "Portuguese"
}
