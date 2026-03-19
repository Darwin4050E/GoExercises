// Package airportrobot implements functionality so that a robot can greet people in their native language.
package airportrobot

type Greeter interface {
	LanguageName() string
	Greet(nameVisitor string) string
}

// SayHello returns the desired greeting string.
func SayHello(visitorName string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(visitorName)
}

type Italian struct{}

// LanguageName returns the string Italian.
func (i Italian) LanguageName() string {
	return "Italian"
}

// Greet returns a string with the greeting message in Italian.
func (i Italian) Greet(visitorName string) string {
	return "Ciao " + visitorName + "!"
}

type Portuguese struct{}

// LanguageName returns the string Portuguese.
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet returns a string with the greeting message in Portuguese.
func (p Portuguese) Greet(visitorName string) string {
	return "Olá " + visitorName + "!"
}
