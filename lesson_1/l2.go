package main

import "fmt"

// Дан целочисленный массив nums, выведите на экран true, если любое значение
// встречается в массиве как минимум дважды, и false, если каждый элемент
// различен.

func main() {
	nums := [...]int32{5, 2, 3, 5, 5}

	count := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == nums[i] {
				count++
			}
		}
		if count >= 2 {
			fmt.Println(true)
			return
		} else {
			count = 0
		}
	}
}
