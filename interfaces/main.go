package main

import "fmt"


type bot interface {
    getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}


func main()  {
    eb := englishBot{}
    sb := spanishBot{}

    // the printGreeting function only accepts type bot but here since the englishBot
    // struct implements the method getGreeting, it is accepted.
    printGreeting(eb)
    printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
    // very custom logic for generating an english greeting
    return "Hi There."
}

func (sb spanishBot) getGreeting() string {
    return "Hola!"
}


func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//    fmt.Println(eb.getGreeting())
//}
//
//func printGreeting(eb spanishBot) {
//    fmt.Println(eb.getGreeting())
//}