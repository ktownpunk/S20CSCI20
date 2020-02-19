// Programmer name: Jacob Harter
// Date completed:  2/18/2020
// Description: have the computer count to a user specified number

package main

import (
    "fmt"
    "math/rand"
  
) //adding the ability to do random numbers

func main() {
    var count = 1
    var userGuess int
    //create variable for count
   fmt.Println("Enter the max range")
    //ask the user to enter a max range for the guessing game and store that value in variable max.
    var max int
    fmt.Scanln(&max)
    fmt.Println(&userGuess)
    //this next line creates a random number from 1 to that guess for the computer to know.  You can test this by printing out the variable computerGuess
    var computerGuess = rand.Intn(max)

    //ask the user to enter a guess for the computer number
    fmt.Println("Enter your guess")
    fmt.Scanln(&userGuess)
    //create a loop that compares the computerGuess to the userGuess while they are NOT equal go into the loop
        //increase the count by 1
        //tell the user that the guessed incorrect
        //ask the user to enter a new guess for the computer number
    for userGuess!=computerGuess{
      count++
      fmt.Println("You guessed incorrect")
      fmt.Println("Try again")
      fmt.Scanln(&userGuess)
    }
    //print out that the user got the answer correctly and how many guesses it took (the count)
    fmt.Println("You guessed correctly, it took",count,"attempts")
}