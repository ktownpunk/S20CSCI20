// Programmer name: Jacob Harter
// Date completed:  2/11/20
// Description: calculate inputed distance of rods to feet miles and meters, then display.

package main

import "fmt" 

func main() {
var input string 
fmt.Println("Enter your temperature scale F, C or K")
fmt.Scanln(&input)
if input == "F" {
  var fahrenheit float32
  fmt.Println("Enter your temperature")
  fmt.Scanln(&fahrenheit)
  var celsius = (fahrenheit-32)*5/9.0
  var kelvin float32
  kelvin= (((fahrenheit - 32) * 5/9.0) + 273.15)
  fmt.Println("Your temperature in Fahrenheit is",fahrenheit,",in Celsius it is",celsius,"and in Kelvin it is",kelvin)
 }else if input == "C"{
   var celsius float32
   fmt.Println("Enter your temperature")
  fmt.Scanln(&celsius)
  var fahrenheit = (celsius *9/5.0) + 32
  var kelvin = celsius + 273.15
  fmt.Println("Your temperature in Fahrenheit is",fahrenheit,",in Celsius it is",celsius,"and in Kelvin it is",kelvin)
 }else if input == "K"{
   var kelvin float32
   fmt.Println("Enter your temperature")
  fmt.Scanln(&kelvin)
  var celsius = kelvin - 273.15
  var fahrenheit = (kelvin - 273.15) *9/5.0 + 32
  fmt.Println("Your temperature in Fahrenheit is",fahrenheit,",in Celsius it is",celsius,"and in Kelvin it is",kelvin)
 }
}
