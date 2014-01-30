package notes

import

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
"fmt"

type Duration byte

// English note names, A-G
type NoteName byte

const (
	C NoteName = 0
	D          = 2
	E          = 4
	F          = 5
	G          = 7
	A          = 9
	B          = 11
)

var noteNames = map[NoteName]string{
	A: "A",
	B: "B",
	C: "C",
	D: "D",
	E: "E",
	F: "F",
	G: "G",
}

func (n NoteName) String() string { return noteNames[n] }

type Accidental int8

const (
	DoubleFlat  Accidental = -2
	Flat                   = -1
	Natural                = 0
	Sharp                  = 1
	DoubleSharp            = 2
)

var accidentalNames = map[Accidental]string{
	DoubleFlat:  "B",
	Flat:        "b",
	Natural:     "",
	Sharp:       "#",
	DoubleSharp: "X",
}

func (a Accidental) String() string { return accidentalNames[a] }

type Interval int32

type Pitch struct {
	Name   NoteName
	Octave int8
	Acc    Accidental
}

func (p *Pitch) String() string {
	return fmt.Sprintf("%v%v%v", p.Name, p.Acc, p.Octave)
}

func (p *Pitch) num() int32 {
	return int32(p.Name) + int32(p.Acc) + 12*int32(p.Octave)
}

func (p *Pitch) Interval(q *Pitch) Interval {
	return Interval(q.num() - p.num())
}

type Note struct {
	Duration
	Pitch
}
