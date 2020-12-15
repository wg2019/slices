package slices

import "fmt"

// RmDInt 整型切片去重.
func RmDInt(slice []int) []int {
	intMap := make(map[int]bool)
	newSlice := make([]int, 0)
	for _, value := range slice {
		if _, ok := intMap[value]; !ok {
			newSlice = append(newSlice, value)
		}
		intMap[value] = true
	}
	return newSlice
}

// RmDString 字符串切片去重.
func RmDString(slice []string) []string {
	intMap := make(map[string]bool)
	newSlice := make([]string, 0)
	for _, value := range slice {
		if _, ok := intMap[value]; !ok {
			newSlice = append(newSlice, value)
		}
		intMap[value] = true
		fmt.Printf("value: %+v", value)
	}
	return slice
}
