package main

import "fmt"

func main() {
	todoList := []string{
		"купить хлеб",
		"купить молоко",
		"купить чипсы",
		"купить пиво",
	}

	rotateList(todoList[:])
	fmt.Println(todoList)

}

func rotateList(arr []string) {
	copiedArr := make([]string, len(arr))
	copy(copiedArr, arr)
	for i := 0; i < len(copiedArr); i++ {
		arr[i] = copiedArr[len(copiedArr)-i-1]
	}
}
