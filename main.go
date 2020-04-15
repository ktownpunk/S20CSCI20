package main

import(
"fmt"
)

func add(num1 int, num2 int) int{

return num1+num2
}

func subtract(num1 int, num2 int) int{
return num1-num2
}

func multiply(num1 int, num2 int) int{
  return num1*num2
}
//write a function for multiply
func divide(num1 int, num2 int) int{
return num1/num2
}
//write a function for division

func main() {
a := 0
b := 75;
a=multiply(b,2)
//double the b value
a=add(a,9)
//add 9 to the result
a=subtract(a,3)
//subtract 3 from the result
a=divide(a,2)
//divide the result by 2
a=subtract(a,b)
//subtract the original value b from the result
fmt.Println(a)
//output the answer
//the answer should be 3
}