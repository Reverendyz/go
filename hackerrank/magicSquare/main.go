package main

import (
	"fmt"
	"math"
)

type Element struct {
	value       int
	isContained bool
}

func calcCost(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func contains(slice []int, v int) bool {
	for _, value := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func getNotIncluded(slice []int) []int {
	notIncluded := []int{}
	for i := 1; i < 10; i++ {
		if !contains(slice, i) {
			notIncluded = append(notIncluded, i)
		}
	}
	return notIncluded
}

func dup_count(list []int) []int {
	valuesInArray := make(map[int]int)
	duplicatedValues := []int{}
	for _, item := range list {
		_, exist := valuesInArray[item]
		if exist {
			duplicatedValues = append(duplicatedValues, item)
		} else {
			valuesInArray[item] = 1
		}
	}
	return duplicatedValues
}

func main() {
	var array []int
	var number int
	for i := 0; i < 9; i++ {
		fmt.Scanf("%d", &number)
		array = append(array, number)

	}
	notIncluded := getNotIncluded(array)
	duplicates := dup_count(array)

	fmt.Println(notIncluded)
	fmt.Println(duplicates)
	for i := 1; i < 10; i++ {

	}
}
