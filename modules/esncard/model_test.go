package esncard

import (
	"log"
	"testing"
)

func setUpModel() {
	NewESNcard(&ESNcard{Id_adherent: 2, Esncard: "Mario"})
	NewESNcard(&ESNcard{Id_adherent: 3, Esncard: "Pikachu"})
	NewESNcard(&ESNcard{Id_adherent: 4, Esncard: "AveryLittleCode"})
}

func TestNewESNcard(t *testing.T) {
	NewESNcard(&ESNcard{Id_adherent: 1, Esncard: "Peach"})
}

func TestFindESNcardByIdadherent(t *testing.T) {
	card := FindESNcardByIdadherent(4)

	if card.Id_adherent != 4 {
		log.Fatal("ESNcard is not the good one.")
	}
}

func TestFindESNcardByESNcard(t *testing.T) {
	card := FindESNcardByESNcard("AveryLittleCode")

	if card.Esncard != "AveryLittleCode" {
		log.Fatal("ESNcard is not the good one.")
	}
}

func TestAllESNcards(t *testing.T) {
	cards := AllESNcards()

	//log.Println(adhs)
	if len(*cards) == 0 {
		log.Fatal("ESNcard is empty")
	}
}

func TestDeleteESNcardById(t *testing.T) {
	err := DeleteESNcardById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	cards := AllESNcards()

	for _, card := range *cards {
		if card.Esncard == "AveryLittleCode" {
			log.Fatal("Card_3 didn't be removed !")
		}
	}
}
