package main

import "fmt"

func main() {
	fmt.Println("Hey Viz")

	var age int = 27
	id := 9

	var items [3]int
	trips := [4]int{2,3,4,5} 
	jobs := []int{1,1}

	fmt.Println(id)
	fmt.Println(jobs)

	if age >= 27 {
		items[1] = age
		fmt.Println("Age is", age)
		jobs = append(jobs, age)
	}

	fmt.Println("Jobs", jobs)

	fmt.Println("Items", items)

	fmt.Println("Trips", trips)
}