// Programmer name: Jacob Harter
// Date completed:  3/29/20
// Description:MPG,tip,average,and comission

package main

import "fmt"
//import "math/rand"
//import "time"
func MPG(miles float64,gallons float64)float64{
return miles/gallons
}
func tip(spent float64,percent float64)float64{
return spent*percent/100
}

func comission(sold float64) float64{
  if sold < 50000{
    sold=sold*4/100
  }else if sold >=50000 && sold < 250000{
 sold=sold*6/100
  }else{
    sold=sold*6.75/100
  }
  return sold
}
func average() float64{
  userInput:= "Y"
  var userNumber float64
  var average float64
  var counter float64
  var total float64
for userInput == "Y"{
  fmt.Println("Enter a number")
  fmt.Scanln(&userNumber)
  total=userNumber+total
  counter++
  fmt.Println("total:",total)
  fmt.Println("count:",counter)
  fmt.Println("Enter Y to continue")
  fmt.Scanln(&userInput)
}
average= (total/counter)
return average
}

func main() {
  var YN = "Y"
  var userInput string
  for YN == "Y"{
  fmt.Println("What would you like to do? MPG, average, tip, or comission (case sensitive)")
  fmt.Scanln(&userInput)
  if userInput == "MPG"{
    var miles float64
    var gallons float64
    fmt.Println("Enter the miles driven")
    fmt.Scanln(&miles)
    fmt.Println("Enter the gallons used")
    fmt.Scanln(&gallons)
MPG:=MPG(miles,gallons)
fmt.Println("Your MPG is",MPG)
  }else if userInput =="tip"{
    var paid float64
    var percent float64
    fmt.Println("Enter the amount due")
    fmt.Scanln(&paid)
    fmt.Println("Enter the % you'd like to tip")
    fmt.Scanln(&percent)
    tip:=tip(paid,percent)
  fmt.Println("Your amount with the tip is",paid+tip)
  }else if userInput =="comission"{
    var sold float64
    fmt.Println("How much did you sell?")
    fmt.Scanln(&sold)
c:=comission(sold)
fmt.Println("Your comission is",c)
  }else if userInput =="average"{
    fmt.Println("Your average is",average())
  }
  fmt.Println("Would you like to do another? Y or N")
  fmt.Scanln(&YN)
  }
}