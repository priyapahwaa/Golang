package main

import (
    "fmt"
    "math/rand"
)

type Student struct {
	name string
	age  int
	cgpa float64
}

	func cal (a int, b int, op func(int, int) int) int {
    return op(a, b)
}
func main() {
    fmt.Println("Hello")
	fmt.Printf("Random number: %d\n", rand.Intn(100))
	sum , diff := addOrsub(10, 5)
	fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)

	//variable 
	i:= 42
	fmt.Printf("The value of i is: %d\n", i)
	i = 100
	fmt.Printf("The new value of i is: %d\n", i)
	fmt.Printf("The sum of numbers from 0 to 9 is: %d\n", loop())


	s1 := Student{
        name: "Ayush Tibrewal hello",
        age:  21,
        cgpa: 8.05,
    }

    fmt.Println(s1.name)
    fmt.Println(s1.age)
    fmt.Println(s1.cgpa)

	// ananymous function
	ayush := func(name string , age int){
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
	ayush("Ayush" , 12)



// funtion inside the function same as ts
    add := func(x int, y int) int {
        return x + y
    }

    result := calculate(10, 5, add)
    fmt.Println(result) // 15
}



func addOrsub(a int , b int) (int, int){
	return a + b, a - b
}

func calculate(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func loop() (int){
	sum:=0
	
	for i:=0 ;i<10 ;i++{
		if(i %2==0) {
			break
		}
		sum += i
	}
	return sum
}
