package main

import "fmt"

func main() {

	//Elements into an array
	var arr [7]int
	fmt.Println("Enter an unsorted array")
	for i := 0; i < 7; i++ {
		fmt.Scanln(&arr[i])
	}

	fmt.Println(arr)

	bubble(arr)
}

func bubble(arr [7]int) {
	//Bubble Sort
	swap := true
	for !swap {
		swap = false
		for i := 0; i < 6; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}

	fmt.Println("Bubble Sort Result:-")
	fmt.Println(arr)
}
