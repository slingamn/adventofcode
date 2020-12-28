package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"sort"
	"strconv"
	"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)


func readStdin() (est int, buses []int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	estStr := scanner.Text()
	est, err = strconv.Atoi(estStr)
	if err != nil {
		return
	}
	scanner.Scan()
	busStr := scanner.Text()
	busVals := strings.Split(busStr, ",")
	var bus int
	for _, busVal := range busVals {
		bus, err = strconv.Atoi(busVal)
		if err != nil {
			continue
		}
		buses = append(buses, bus)
	}

	return
}

func solve(est int, buses []int) (result int) {
	for i := est; true; i++ {
		for _, bus := range buses {
			if i % bus == 0 {
				return (i - est) * bus
			}
		}
	}
	return
}

func main() {
	est, buses, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(est, buses))
}
