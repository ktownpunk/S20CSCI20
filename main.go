// Programmer name: Jacob Harter
// Date completed:  3/29/20
// Description:Random feedback

package main

import "fmt"
//import "math/rand"
//import "time"

func age(bYear int,cYear int){
  fmt.Println("Your age is ",cYear-bYear)
}
func eggs(doz int){
fmt.Println(doz,"dozens are ",doz*12," eggs")
}
func swap(num1 int,num2 int){
fmt.Println("Your numbers swapped are",num2,"and",num1,"Magic!")
}
func gradeReport(score int){
  grade:= score/10
  switch grade{
    case 10:
      fmt.Println("The grade is an A+")
    case 9:
      fmt.Println("The grade is an A")
    case 8:
      fmt.Println("The grade is a B")
    case 7:
      fmt.Println("The grade is a C")
    default:
      fmt.Println("The grade is a Fail")
  }
}
func ask(){
  var userInput string
 fmt.Println("Enter which function you would like to use: Accepted choices are AGE, EGGS, SWAP, and GRADEREPORT")
  fmt.Scanln(&userInput)
  if userInput == "AGE"{
    var bYear int
    var cYear int
    fmt.Println("Enter your birth year in YYYY format, example 2003")
    fmt.Scanln(&bYear)
    fmt.Println("Enter the current year")
    fmt.Scanln(&cYear)
    age(bYear,cYear)
  }else if userInput == "EGGS"{
    var doz int
    fmt.Println("How many dozen's of eggs would you like to translate?")
    fmt.Scanln(&doz)
    eggs(doz)
}else if userInput == "SWAP"{
  var num1 int
  var num2 int
  fmt.Println("Enter a number")
  fmt.Scanln(&num1)
  fmt.Println("Enter another number")
  fmt.Scanln(&num2)
  swap(num1,num2)
}else if userInput == "GRADEREPORT"{
  var score int
  fmt.Println("Enter the score as a number between 0-100")
  fmt.Scanln(&score)
  gradeReport(score)
} else {
  fmt.Println("Sorry that isn't an option, try again")
  ask()
}
}
func main(){
  fmt.Println("What would you like to")
  ask()
}
