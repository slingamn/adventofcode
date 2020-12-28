package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"strings"
	"strconv"
)

var (
	ErrBadLine = errors.New("bad line")
)

type Bag struct {
	Color  string
	Holds  []*Bag
	Counts []int
}

func (bag *Bag) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s bags contain ", bag.Color)
	for i := 0; i < len(bag.Holds); i++ {
		fmt.Fprintf(&buf, "%d %s bags, ", bag.Counts[i], bag.Holds[i].Color)
	}
	return buf.String()
}

type empty struct{}

func reachable(allBags map[string]*Bag, startPos *Bag) (result map[*Bag]empty) {
	result = make(map[*Bag]empty)
	queue := []*Bag{startPos,}

	for len(queue) != 0 {
		bag := queue[0]
		queue = queue[1:]

		// visit
		result[bag] = empty{}

		for _, child := range bag.Holds {
			if _, ok := result[child]; !ok {
				queue = append(queue, child)
			}
		}
	}

	return
}

func readStdin() (allBags map[string]*Bag, err error) {
	allBags = make(map[string]*Bag)

	getBag := func(c string) (result *Bag) {
		result = allBags[c]
		if result == nil {
			result = &Bag {
				Color: c,
				//Holds: make([]*Bag, 0, 10),
				//Counts: make([]int, 0, 10),
			}
			allBags[c] = result
		}
		return result
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		bagIdx := strings.Index(line, " bags contain ")
		color := line[:bagIdx]
		bag := getBag(color)
		rest := line[bagIdx+len(" bags contain "):]
		var pieces []string
		if rest != "no other bags." {
			pieces = strings.Split(rest, ", ")
		}
		for _, piece := range pieces {
			spaceIdx := strings.IndexByte(piece, ' ')
			bagsIdx := strings.Index(piece, " bag")
			count, err := strconv.Atoi(piece[:spaceIdx])
			if err != nil {
				return nil, err
			}
			bagColor := piece[spaceIdx+1:bagsIdx]
			childBag := getBag(bagColor)
			bag.Holds = append(bag.Holds, childBag)
			bag.Counts = append(bag.Counts, count)
		}
		allBags[color] = bag
	}

	return
}

func main() {
	allBags, err := readStdin()
	if err != nil {
		panic(err)
	}

	shinyGold := allBags["shiny gold"]
	count := 0
	for _, bag := range allBags {
		if bag == shinyGold {
			continue
		}
		if _, ok := reachable(allBags, bag)[shinyGold]; ok {
			count++
		}
	}
	fmt.Println(count)
}
