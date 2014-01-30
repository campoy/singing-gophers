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

	s, err := newSine(150, 44000)
	if err != nil {
		log.Fatalf("creating: %v", err)
	}
	defer s.Close()

	if err := s.Start(); err != nil {
		log.Fatalf("starting: %v", err)
	}

	// start notes OMIT
	swap := 0.
	for i := 0; i <= 24; i++ {
		s.freq *= math.Pow(2, 1./12.)
		time.Sleep(100 * time.Millisecond)

		swap, s.freq = s.freq, swap
		time.Sleep(150 * time.Millisecond)
		swap, s.freq = s.freq, swap
	}

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
