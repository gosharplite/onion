package onion

import (
	"testing"
)

func TestOnion(t *testing.T) {

	cli := Decorate(Onion{},
		Add(3),
		Mul(5),
	)

	a, _ := cli.Do(2)

	if a != 13 {
		t.Errorf("%v", a)
	}
}
