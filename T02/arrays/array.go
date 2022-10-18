package main

import (
	"fmt"
	"math"
	"math/rand"
)

func genSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int()%200 - 100
	}
	return s
}

func secondMin(s []int) int {
	stmin, ndmin := math.MaxInt, math.MaxInt
	for _, v := range s {
		if v < stmin {
			ndmin = stmin
			stmin = v
		} else if v < ndmin && v != stmin {
			ndmin = v
		}
	}
	return ndmin
}

func uniqVal(s []int) []int {
	m := make(map[int]int)
	for _, v := range s {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	var u []int
	for k, v := range m {
		if v == 1 {
			u = append(u, k)
		}
	}
	return u
}

func reverseArray(s *[]int) []int {
	n := len(*s) - 1
	for i, _ := range *s {
		(*s)[i], (*s)[n-i] = (*s)[n-i], (*s)[i]
		if i == n/2 {
			break
		}
	}
	return *s
}

func delDouble(s *[]int) []int {
	m := make(map[int]bool)
	newS := []int{}
	for _, v := range *s {
		if !m[v] {
			newS = append(newS, v)
			m[v] = true
		}
	}
	s = &newS
	return newS
}

func numOf(s string) (int, int) {
	var lat, cyr int
	for _, v := range s {
		if 'А' <= int(v) && int(v) <= 'ї' {
			cyr++
		} else if 'A' < int(v) && int(v) < 'z' {
			lat++
		}
	}
	return lat, cyr
}

//Слайс можна псевдозгенерувати за допомогою функції genSlice() змінивши значення n на потрібне
//або можна закоментувати 89, 90 рядки й вписати свої значення у 91 рядку
//Рядок для обчислення літер можна змінити на 98 рядку
func main() { //go run main.go
	var n = 10
	array := genSlice(n)
	//array := []int{1, 1, 1, 1, 1, 1, 5, 6, 7, 8, 9, 9, 9, 9}
	fmt.Println("Generated slice: ", array)
	fmt.Println("Second min: ", secondMin(array))
	fmt.Println("Unique values: ", uniqVal(array))
	fmt.Println("Reversed slice: ", reverseArray(&array))
	fmt.Println("Array after doubles removed: ", delDouble(&array))

	str := "Надія є - Mad Heads"
	a, b := numOf(str)
	fmt.Println("\nString: ", str)
	fmt.Println("Number of Latin letters: ", a, "\nNumber of Cyrillic letters: ", b)
}
