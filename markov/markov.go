package markov

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Value int8

type suite []Value

func (m suite) hash() string {
	b := make([]byte, len(m))
	for i := range m {
		b[i] = byte(m[i])
	}
	return string(b)
}

type Chain struct {
	after map[string]suite
}

func NewChain() *Chain {
	return &Chain{make(map[string]suite)}
}

func (c *Chain) Feed(vs ...Value) {
	for i, d := range vs {
		for j := 0; j <= i; j++ {
			h := suite(vs[j:i]).hash()
			c.after[h] = suite(append(c.after[h], d))
		}
	}
}

func (c *Chain) String() string {
	w := new(bytes.Buffer)
	for k, v := range c.after {
		fmt.Fprintf(w, "%v\t: %v\n", []byte(k), v)
	}
	return w.String()
}

func (c *Chain) After(vs ...Value) []Value {
	return c.after[suite(vs).hash()]
}

func (c *Chain) Generate(length, memory int) []Value {
	res := make(suite, 0, length)
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
