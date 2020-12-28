package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type empty struct{}

type StringSet map[string]empty

func (s StringSet) Add(str string) {
	s[str] = empty{}
}

func (s StringSet) Has(str string) bool {
	_, ok := s[str]
	return ok
}

func (s StringSet) Copy() (c StringSet) {
	c = make(StringSet, len(s))
	for x := range s {
		c.Add(x)
	}
	return
}

func (s StringSet) Intersection(o StringSet) {
	for x := range s {
		if !o.Has(x) {
			delete(s, x)
		}
	}
	return
}

func solve(input []string) (result int, err error) {
	allergenToCandidates := make(map[string]StringSet)
	allIngredients := make(StringSet)
	var originalRecipes []StringSet

	for _, line := range input {
		fields := strings.Fields(line[:len(line)-1])
		ingredients := make(StringSet)
		originalIngredients := make(StringSet)
		i := 0
		for ; i < len(fields); i++ {
			if fields[i][0] == '(' {
				break
			}
			ingredients.Add(fields[i])
			originalIngredients.Add(fields[i])
			allIngredients.Add(fields[i])
		}
		i++
		originalRecipes = append(originalRecipes, originalIngredients)
		for ; i < len(fields); i++ {
			allergen := strings.TrimSuffix(fields[i], ",")
			existing, found := allergenToCandidates[allergen]
			if !found {
				allergenToCandidates[allergen] = ingredients.Copy()
			} else {
				existing.Intersection(ingredients)
			}
		}
	}

	var innocentIngredients []string
	for ingredient := range allIngredients {
		ok := true
		for _, candidates := range allergenToCandidates {
			if candidates.Has(ingredient) {
				ok = false
				break
			}
		}
		if ok {
			innocentIngredients = append(innocentIngredients, ingredient)
		}
	}

	for _, ingredient := range innocentIngredients {
		for _, recipe := range originalRecipes {
			if recipe.Has(ingredient) {
				result++
			}
		}
	}

	return
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	solution, err := solve(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}
