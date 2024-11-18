package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Привет Интернет")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Query().Get("message"))
	})

	http.HandleFunc("/area", func(w http.ResponseWriter, r *http.Request) {
		radiusStr := r.URL.Query().Get("radius")
		radius, err := strconv.ParseFloat(radiusStr, 64)
		if err != nil || radius <= 0 {
			err := errors.New("Некоректные данные")
			fmt.Fprint(w, err)
		} else {
			area := math.Pi * radius * radius
			fmt.Fprint(w, area)
		}

	})

	http.ListenAndServe(":80", nil)
}
