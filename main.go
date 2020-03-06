// Programmer name: Jacob Harter
// Date completed:  3/05/2020
// Description: tic tac toe, min max, name locator

package main

import (
    "fmt"
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
 //}    
    

fmt.Println(row)
   fmt.Println(row2)
    fmt.Println(row3)
   }
   }
  



 
 
