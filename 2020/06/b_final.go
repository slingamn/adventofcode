package main

import (
	"fmt"
	"bufio"
	"os"
)

func countSetBits(v uint32) (c int) {
	// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
	for c = 0; v != 0; c++ {
		v &= v - 1; // clear the least significant bit set
	}
	return
}

func readStdin() (sum int) {
	scanner := bufio.NewScanner(os.Stdin)
	first := true
	var groupForm uint32
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum += countSetBits(groupForm)
			first = true
			groupForm = 0
			continue
		}
		var form uint32
		for i := 0; i < len(line); i++ {
			form = form | (1 << (line[i] - 'a'))
		}
		if first {
			groupForm = form
		} else {
			groupForm = groupForm & form
		}
		first = false
	}
	sum += countSetBits(groupForm)

	return
}

func main() {
	result := readStdin()
	fmt.Println(result)
}
