package main

import (
	"log"
	"math"
	"time"

	"code.google.com/p/portaudio-go/portaudio"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	s, err := newSine(200, 44000)
	if err != nil {
		log.Fatalf("creating: %v", err)
	}
	defer s.Close()

	if err := s.Start(); err != nil {
		log.Fatalf("starting: %v", err)
	}

	notes := []float64{0, 0, 0, -4, +3, 0, -4, +3, 0}
	length := []time.Duration{4, 4, 4, 3, 1, 4, 3, 1, 4}

	for i, n := range notes {
		s.freq = 220 * math.Pow(2, n/12)
		time.Sleep(length[i] * 150 * time.Millisecond)

		s.freq = 0
		time.Sleep(10 * time.Millisecond)
	}
	// end_notes OMIT
	if err := s.Stop(); err != nil {
		log.Fatalf("stopping: %v", err)
	}
}

type sine struct {
	*portaudio.Stream
	freq, sampleRate, phase float64
}

func newSine(freq, sampleRate float64) (*sine, error) {
	s := &sine{nil, freq, sampleRate, 0}
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, 0, s.processAudio)
	if err != nil {
		return nil, err
	}
	s.Stream = stream
	return s, nil
}

func (g *sine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phase))
		_, g.phase = math.Modf(g.phase + g.freq/g.sampleRate)
	}
}
