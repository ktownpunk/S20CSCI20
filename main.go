// Programmer name: Jacob Harter
// Date completed:  3/05/2020
// Description: tic tac toe, min max, name locator

package main

import (
    "fmt"
    "time"
    "math/rand"
)
func main() {
 var column [3] string
 column[0] = "O"
 column[1] = "O"
 column[2] = "O"
 var row [3] string
 row[0] = "O"
 row[1] = "O"
 row[2] = "O"
  var row2 [3] string
 row2[0] = "O"
 row2[1] = "O"
 row2[2] = "O"
  var row3 [3] string
 row3[0] = "O"
 row3[1] = "O"
 row3[2] = "O"
 
 var userColumn int
 var userRow int
 for userRow != 3 {
  fmt.Println("Enter your row as 0-2, or 3 to STOP")
    fmt.Scanln(&userRow)
    if userRow == 0 {
      fmt.Println("Enter your column as 0-2")
        fmt.Scanln(&userColumn)
          if row[userColumn] == "O" {}
            row[userColumn] = "X"
    
  
    }else if userRow == 1 {
        fmt.Println("Enter your column as 0-2")
        fmt.Scanln(&userColumn)
        if row2[userColumn] == "O" {
            row2[userColumn] = "X"
          }
          }else if userRow == 2 {
              fmt.Println("Enter your column as 0-2")
        fmt.Scanln(&userColumn)
        if row3[userColumn] == "O" {
            row3[userColumn] = "X"
          }
          }
 
    

fmt.Println(row)
   fmt.Println(row2)
    fmt.Println(row3)
   }
   //ID lookup by name
   var ID [5] string
   var userName string
   var counter int
   var IDFound = false
   ID[0] = "John"
   ID[1] = "James"
   ID[2] = "Luke"
   ID[3] = "Sam"
   ID[4] = "Jessie"
fmt.Println("Enter Name (Case Sensitive)")
fmt.Scanln(&userName)
for counter <6 {
  if userName == ID[counter] {
  fmt.Println("Your ID is", counter)
  IDFound = true
  counter++
  }else if IDFound == false{
    fmt.Println("The name you entered is not in our system")
    break
  }
}
seed := rand.NewSource(time.Now().UnixNano())
    randomNumber := rand.New(seed)
var numbers = [5]int{(randomNumber.Intn(5))}
	min, max := findMinAndMax(numbers)
	fmt.Println("Minimum: ", min) //prints the min value
	fmt.Println("Maximum: ", max) //print the max value
  fmt.Println(numbers)
}
func findMinAndMax(numbers [5]int) (min int, max int) { //function to find the min and max from the array "numbers"
	min = numbers[0] //sets min
	max = numbers[0] //sets max
	for _, value := range numbers { //for each value compare the number to value
		if value < min { //checks the min to another value
			min = value //sets value to the smaller one
		}
		if value > max { //checks the max to another value
			max = value //sets value to the greater one
		}
	}
	return min, max //gives the minimum and maximum values
}


 
 
