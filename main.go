// Programmer name: Jacob Harter
// Date completed:  3/29/20
// Description:Random feedback

package main

import "fmt"
import "math/rand"
import "time"
func Correct(){ //correct function, makes a random numebr which the switch outputs a random phrase
seed := rand.NewSource(time.Now().UnixNano())
random := rand.New(seed)
var feedback = (random.Intn(3))
  switch feedback{
    case 0:fmt.Println("Very Good!")
    case 1:fmt.Println("Nice Job!")
    case 2:fmt.Println("Keep it up!")
  }
  

}
func Incorrect(){//incorrect funtion, random number, switch same as the previous function
seed2 := rand.NewSource(time.Now().UnixNano())
random2 := rand.New(seed2)
var feedback = (random2.Intn(3))
  switch feedback{
    case 0:fmt.Println("Try Again")
    case 1:fmt.Println("Don't give up!")
    case 2:fmt.Println("Give it another shot!")
  }

}

func main() {//asks user the question and displays the choices to answer, then prompts input
  var UserInput string
  fmt.Println("What is a variable")
  fmt.Println("A. A reserved place in memory which can store one value")
  fmt.Println("B. A named number")
  fmt.Println("C. A math term")
  fmt.Println("D. A series of characters")
  fmt.Println("Enter your answer as a capital single character A B C or D")
  fmt.Scanln(&UserInput)
  if UserInput == "A" { //checks if the input is true or false
  Correct() //if true, calls on the correct function
  }else{
  Incorrect()//if false calls on the incorrect function
  }
}