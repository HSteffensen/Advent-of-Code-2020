package day22

import (
	"fmt"
	"strconv"
	"strings"
)

var Input = `Player 1:
3
42
4
25
14
36
32
18
33
10
35
50
16
31
34
46
9
6
41
7
15
45
30
27
49

Player 2:
8
11
47
21
17
39
29
43
23
28
13
22
5
20
44
38
26
37
2
24
48
12
19
1
40`

var ExampleInput1 = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func ReadInput(input string) CombatGame {
	blocks := strings.Split(input, "\n\n")
	decks := CombatGame{}
	for i, block := range blocks {
		cards := strings.Split(block, "\n")[1:]
		decks[i] = make(Deck, len(cards))
		for j, card := range cards {
			cardValue, err := strconv.Atoi(card)
			if err != nil {
				panic(fmt.Sprintf("Bad card value: expected integer, got %#v", card))
			}
			decks[i][j] = cardValue
		}
	}
	return decks
}
