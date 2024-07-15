package main 

import "fmt"

func main() {
	PrintGreeting()
	NotifyStart()
	fmt.Println(GetWelcomeMessage())
}

func PrintGreeting(){
	fmt.Println("Hello Word")
}

   func NotifyStart() {
	fmt.Println("Programm started")
   }

  func GetWelcomeMessage() string {
return "Welcome"
  }