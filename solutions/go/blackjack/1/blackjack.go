package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    numCard := 0
	switch card {
        case "one":
                numCard = 1
        case "two":
                numCard = 2 
        case "three":
                numCard = 3
            case "four":
                numCard = 4
        case "five":
                numCard = 5 
        case "six":
                numCard = 6
        case "seven":
                numCard = 7
        case "eight":
                numCard = 8 
        case "nine":
                numCard = 9
        case "ten":
                numCard = 10
        case "jack":
                numCard = 10 
        case "queen":
                numCard = 10
        case "king":
                numCard = 10
        case "ace":
        	numCard = 11
        default:
            numCard = 0
	} 
    return numCard
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    value1 := ParseCard(card1)
    value2 := ParseCard(card2)
    sumValue := value1 + value2
    valueDealer := ParseCard(dealerCard)
	switch{
        case card1 == "ace" && card2 == "ace":
        	return "P"
        case sumValue == 21 && valueDealer != 10 && dealerCard != "ace":
        	return "W"
        case sumValue >= 17 && sumValue <= 20:
        	return "S"
        case sumValue >= 12 && sumValue <= 16 && valueDealer >=7:
        	return "H"
        case sumValue <= 11 :
        	return "H"
        default:
        	return "S"
    }
}
