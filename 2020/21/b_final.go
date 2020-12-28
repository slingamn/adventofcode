package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func (s StringSet) Peek() string {
	for x := range s {
		return x
	}
	return ""
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

	allergenToIngredient := make(map[string]string)
	resolvedIngredients := make(StringSet)
	for {
		progress := false
		for allergen, candidates := range allergenToCandidates {
			if len(candidates) == 0 {
				panic(allergen)
			} else if len(candidates) == 1 {
				candidate := candidates.Peek()
				allergenToIngredient[allergen] = candidate
				delete(allergenToCandidates, allergen)
				resolvedIngredients.Add(candidate)
				progress = true
			} else {
				for candidate := range candidates {
					if resolvedIngredients.Has(candidate) {
						delete(candidates, candidate)
						progress = true
					}
				}
			}
		}
		if len(allergenToCandidates) == 0 {
			break
		}
		if !progress {
			panic("failed to make progress")
		}
	}

	allergens := make([]string, 0, len(allergenToIngredient))
	for allergen := range allergenToIngredient {
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)
	ingredients := make([]string, 0, len(allergens))
	for _, allergen := range allergens {
		ingredients = append(ingredients, allergenToIngredient[allergen])
	}
	fmt.Println(strings.Join(ingredients, ","))

	return 0, nil
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
