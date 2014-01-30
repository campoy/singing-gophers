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

	forms := map[string]func(*wave, [][]float32){
		"sine":     (*wave).sine,
		"square":   (*wave).square,
		"sawtooth": (*wave).sawtooth,
		"triangle": (*wave).triangle,
	}

	for name, form := range forms {
		log.Println(name)

		s, err := newWave(200, 44000, form)
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

		if err := s.Stop(); err != nil {
			log.Fatalf("stopping: %v", err)
		}

		time.Sleep(time.Second)

	}
}

type wave struct {
	*portaudio.Stream
	freq, sampleRate, phase float64
}

func newWave(freq, sampleRate float64, form func(*wave, [][]float32)) (*wave, error) {
	s := &wave{nil, freq, sampleRate, 0}
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, 0, call(s, form))
	if err != nil {
		return nil, err
	}
	s.Stream = stream
	return s, nil
}

func call(w *wave, f func(*wave, [][]float32)) func([][]float32) {
	return func(out [][]float32) {
		f(w, out)
	}
}

func (g *wave) sine(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phase))
		_, g.phase = math.Modf(g.phase + g.freq/g.sampleRate)
	}
}

func (g *wave) square(out [][]float32) {
	for i := range out[0] {
		if g.phase > 0.5 {
			out[0][i] = 1
		} else {
			out[0][i] = -1
		}
		_, g.phase = math.Modf(g.phase + g.freq/g.sampleRate)
	}
}

func (g *wave) sawtooth(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(2*(g.phase-math.Floor(g.phase)) - 1)
		_, g.phase = math.Modf(g.phase + g.freq/g.sampleRate)
	}
}

func (g *wave) triangle(out [][]float32) {
	for i := range out[0] {
		p := math.Floor((g.phase + 1) / 2)
		out[0][i] = float32(g.phase - 2*p*math.Pow(-1, p))
		_, g.phase = math.Modf(g.phase + g.freq/g.sampleRate)
	}
}
