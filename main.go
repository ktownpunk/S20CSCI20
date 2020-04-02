package main

import(
"fmt"
)

func addNumbers(input1 int, input2 int){
 fmt.Println(input1, "+", input2, "=", input1+input2)
}

func subNumbers(input1 int, input2 int){
 fmt.Println(input1, "-", input2, "=", input1-input2)
}

func divNumbers(input1 int, input2 int){
  fmt.Println(input1, "/",input2,"=",input1/input2)
}
func mulNumbers(input1 int, input2 int){
  fmt.Println(input1,"*",input2,"=",input1*input2)
}
func main() {
var num1, num2 int
var op string

 fmt.Println("Please enter a whole number between 1 and 10")
 fmt.Scan(&num1)
 fmt.Println("Please enter another whole number between 1 and 10")
 fmt.Scan(&num2)

 fmt.Println("Please enter a mathematical operator (+,-,/,*)")
 fmt.Scan(&op)

if(op == "+"){
 addNumbers(num1,num2)
 } else if(op == "-"){
 subNumbers(num1,num2)
 }else if (op == "/"){
 divNumbers(num1,num2)
 }else if (op == "*"){
 mulNumbers(num1,num2)
 }else{
     fmt.Println("Invalid Operator")
   }
 }
