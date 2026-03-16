// Package blackjack simulates the first turn of a Blackjack game.
// Note: Commonly, aces can take the value of 1 or 11 but for simplicity it will assume that they can only take the value of 11.
package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	result := -1
	switch card {
	case "ace":
		result = 11
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	case "ten", "jack", "king", "queen":
		result = 10
	default:
		result = 0
	}
	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	parse_card1 := ParseCard(card1)
	parse_card2 := ParseCard(card2)
	parse_dealer := ParseCard(dealerCard)
	sum_c12 := parse_card1 + parse_card2
	result := ""
	switch {
	case parse_card1 == 11 && parse_card2 == 11:
		result = "P"
	case sum_c12 == 21 && (parse_dealer != 11 && parse_dealer != 10):
		result = "W"
	case ((sum_c12 >= 12 && sum_c12 <= 16) && parse_dealer >= 7) || sum_c12 <= 11:
		result = "H"
	case parse_dealer == 11 || parse_dealer == 10 || (sum_c12 >= 17 && sum_c12 <= 20) || ((sum_c12 >= 12 && sum_c12 <= 16) && parse_dealer < 7):
		result = "S"
	}
	return result
}
