package markov

import (
	"bytes"
	"fmt"
	"math/rand"
)

// rhythms are limited to 256 values
// Base rhythm is 1/36
// Negra: 36
// Corchea: 18
// Tresillo de corcheas: 12
// Semicorchea: 9
// Tresillo de semicorcheas: 6
// Blanca 72
// Redonda 144
// Max: 255 ->
// 255 - 144 (redonda)
// 111 - 72 (blanca)
// 39 - 36 (negra)
// 3 (tresillo de fusas)
type Duration byte

// A suite contains a suite of durations ... duh
type suite []Duration

func (ds suite) hash() hash {
	bs := make([]byte, len(ds))
	for i := range ds {
		bs[i] = byte(ds[i])
	}
	return hash(bs)
}

type hash string

type RythmChain struct {
	after map[hash][]Duration
}

func NewRhythmChain() *RythmChain {
	return &RythmChain{make(map[hash][]Duration)}
}

func (c *RythmChain) Feed(ds ...Duration) {
	for i, d := range ds {
		for j := 0; j <= i; j++ {
			h := suite(ds[j:i]).hash()
			c.after[h] = append(c.after[h], d)
		}
	}
}

func (c *RythmChain) String() string {
	w := new(bytes.Buffer)
	for k, v := range c.after {
		fmt.Fprintf(w, "%v\t: %v\n", []byte(k), v)
	}
	return w.String()
}

func (c *RythmChain) After(ds ...Duration) []Duration {
	return c.after[suite(ds).hash()]
}

func (c *RythmChain) Generate(length, memory int) []Duration {
	res := make([]Duration, 0, length)
	for i := 0; i < length; i++ {
		prefStart := len(res) - memory
		if prefStart < 0 {
			prefStart = 0
		}

		for ; prefStart <= len(res); prefStart++ {
			ops := c.After(res[prefStart:]...)
			if len(ops) == 0 {
				continue
			}
			choice := rand.Intn(len(ops))
			res = append(res, ops[choice])
			break
		}
	}
	return res
}
