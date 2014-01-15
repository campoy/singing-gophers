package markov

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestNewRhythmChain(t *testing.T) {
	c := NewRhythmChain()
	c.Feed(1, 2, 3, 4, 5)
	c.Feed(5, 4, 3, 2, 1)

	next := c.After(3)
	if len(next) != 2 {
		t.Errorf("expected two options, got %v", len(next))
	}

	if next[0] > next[1] {
		next[0], next[1] = next[1], next[0]
	}
	exp := []Duration{2, 4}
	if !reflect.DeepEqual(next, exp) {
		t.Errorf("expected options %v, got %v", exp, next)
	}
}

func TestGenerate(t *testing.T) {
	c := NewRhythmChain()
	rand.Seed(time.Now().Unix())
	c.Feed(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	c.Feed(0, 2, 4, 6, 8, 0)
	c.Feed(0, 3, 6, 9, 0)

	fmt.Println(c.Generate(100, 0))
}
