package main

import(
"fmt"
"os"
//"bufio"
"strconv"

)
func main() {
  var userInput int
  var YN string
  var gNumber int
  Stats, err := os.Create("Stats.txt")
  if err != nil{
    panic(err)
  }
  fmt.Println("Enter Y to continue or N to Stop")
  fmt.Scanln(&YN)
  for YN == "Y"{
    Stats.WriteString(strconv.Itoa(gNumber))
    Stats.WriteString("/n")
    fmt.Println("Enter the damage for this test")
    fmt.Scanln(&userInput)
    Stats.WriteString("/n")
    fmt.Println("Enter Y to continue, N to stop")
    fmt.Scanln(&YN)
  }
  
  Stats.Sync()
  Stats.Close()
}
