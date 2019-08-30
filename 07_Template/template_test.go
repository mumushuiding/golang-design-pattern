package template

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {

	toFriend := &holidayCardToFriend{}
	toFriend.ICard = toFriend
	toFriend.cardWriting()
	fmt.Println("---------------------")
	toCompany := new(holidayCardToCompany)
	toCompany.ICard = toCompany
	toCompany.cardWriting()

	// fmt.Println()
	// fmt.Println()

	// var card ICard
	// card = new(holidayCardToCompany)
	// card.cardWriting()

	// card = new(holidayCardToFriend)
	// card.cardWriting()

}
