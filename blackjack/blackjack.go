package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	total := card1Value + card2Value
	blackjack := 21

	switch {
	case card1Value == ParseCard("ace") && card2Value == ParseCard("ace"):
		return "P"
	case total == blackjack && (dealerCardValue != ParseCard("ace") && !isFigure(dealerCard) && dealerCardValue != ParseCard("ten")):
		return "W"
	case total == blackjack && (dealerCardValue == ParseCard("ace") || isFigure(dealerCard) || dealerCardValue == ParseCard("ten")):
		return "S"
	case total >= 17 && total <= 20:
		return "S"
	case total >= 12 && total <= 16 && dealerCardValue >= ParseCard("seven"):
		return "H"
	case total >= 12 && total <= 16 && dealerCardValue < ParseCard("seven"):
		return "S"
	default:
		return "H"
	}
}

func isFigure(card string) bool {
	switch card {
	case "jack":
		return true
	case "queen":
		return true
	case "king":
		return true
	default:
		return false
	}
}
