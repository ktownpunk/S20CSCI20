// Programmer name: Jacob Harter
// Date completed:  3/29/20
// Description:Random feedback

package main

import "fmt"
//import "math/rand"
//import "time"
func greeting(userInput string){ 
fmt.Println("Greetings,"+userInput)
}

func main() {
  var userInput string
  fmt.Println("Enter your name")
  fmt.Scanln(&userInput)
  greeting(userInput)
  }