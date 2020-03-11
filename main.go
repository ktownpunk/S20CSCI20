// Programmer name: Jacob Harter
// Date completed:  3/05/2020
// Description: Pig

package main

import (
    "fmt"
    "time"
    "math/rand"
)
func main() {
 var computerScore int //computer's score
 var userScore int //user's score
 var turnCounter =99 //monitors turns
 var dice int //stores dice max value
 var currentRoll int //short term value hold
 var pot int //point pot
 var answer string //used for user's choice
 var choice int //used for computer's choice
 seed := rand.NewSource(time.Now().UnixNano()) //seed for random#
 roll := rand.New(seed) //roll is the random#
 fmt.Println("What is the max value of the dice")
 fmt.Scanln(&dice)
 for userScore<100 && computerScore<100{
   switch turnCounter{
     case 99: //user's turn
      currentRoll =(roll.Intn(dice)) +1 //roll
      fmt.Println(currentRoll) //print the roll
      if currentRoll == 1 { 
        fmt.Println("Your turn is over")
        turnCounter++ //pass turn
        }else if currentRoll!=1{
      pot = pot + currentRoll //add roll to pot
      fmt.Println("Type roll to roll again, type anything else to end") //ask if User wants to roll again
      fmt.Scanln(&answer)
        if answer == "roll"{
      pot = pot + currentRoll
        }else{
        userScore = userScore + pot
        fmt.Println("Your score is ",userScore)
        pot = 0
        turnCounter++
      }

        }
        case 100: //computer's turn
        currentRoll =(roll.Intn(dice)) +1
      fmt.Println(currentRoll)
      if currentRoll == 1{
        fmt.Println("The computer's turn is over")
        turnCounter--
        }else if currentRoll != 1{
          pot = pot + currentRoll
          choice = (roll.Intn(6))
          if choice != 5{
      pot = pot + currentRoll
          }else{
          computerScore = computerScore + pot
        fmt.Println("Computer score is ",computerScore)
        pot = 0
        turnCounter--
          }
          
        }
      }
}
if computerScore >= 100{
fmt.Println("Computer Wins")
}else if userScore >=100{
  fmt.Println("You Win!")
}
}