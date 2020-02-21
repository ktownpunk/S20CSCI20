// Programmer name: Jacob Harter
// Date completed:  2/11/20
// Description: calculate average of test scores, and make a pascls triangle with a user inputed height.
package main
import "fmt"
 
func main(){
    var rows int
    var number int = 1
    fmt.Print("Enter number of rows : ") //prompt user to input height
    fmt.Scan(&rows) // input heigh
    for i := 0; i < rows; i++ {  //loop to be used in formula
        for k := 0; k <= i; k++ { //loop to be used in formula
 
            if (k==0 || i==0) { //if the loops equal 0 than the number is 1, the sides/first middle
                    number = 1
                }else{
                    number = number*(i-k+1)/k //formula for pascals triangle
                }
 
            fmt.Print(number) //print the numbers        
        }
        fmt.Println(" ") //space the levels
    }
}