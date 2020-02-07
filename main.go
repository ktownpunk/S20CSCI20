// Programmer name: Jacob Harter
// Date completed:  2/6/20
// Description: variables for my favourite, name, age, prompt to enter them, output them.

package main

import "fmt" 

func main() {
    //declare variable for favorite food and store your favorite food.
    var favFood = "a nice sandwich"
    //declare variables for name and age (make sure they are appropriate data types)
    var name string
    var age string
    //ask user to enter name and age
    fmt.Scanln(&name)
    fmt.Scanln(&age)
    //output a welcome statement using their name
    fmt.Println("Greetings and well-met," + name + ",fellow seeker of truth")
    //output a statement that says At their age you enjoyed the favorite food
    fmt.Println("When I was",age,"I enjoyed",favFood)
}