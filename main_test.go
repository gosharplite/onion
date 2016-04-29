package onion

import (
	"testing"
)

func TestOnion(t *testing.T) {

	cli := Decorate(Onion{},
		Add(3),
		Mul(4),
	)

	a, _ := cli.Do(2)

	t.Errorf("%v", a)
}
