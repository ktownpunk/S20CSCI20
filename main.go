// Programmer Name: Jacob Harter
// Date Completed: 1/30/2020
// Description: Calculate the poulation gain of the U.S. with current population, net gain, and time elapsed

package main

import "fmt"

func main() {
//
var Population = 329252103 //
var NetGain = 24 // 
var TimeY = 10 // 
var Time = TimeY*365*24*60*60 // 
var PopulationIncrease = NetGain*Time
fmt.Println("Current Population of the U.S. is", Population) //
fmt.Print("The rate of growth is ", NetGain, " people per second") //
fmt.Println(" over", Time, "seconds") //
fmt.Print("For an increase of ", PopulationIncrease) //
fmt.Println(" population in",Time,"seconds or",TimeY,"years")// 
fmt.Println("The Population of the U.S. in",Time,"seconds or",TimeY, "years will be", Population+PopulationIncrease,"people") // 

//Description: create string variables that hold names, int variables that hold date, put them together for today's date
var DayOfTheWeek = "Tuesday" //
var Month = "February" //
var Date = 204 // 
var Year = 2020 //
fmt.Print("Today is 0") //
fmt.Println(Date,"a",DayOfTheWeek,"in",Month,Year) //
}
