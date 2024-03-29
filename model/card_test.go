package Model

import (
	"testing"
)

var (
)

func TestGetDeckValue(t *testing.T) {
	msg := "Value must be %d, got %d"

	card := Card{COLOR_HEART, 2}
	if card.GetDeckValue() != 0 {
		t.Errorf(msg, 0, card.GetDeckValue())
	}

	card = Card{COLOR_CLOVER, 2}
	if card.GetDeckValue() != 13 {
		t.Errorf(msg, 13, card.GetDeckValue())
	}

	card = Card{COLOR_DIAMOND, 2}
	if card.GetDeckValue() != 26 {
		t.Errorf(msg, 26, card.GetDeckValue())
	}

	card = Card{COLOR_SPADE, 2}
	if card.GetDeckValue() != 39 {
		t.Errorf(msg, 39, card.GetDeckValue())
	}

	card = Card{COLOR_SPADE, VALUE_ACE}
	if card.GetDeckValue() != 51 {
		t.Errorf(msg, 51, card.GetDeckValue())
	}

	card = Card{COLOR_HEART, VALUE_JACK}
	if card.GetDeckValue() != 9 {
		t.Errorf(msg, 9, card.GetDeckValue())
	}

	card = Card{COLOR_CLOVER, 7}
	if card.GetDeckValue() != 18 {
		t.Errorf(msg, 18, card.GetDeckValue())
	}
}

func TestMakeCardFromDeckValue(t *testing.T) {
	
	var card Card

	card = MakeCardFromDeckValue(0)
	if card.Color != COLOR_HEART || card.Value != 2 {
		t.Error("Failed 0")
	}

	card = MakeCardFromDeckValue(13)
	if card.Color != COLOR_CLOVER || card.Value != 2 {
		t.Errorf("Failed 13")
	}

	card = MakeCardFromDeckValue(26)
	if card.Color != COLOR_DIAMOND || card.Value != 2 {
		t.Error("Failed 26")
	}

	card = MakeCardFromDeckValue(39)
	if card.Color != COLOR_SPADE || card.Value != 2 {
		t.Error("Failed 39")
	}

	card = MakeCardFromDeckValue(51)
	if card.Color != COLOR_SPADE || card.Value != VALUE_ACE {
		t.Error("Failed 51")
	}

	card = MakeCardFromDeckValue(9)
	if card.Color != COLOR_HEART || card.Value != VALUE_JACK {
		t.Error("Failed 51")
	}

	card = MakeCardFromDeckValue(18)
	if card.Color != COLOR_CLOVER || card.Value != 7 {
		t.Error("Failed 51")
	}
}