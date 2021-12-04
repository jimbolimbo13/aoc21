package utils

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func Day1() {
	file, err := os.Open("./data/day1_data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new, _ := strconv.Atoi(scanner.Text())
		data =  append(data, new)
	}

	// data = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	increases := 0

	for i := 4; i <= len(data); i++ {
		first := sum(data[i-4:i-1])
		// fmt.Println(data[i-4:i-1]) //
		second := sum(data[i-3:i])
		// fmt.Println("first", first, "second", second)
		if second > first {
			increases += 1
		}
	}

	fmt.Println("increases", increases)

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
