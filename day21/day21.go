package day21

import (
	"sort"
	"strings"
)

type IngredientLabel struct {
	Ingredients []string
	Allergens   []string
}

func UniqueStrings(words []string) []string {
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}
	result := make([]string, len(wordMap))
	i := 0
	for item := range wordMap {
		result[i] = item
		i++
	}
	return result
}

func SetIntersection(left, right []string) []string {
	leftMap := make(map[string]bool)
	rightMap := make(map[string]bool)
	for _, item := range left {
		leftMap[item] = true
	}
	for _, item := range right {
		rightMap[item] = true
	}
	result := make([]string, 0, len(leftMap)+len(rightMap))
	for item := range leftMap {
		if rightMap[item] {
			result = append(result, item)
		}
	}
	return result
}

func SetSubtract(left, right []string) []string {
	leftMap := make(map[string]bool)
	rightMap := make(map[string]bool)
	for _, item := range left {
		leftMap[item] = true
	}
	for _, item := range right {
		rightMap[item] = true
	}
	result := make([]string, 0, len(leftMap)+len(rightMap))
	for item := range leftMap {
		if !rightMap[item] {
			result = append(result, item)
		}
	}
	return result
}

func DetermineAllergenPossibilities(labels []IngredientLabel) map[string][]string {
	allergenToIngredients := make(map[string][]string)
	for _, label := range labels {
		for _, allergen := range label.Allergens {
			if possibleIngredients, ok := allergenToIngredients[allergen]; !ok {
				allergenToIngredients[allergen] = label.Ingredients
			} else {
				allergenToIngredients[allergen] = SetIntersection(possibleIngredients, label.Ingredients)
			}
		}
	}
	return allergenToIngredients
}

func Part1(input string) int {
	labels := ReadInput(input)
	totalIngredients := make([]string, 0, 1000)
	totalAllergens := make([]string, 0, 1000)
	for _, label := range labels {
		totalIngredients = append(totalIngredients, label.Ingredients...)
		totalAllergens = append(totalAllergens, label.Allergens...)
	}

	allergenPossibilities := DetermineAllergenPossibilities(labels)
	allergenIngredients := make([]string, 0, 100)
	for _, aIngredients := range allergenPossibilities {
		allergenIngredients = append(allergenIngredients, aIngredients...)
	}
	allergenIngredients = UniqueStrings(allergenIngredients)

	nonAllergenIngredients := SetSubtract(totalIngredients, allergenIngredients)
	count := 0
	for _, label := range labels {
		noticed := SetIntersection(nonAllergenIngredients, label.Ingredients)
		count += len(noticed)
	}

	return count
}

func DetermineAllergenActual(possibilities map[string][]string) map[string]string {
	actualAllergens := make(map[string]string)

	for len(possibilities) > 0 {
		for allergen, ingredients := range possibilities {
			if len(ingredients) == 1 {
				actualAllergens[allergen] = ingredients[0]
				delete(possibilities, allergen)
				for a2, i2 := range possibilities {
					possibilities[a2] = SetSubtract(i2, ingredients)
				}
			} else if len(ingredients) == 0 {
				panic("ingredients list should not be empty")
			}
		}
	}

	return actualAllergens
}

func Part2(input string) string {
	labels := ReadInput(input)
	totalIngredients := make([]string, 0, 1000)
	totalAllergens := make([]string, 0, 1000)
	for _, label := range labels {
		totalIngredients = append(totalIngredients, label.Ingredients...)
		totalAllergens = append(totalAllergens, label.Allergens...)
	}
	totalAllergens = UniqueStrings(totalAllergens)
	sort.Strings(totalAllergens)

	allergenPossibilities := DetermineAllergenPossibilities(labels)
	actualAllergens := DetermineAllergenActual(allergenPossibilities)
	result := make([]string, len(totalAllergens))

	for i, allergen := range totalAllergens {
		result[i] = actualAllergens[allergen]
	}
	return strings.Join(result, ",")
}
