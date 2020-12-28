package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bytes"
	"strconv"
)

func solve(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] + nums[j] == 2020 {
				return nums[i] * nums[j]
			}
		}
	}
	return 0
}

func main() {
	s, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fields := bytes.Fields(s)
	nums := make([]int, 0, len(fields))
	for _, field := range fields {
		if len(field) != 0 {
			num, err := strconv.Atoi(string(field))
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
	}
	fmt.Println(solve(nums))
}
