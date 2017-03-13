package main

import (
	"strconv"
    "fmt"
)

type Card struct {
	Suit, Rank string
}

func (h Card)IsBetterThan(c Card) bool {
	face_values := make(map[string]int)
	var h_value, c_value int

	face_values["A"] = 1
	face_values["J"] = 11
	face_values["Q"] = 12
	face_values["K"] = 13

	if _, err := strconv.Atoi(h.Rank); err == nil {
		h_value,err = strconv.Atoi(h.Rank)
	} else {
		h_value = face_values[h.Rank]
	}

	if _, err := strconv.Atoi(c.Rank); err == nil {
		c_value,err = strconv.Atoi(c.Rank)
	} else {
		c_value = face_values[c.Rank]
	}

	return h_value > c_value;
}

func (h Card)String() string {
	return fmt.Sprintf("%s of %s", h.Rank, h.Suit)
}
