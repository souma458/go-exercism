package dndcharacter

import (
	"math"
	"math/rand"
)

const minDiceRoll = 1
const maxDiceRoll = 6

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))

}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	min := maxDiceRoll
	total := 0
	for i := 0; i < 4; i++ {
		diceRoll := throwDice()
		total += diceRoll
		if diceRoll < min {
			min = diceRoll
		}
	}

	total -= min
	return total
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	newCharacter := &Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	newCharacter.Hitpoints = 10 + Modifier(newCharacter.Constitution)

	return *newCharacter
}

func throwDice() int {
	return minDiceRoll + rand.Intn(maxDiceRoll-minDiceRoll)
}
