package main

import(
"fmt"
"os"
//"bufio"
"strconv"
)
 

func damageTest(){
  var userInput int
  var gName string
 Stats,err:=os.Create("Stats.txt")
  if err != nil{
    panic(err)
  }
    fmt.Println("What weapon is being tested?")
    fmt.Scanln(&gName)
    Stats.WriteString(gName)
    Stats.WriteString("\n")
    fmt.Println("Enter the damage for Headshot")
    fmt.Scanln(&userInput)
    Stats.WriteString("Headshot Damage:")
    Stats.WriteString(strconv.Itoa(userInput))
    Stats.WriteString("\n")
    fmt.Println("Enter the damage for Body/Arm")
    fmt.Scanln(&userInput)
    Stats.WriteString("Body/Arm Damage:")
    Stats.WriteString(strconv.Itoa(userInput))
    Stats.WriteString("\n")
    fmt.Println("Enter the damage for Leg")
    fmt.Scanln(&userInput)
    Stats.WriteString("Leg Damage:")
    Stats.WriteString(strconv.Itoa(userInput))
    Stats.WriteString("\n")
    Stats.WriteString("\n")
}
func main() {
  var YN string
  //var gNumber int
  Stats, err := os.Create("Stats.txt")
  if err != nil{
    panic(err)
  }
  fmt.Println("Enter Y to continue or N to stop")
  fmt.Scanln(&YN)
  for YN == "Y"{
    damageTest()
    fmt.Println("Enter Y to continue or N to stop")
    fmt.Scanln(&YN)
  }
  
  Stats.Sync()
  Stats.Close()
}