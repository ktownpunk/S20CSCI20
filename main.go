// Programmer name: Jacob Harter
// Date completed:  2/25/2020
// Description: Rock Paper Scissors w/ computer

package main

import (
    "fmt"
    "math/rand"
    //"time"
) //adding the ability to do random numbers

func main() {
    var computerchoice = rand.Intn(2)
    var playerchoice int
    //create two variables - one for the computer and one for the user
    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 0=rock, 1=scissors, 2=paper
    fmt.Println("Enter your choice as 0 (rock) 1 (paper) 2 (scissor)")
    fmt.Scanln(&playerchoice)
    //prompt the user for an integer value representing the player's choice
    if computerchoice == 0{
      fmt.Println("Computer chooses rock")
    } else if computerchoice == 1{
      fmt.Println("Computer chooses paper")
    } else if computerchoice == 2{
      fmt.Println("Computer chooses scissors")
    }
    if playerchoice == 0{
      fmt.Println("You chose rock")
    } else if playerchoice == 1{
      fmt.Println("You chose paper")
    } else if playerchoice == 2{
      fmt.Println("You chose scissors")
    }
    
    //Print out the values using the words rock, scissors, paper.  ie. "Computer chose rock and player chose paper"
    //You will need to use decisions for this
}