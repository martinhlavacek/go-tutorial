package main

import (
	"testing"
	"fmt"
)

func TestNewDeck(t *testing.T){
	d := newDeck()
	fmt.Println(len(d))
	if len(d) != 4 {
		t.Errorf("error d: %v", len(d))
	}

	if d[0] != "one of a"{
		t.Errorf("excepted one of a but we got: %v",d[0])
	}

	if d[len(d) - 1] != "two of b"{
		t.Errorf("excepted two of b but we got: %v",d[len(d) -1])
	}
}