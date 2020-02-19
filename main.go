// Programmer name: Jacob Harter
// Date completed:  2/11/20
// Description: calculate inputed distance of rods to feet miles and meters, then display.

package main

import "fmt" 

func main() {
var input string 
fmt.Println("Enter your temperature scale F, C or K") //prompt user to input their temp scale's initial
fmt.Scanln(&input) //scan and save the input
if input == "F" { //if it is Fahrenheit
  var fahrenheit float32 //make float variable fahrenheit
  fmt.Println("Enter your temperature") //prompt user to ask for their temperature
  fmt.Scanln(&fahrenheit) //scan and save the input
  var celsius = (fahrenheit-32)*5/9.0 //F to C
  var kelvin float32 // make kelvin a variable
  kelvin= (((fahrenheit - 32) * 5/9.0) + 273.15) //F to K
  fmt.Println("Your temperature in Fahrenheit is",fahrenheit,",in Celsius it is",celsius,"and in Kelvin it is",kelvin) //print the temp as a number with the corresponding scale
 }else if input == "C"{ //if C
   var celsius float32 //make float variable celsius
   fmt.Println("Enter your temperature") //prompt user to input the temp
  fmt.Scanln(&celsius) //scan and save the input
  var fahrenheit = (celsius *9/5.0) + 32 // C to F
  var kelvin = celsius + 273.15 // C to K
  fmt.Println("Your temperature in Fahrenheit is",fahrenheit,",in Celsius it is",celsius,"and in Kelvin it is",kelvin) //print the temp as a number with the corresponding scale
 }else if input == "K"{ //if K
   var kelvin float32 //make float variable kelvin
   fmt.Println("Enter your temperature") //prompt user to enter their temp
  fmt.Scanln(&kelvin) //scan and save input
  var celsius = kelvin - 273.15 // K to C
  var fahrenheit = (kelvin - 273.15) *9/5.0 + 32 //K to F
  fmt.Println("Your temperature in Fahrenheit is",fahrenheit,",in Celsius it is",celsius,"and in Kelvin it is",kelvin) //print temp as a number with its corresponding scale
 }
}
