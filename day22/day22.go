package day22

import (
	"fmt"
	"strings"
)

type Deck []int
type CombatGame [2]Deck

func (d *Deck) PlayTopCard() int {
	top := (*d)[0]
	copy(*d, (*d)[1:])
	*d = (*d)[:len(*d)-1]
	return top
}

func (d *Deck) PutOnBottom(a, b int) {
	if a < b {
		a, b = b, a
	}
	*d = append(*d, a, b)
}

func (d Deck) String() string {
	dString := make([]string, len(d))
	for i, card := range d {
		dString[i] = fmt.Sprint(card)
	}
	return strings.Join(dString, ", ")
}

func (cg CombatGame) Winner() int {
	for i, deck := range cg {
		if len(deck) == 0 {
			return 1 - i
		}
	}
	return -1
}

func (cg *CombatGame) Step() int {
	card0, card1 := cg[0].PlayTopCard(), cg[1].PlayTopCard()
	stepWinner := 0
	if card1 > card0 {
		stepWinner = 1
	}
	cg[stepWinner].PutOnBottom(card0, card1)

	return cg.Winner()
}

func (cg *CombatGame) PlayGame() int {
	for {
		if winner := cg.Winner(); winner >= 0 {
			return winner
		}
		cg.Step()
	}
}

func (cg *CombatGame) RecursiveStep(gameCounter *int) int {
	card0, card1 := cg[0].PlayTopCard(), cg[1].PlayTopCard()
	// fmt.Printf("Player 1 plays: %v\nPlayer 2 plays: %v\n", card0, card1)
	stepWinner := 0
	if len(cg[0]) >= card0 && len(cg[1]) >= card1 {
		newGame := CombatGame{{}, {}}
		newGame[0] = make(Deck, card0)
		copy(newGame[0], cg[0][:card0])
		newGame[1] = make(Deck, card1)
		copy(newGame[1], cg[1][:card1])
		// fmt.Println("Playing a sub-game to determine the winner...")
		stepWinner = newGame.PlayRecursiveGame(gameCounter)
		// fmt.Println("...Sub-game complete!")
		cards := Deck{card0, card1}
		cg[stepWinner] = append(cg[stepWinner], cards[stepWinner], cards[1-stepWinner])
	} else {
		if card1 > card0 {
			stepWinner = 1
		}
		cg[stepWinner].PutOnBottom(card0, card1)
	}

	return stepWinner
}

func (cg *CombatGame) PlayRecursiveGame(gameCounter *int) int {
	*gameCounter++
	// gameID := *gameCounter
	visitedGameStates := make(map[string]bool)
	// fmt.Printf("\n=== Game %v ===\n", gameID)
	for round := 1; ; round++ {
		// fmt.Printf("\n-- Round %v (Game %v) --\nPlayer 1's deck: %v\nPlayer 2's deck: %v\n", round, gameID, cg[0], cg[1])
		gameState := fmt.Sprint(*cg)
		if visitedGameStates[gameState] {
			// fmt.Printf("\n...We've already been here before...\nPlayer 1 wins by gameloop rule!\n")
			return 0
		}
		visitedGameStates[gameState] = true
		// roundWinner := cg.RecursiveStep(gameCounter)
		cg.RecursiveStep(gameCounter)
		// fmt.Printf("Player %v wins round %v of game %v!\n", roundWinner+1, round, gameID)
		if winner := cg.Winner(); winner >= 0 {
			// fmt.Printf("The winner of game %v is player %v!\n", gameID, winner+1)
			return winner
		}
	}
}

func Score(deck Deck) int {
	factor := len(deck)
	total := 0
	for i, card := range deck {
		total += card * (factor - i)
	}
	return total
}

func Part1(input string) int {
	game := ReadInput(input)
	winner := game.PlayGame()
	return Score(game[winner])
}

func Part2(input string) int {
	game := ReadInput(input)
	gameCounter := 0
	winner := game.PlayRecursiveGame(&gameCounter)
	// fmt.Printf("== Post-game results ==\nPlayer 1's deck: %v\nPlayer 2's deck: %v\n", game[0], game[1])
	return Score(game[winner])
}
