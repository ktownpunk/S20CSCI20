// Programmer Name: Jacob Harter
// Date Completed: 1/30/2020
// Description: Calculate the poulation gain of the U.S. with current population, net gain, and time elapsed

package main

import "fmt"

func main() { 
//
var Population = 329252103 // population of U.S.
var NetGain = 24 // rate of growth in people/seconds
var TimeY = 10 // time in years
var Time = TimeY*365*24*60*60 // variable of time elapsed in seconds
var PopulationIncrease = NetGain*Time
fmt.Println("Current Population of the U.S. is", Population) //state current population
fmt.Print("The rate of growth is ", NetGain, " people per second") //state growth/time
fmt.Println(" over", Time, "seconds") //state time elapsed
fmt.Print("For an increase of ", PopulationIncrease) //state increase
fmt.Println(" population in",Time,"seconds or",TimeY,"years")// state increase and time in seconds and years
fmt.Println("The Population of the U.S. in",Time,"seconds or",TimeY, "years will be", Population+PopulationIncrease,"people") // put it all together

//Description: create string variables that hold names, int variables that hold date, put them together for today's date
var DayOfTheWeek = "Tuesday" //Day of Week
var Month = "February" //Month
var Date = 204 // MonthDayDay
var Year = 2020 //YearYearYearYear
fmt.Print("Today is 0") // 0 place-holder if month is 1-9
fmt.Println(Date,"a",DayOfTheWeek,"in",Month,Year) //Statement
}
