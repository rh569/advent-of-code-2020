package main

import (
	"fmt"
	"sort"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   []string
}

func main() {
	fmt.Println("Part 1:")

	countNonAllergenIngredients, ingredientsByAllergen := part1(FoodInput)
	fmt.Printf("Non-allergen containing ingredients: %v\n", countNonAllergenIngredients)

	fmt.Println("Part 2:")
	fmt.Printf("Canonical dangerous ingredient list: %v\n", part2(ingredientsByAllergen))
}

func parseInput(input string) []Food {
	lines := strings.Split(input, "\n")
	foods := []Food{}

	for _, l := range lines {
		parts := strings.Split(l, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		parts[1] = strings.Replace(parts[1], ")", "", 1)
		allergens := strings.Split(parts[1], ", ")

		foods = append(foods, Food{ingredients, allergens})
	}

	return foods
}

func part1(input string) (int, map[string][]string) {
	foods := parseInput(input)

	ingredientsByAllergen := getIngredientsByAllergen(foods)

	countNonAllergenIngredients := 0
	allergenIngredients := []string{}

	for k, v := range ingredientsByAllergen {
		fmt.Printf("Allergen: %v | Ingredients: %v\n", k, v)
		allergenIngredients = append(allergenIngredients, v[0])
	}

	for _, f := range foods {
		for _, i := range f.ingredients {

			if !contains(allergenIngredients, i) {
				countNonAllergenIngredients++
			}
		}
	}

	return countNonAllergenIngredients, ingredientsByAllergen
}

func part2(dangerousIngredientsByAllergen map[string][]string) string {

	alphaAllergens := []string{}

	for k, _ := range dangerousIngredientsByAllergen {
		alphaAllergens = append(alphaAllergens, k)
	}

	sort.Strings(alphaAllergens)

	CDIL := ""

	for i, a := range alphaAllergens {
		CDIL += dangerousIngredientsByAllergen[a][0]

		if i < len(alphaAllergens)-1 {
			CDIL += ","
		}
	}

	return CDIL
}

// returns a map of allergens, each pointing to an array of ingredients that may contain them
func getIngredientsByAllergen(foods []Food) map[string][]string {
	ingredientsByAllergen := make(map[string][]string)

	for _, f := range foods {

		for _, a := range f.allergens {

			ingredients, present := ingredientsByAllergen[a]

			if present {
				ingredientsByAllergen[a] = resolveMatchingIngredients(ingredients, f.ingredients)
			} else {
				ingredientsByAllergen[a] = f.ingredients
			}
		}
	}

	loops := 0

	// Compare ingredients by allergen to each other to narrow down
	// This might be completely unnecessary for part 1...
	for !isOneToOne(ingredientsByAllergen) && loops < len(ingredientsByAllergen) {
		for _, oV := range ingredientsByAllergen {
			// k allergen must be in v[0] ingredient only, remove v[0] ingredient from other allergens
			if len(oV) == 1 {
				ingredientToRemove := oV[0]

				for iK, iV := range ingredientsByAllergen {

					if len(iV) > 1 {
						ingredientsByAllergen[iK] = removeIngredient(iV, ingredientToRemove)
					}
				}
			}
		}
		loops++
	}

	return ingredientsByAllergen
}

// returns true if every value has only one element
func isOneToOne(m map[string][]string) bool {
	isOneToOne := true

	for _, v := range m {
		if len(v) > 1 {
			isOneToOne = false
		}
	}

	return isOneToOne
}

// removes the given element from the list, returning the new list
// returns the given list unchanged if no such element found
func removeIngredient(list []string, element string) []string {

	for i, e := range list {
		if element == e {

			if i == 0 {
				return list[1:]
			}

			if i == len(list)-1 {
				return list[:len(list)-1]
			}

			return append(list[:i], list[i+1:]...)
		}
	}

	return list
}

// returns only the elements that exist in both lists
func resolveMatchingIngredients(list1, list2 []string) []string {
	listOut := []string{}

	for _, ingredient1 := range list1 {
		inBoth := false

		for _, ingredient2 := range list2 {
			if ingredient1 == ingredient2 {
				inBoth = true
			}
		}

		if inBoth {
			listOut = append(listOut, ingredient1)
		}
	}

	return listOut
}

func contains(list []string, element string) bool {
	for _, e := range list {
		if element == e {
			return true
		}
	}

	return false
}
