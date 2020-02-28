// Programmer name: Jacob Harter
// Date completed:  2/25/2020
// Description: Rock Paper Scissors w/ computer

package main

import (
    "fmt"
)

func main() {
 var grossIncome float32 //income variable
 var status int //status variable
fmt.Println("Enter your gross income") //prompt input for income
fmt.Scanln(&grossIncome)
fmt.Println("Enter your status as a integer 1 for single, 2 for married filing jointed, 3 for head of household, 4 for married filing seperately")//prompt input for status
fmt.Scanln(&status)

switch status{
  case 1: { //tax brackets for single
if grossIncome <= 9700{
  fmt.Println("You are in the 10% bracket, you owe",grossIncome*.1)
}else if grossIncome <= 39475{
  fmt.Println("You are in the 12% bracket, you owe",grossIncome*.12)
}else if grossIncome <= 84200{
  fmt.Println("You are in the 22% bracket, you owe",grossIncome*.22)
}else if grossIncome <= 160725{
  fmt.Println("You are in the 24% bracket, you owe",grossIncome*.24)
}else if grossIncome <= 204100{
  fmt.Println("You are in the 32% bracket, you owe",grossIncome*.32)
}else if grossIncome <= 510300{
  fmt.Println("You are in the 35% bracket, you owe",grossIncome*.35)
}else{
  fmt.Println("You are in the 37% bracket, you owe",grossIncome*.37) 
}
}
case 2: //tax brackets for married filing jointed
  if grossIncome <= 19400{
  fmt.Println("You are in the 10% bracket, you owe",grossIncome*.1)
}else if grossIncome <= 78950{
  fmt.Println("You are in the 12% bracket, you owe",grossIncome*.12)
}else if grossIncome <= 168400{
  fmt.Println("You are in the 22% bracket, you owe",grossIncome*.22)
}else if grossIncome <= 321450{
  fmt.Println("You are in the 24% bracket, you owe",grossIncome*.24)
}else if grossIncome <= 408200{
  fmt.Println("You are in the 32% bracket, you owe",grossIncome*.32)
}else if grossIncome <= 612350{
  fmt.Println("You are in the 35% bracket, you owe",grossIncome*.35)
}else{
  fmt.Println("You are in the 37% bracket, you owe",grossIncome*.37) 
}

  case 3: //tax brackets for head of household
  if grossIncome <= 13850{
  fmt.Println("You are in the 10% bracket, you owe",grossIncome*.1)
}else if grossIncome <= 52850{
  fmt.Println("You are in the 12% bracket, you owe",grossIncome*.12)
}else if grossIncome <= 84200{
  fmt.Println("You are in the 22% bracket, you owe",grossIncome*.22)
}else if grossIncome <= 160700{
  fmt.Println("You are in the 24% bracket, you owe",grossIncome*.24)
}else if grossIncome <= 204100{
  fmt.Println("You are in the 32% bracket, you owe",grossIncome*.32)
}else if grossIncome <= 510300{
  fmt.Println("You are in the 35% bracket, you owe",grossIncome*.35)
}else{
  fmt.Println("You are in the 37% bracket, you owe",grossIncome*.37) 
}
  case 4: //tax brackets for married filing seperately
  if grossIncome <= 9700{
  fmt.Println("You are in the 10% bracket, you owe",grossIncome*.1)
}else if grossIncome <= 39475{
  fmt.Println("You are in the 12% bracket, you owe",grossIncome*.12)
}else if grossIncome <= 84200{
  fmt.Println("You are in the 22% bracket, you owe",grossIncome*.22)
}else if grossIncome <= 160725{
  fmt.Println("You are in the 24% bracket, you owe",grossIncome*.24)
}else if grossIncome <= 204100{
  fmt.Println("You are in the 32% bracket, you owe",grossIncome*.32)
}else if grossIncome <= 306175{
  fmt.Println("You are in the 35% bracket, you owe",grossIncome*.35)
}else{
  fmt.Println("You are in the 37% bracket, you owe",grossIncome*.37) 
}

}
}