package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	Radius int
}

type Triangle struct {
	Base   int
	Height int
}

type Figure interface {
	GetArea() (float32, error)
}

func NewTriangle(b int, h int) Triangle {
	return Triangle{
		Base:   b,
		Height: h,
	}
}

func (t Triangle) GetArea() (float32, error) {
	if t.Base <= 0 || t.Height <= 0 {
		return float32(0), errors.New("основание и высота должны быть больше 0")
	} else {
		return float32((t.Base * t.Height) / 2), nil
	}
}

func NewCircle(radius int) Circle {
	return Circle{
		Radius: radius,
	}
}

func (c Circle) GetArea() (float32, error) {
	if c.Radius <= 0 {
		return float32(0), errors.New("радиус должен быть больше 0")
	} else {
		return float32(c.Radius) * float32(c.Radius) * math.Pi, nil
	}
}

func SpawnFigures(f Figure) {
	switch fig := f.(type) {
	case Triangle:
		for i := 1; i <= 10; i++ {
			area, err := f.GetArea()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Площадь фигуры:", area)
			}
			fig.Base++
			fig.Height++
		}
	case Circle:
		for i := 1; i <= 10; i++ {
			area, err := f.GetArea()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Площадь фигуры:", area)
			}
			fig.Radius++
		}
	}
}
