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
	selection(arr)
}

func bubble(arr [7]int) {
	//Bubble Sort
	swap := true
	//count := 0
	for swap {
		//count++
		swap = false
		for i := 0; i < 6; i++ {
			//count++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}

	fmt.Println("Bubble Sort Result:-")
	fmt.Println(arr)
	//fmt.Println(count)
}

func selection(arr [7]int) {
	//Selection Sort
	//count := 0
	for i := 0; i < 6; i++ {
		//count++
		minIndex := i
		for j := i + 1; j < 7; j++ {
			//count++
			if arr[j] < arr[minIndex] {
				minIndex = j
			}

		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	fmt.Println("Selection Sort Result:-")
	fmt.Println(arr)
	//fmt.Println(count)
}
