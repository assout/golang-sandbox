package main

import (
	"fmt"
	"time"
)

func main2() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	sub()
}

const Pi = 3.14

func sub() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
