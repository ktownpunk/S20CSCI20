// Programmer name: Jacob Harter
// Date completed:  2/11/20
// Description: calculate inputed distance of rods to feet miles and meters, then display.

package main

import "fmt" 

func main() {
var rods float32
  fmt.Println("Enter your distance in rods as a decimal") //tell user to input their distance in rods as a decimal
  fmt.Scanln(&rods) //prompt user to input
  var meter float32 = rods * 5.0292
 var feet float32 = rods * 16.5
 var miles float32= rods * 0.003125
  fmt.Println("Your distance is",meter,"meters",feet,"feet",miles,"miles")

  
}