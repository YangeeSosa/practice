package main

import (
	shapes "figures/figures"
	"fmt"
)

func main() {
	var choose int
	var r, b, h int
chooseOption:
	fmt.Println("Выберите, что нужно найти:\n 1) площадь круга \n 2) площадь треугольника")
	fmt.Scan(&choose)
	switch choose {
	case 1:
		fmt.Printf("введите значение радиуса круга: ")
		fmt.Scan(&r)
		circle := shapes.NewCircle(r)
		area, err := circle.GetArea()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Площадь круга: ", area)
	case 2:
		fmt.Printf("введите значение основания и высоты треугольника: ")
		fmt.Scan(&b, &h)
		triangle := shapes.NewTriangle(b, h)
		tArea, err := triangle.GetArea()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Площадь треугольника", tArea)
	default:
		fmt.Printf("Выберите из существующих вариантов!!!\n")
		goto chooseOption
	}
	shapes.SpawnFigures(&shapes.Triangle{Base: 3, Height: 7})
}
