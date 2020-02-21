// Programmer name: Jacob Harter
// Date completed:  2/11/20
// Description: calculate average of test scores, and make a pascls triangle with a user inputed height.

package main

import "fmt" 

func main() {
  var totalScore float32 = 0 //totalofScores
  var count float32 = 0 //amount of inputed scores, not including stop input
  var input float32 = 0 //value of the input
  
fmt.Println("Enter your Score") //prompt and input the score
fmt.Scanln(&input)
for input != -1 {
  totalScore = totalScore + input
fmt.Println("Enter your Score of -1 to STOP") //prompt another input, either a score or stop
fmt.Scanln(&input)
count++ //increase count
}
var averageofscores = totalScore/count //take the average of total scores
fmt.Println("The average is",averageofscores) //print the average
}