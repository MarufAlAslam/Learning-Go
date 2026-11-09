package main

import "fmt"

/*
Write exercise.go that prints:

Name: <your name>
Goal: Learn Go in 60 days
Today is day 1
*/
func variables() {
	myName := "Maruf"
	myGoal := "Learn Go in 60 days"
	dayNum := 1
	fmt.Printf("\nName: %s\n", myName)
	fmt.Printf("Goal: %s\n", myGoal)
	fmt.Printf("Today is day %d\n", dayNum)
}
