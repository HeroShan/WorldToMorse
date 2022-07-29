package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"strings"
	"time"
)

func Voice(fd string) {
	f, err := os.Open(fd)
	defer f.Close()
	if err != nil {
		report(err)
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/20))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	return
}

func (m *MorseWay) Dot() {
	fmt.Printf("·")
	Voice("./b.wav")
}

func (m *MorseWay) Dash() {
	fmt.Printf("-")
	Voice("./a.wav")
}

func (m *MorseWay) Pause() {
	fmt.Printf(m.PauseWat)
	time.Sleep(300 * time.Millisecond)
}

const (
	A  = "A"
	B  = "B"
	C  = "C"
	D  = "D"
	E  = "E"
	F  = "F"
	G  = "G"
	H  = "H"
	I  = "I"
	J  = "J"
	K  = "K"
	L  = "L"
	M  = "M"
	N  = "N"
	O  = "O"
	P  = "P"
	Q  = "Q"
	R  = "R"
	S  = "S"
	T  = "T"
	U  = "U"
	V  = "V"
	W  = "W"
	X  = "x"
	Y  = "Y"
	Z  = "Z"
	a  = "a"
	b  = "b"
	c  = "c"
	d  = "d"
	e  = "e"
	f  = "f"
	g  = "g"
	h  = "h"
	i  = "i"
	j  = "j"
	k  = "k"
	l  = "l"
	lm = "m"
	n  = "n"
	o  = "o"
	p  = "p"
	q  = "q"
	r  = "r"
	s  = "s"
	t  = "t"
	u  = "u"
	v  = "v"
	w  = "w"
	x  = "x"
	y  = "y"
	z  = "z"
)

type MorseWay struct {
	DotWay   string
	DashWay  string
	PauseWat string
	Col      Collect
}

type Collect map[string]string

func (m *MorseWay) InitData() {
	m.DotWay = "·"
	m.DashWay = "-"
	m.PauseWat = "  "
	m.Col = make(map[string]string, 0)
	m.Col[a] = m.DotWay + m.DashWay + m.PauseWat //。-
	m.Col[A] = m.DotWay + m.DashWay + m.PauseWat
	m.Col[b] = m.DashWay + m.DotWay + m.DotWay + m.DotWay + m.PauseWat //- 。。。
	m.Col[B] = m.DashWay + m.DotWay + m.DotWay + m.DotWay + m.PauseWat
	m.Col[c] = m.DashWay + m.DotWay + m.DashWay + m.DotWay + m.PauseWat //- 。- 。
	m.Col[C] = m.DashWay + m.DotWay + m.DashWay + m.DotWay + m.PauseWat
	m.Col[d] = m.DashWay + m.DotWay + m.DotWay + m.PauseWat //- 。。
	m.Col[D] = m.DashWay + m.DotWay + m.DotWay + m.PauseWat
	m.Col[e] = m.DotWay + m.PauseWat //。
	m.Col[E] = m.DotWay + m.PauseWat
	m.Col[f] = m.DotWay + m.DotWay + m.DashWay + m.DotWay + m.PauseWat //。。- 。
	m.Col[F] = m.DotWay + m.DotWay + m.DashWay + m.DotWay + m.PauseWat
	m.Col[g] = m.DashWay + m.DashWay + m.DotWay + m.PauseWat //- - 。
	m.Col[G] = m.DashWay + m.DashWay + m.DotWay + m.PauseWat
	m.Col[h] = m.DotWay + m.DotWay + m.DotWay + m.DotWay + m.PauseWat //。。。。
	m.Col[H] = m.DotWay + m.DotWay + m.DotWay + m.DotWay + m.PauseWat
	m.Col[i] = m.DotWay + m.DotWay + m.PauseWat //。。
	m.Col[I] = m.DotWay + m.DotWay + m.PauseWat
	m.Col[j] = m.DotWay + m.DashWay + m.DashWay + m.PauseWat //。 - -
	m.Col[J] = m.DotWay + m.DashWay + m.DashWay + m.PauseWat
	m.Col[k] = m.DashWay + m.DotWay + m.DashWay + m.PauseWat //-。-
	m.Col[K] = m.DashWay + m.DotWay + m.DashWay + m.PauseWat
	m.Col[l] = m.DotWay + m.DashWay + m.DotWay + m.DotWay + m.PauseWat //。-。。
	m.Col[L] = m.DotWay + m.DashWay + m.DotWay + m.DotWay + m.PauseWat
	m.Col[lm] = m.DashWay + m.DashWay + m.PauseWat //- -
	m.Col[M] = m.DashWay + m.DashWay + m.PauseWat
	m.Col[n] = m.DashWay + m.DotWay + m.PauseWat //-。
	m.Col[N] = m.DashWay + m.DotWay + m.PauseWat
	m.Col[o] = m.DashWay + m.DashWay + m.DashWay + m.PauseWat //- -
	m.Col[O] = m.DashWay + m.DashWay + m.DashWay + m.PauseWat
	m.Col[p] = m.DotWay + m.DashWay + m.DashWay + m.DotWay + m.PauseWat //。 - - 。
	m.Col[P] = m.DotWay + m.DashWay + m.DashWay + m.DotWay + m.PauseWat
	m.Col[q] = m.DashWay + m.DashWay + m.DotWay + m.DashWay + m.PauseWat //- - 。-
	m.Col[Q] = m.DashWay + m.DashWay + m.DotWay + m.DashWay + m.PauseWat
	m.Col[r] = m.DotWay + m.DashWay + m.DotWay + m.PauseWat //。- 。
	m.Col[R] = m.DotWay + m.DashWay + m.DotWay + m.PauseWat
	m.Col[s] = m.DotWay + m.DotWay + m.DotWay + m.PauseWat //。。。。
	m.Col[S] = m.DotWay + m.DotWay + m.DotWay + m.PauseWat
	m.Col[t] = m.DashWay + m.PauseWat //-
	m.Col[T] = m.DashWay + m.PauseWat
	m.Col[u] = m.DotWay + m.DotWay + m.DashWay + m.PauseWat //。。-
	m.Col[U] = m.DotWay + m.DotWay + m.DashWay + m.PauseWat
	m.Col[v] = m.DotWay + m.DotWay + m.DotWay + m.DashWay + m.PauseWat //。。。-
	m.Col[V] = m.DotWay + m.DotWay + m.DotWay + m.DashWay + m.PauseWat
	m.Col[w] = m.DotWay + m.DashWay + m.DashWay + m.PauseWat //。 - -
	m.Col[W] = m.DotWay + m.DashWay + m.DashWay + m.PauseWat
	m.Col[x] = m.DashWay + m.DotWay + m.DashWay + m.DashWay + m.PauseWat //-。 - -
	m.Col[X] = m.DashWay + m.DotWay + m.DashWay + m.DashWay + m.PauseWat
	m.Col[y] = m.DashWay + m.DotWay + m.DashWay + m.DashWay + m.PauseWat //-。- -
	m.Col[Y] = m.DashWay + m.DotWay + m.DashWay + m.DashWay + m.PauseWat
	m.Col[z] = m.DashWay + m.DashWay + m.DotWay + m.DotWay + m.PauseWat //- -
	m.Col[Z] = m.DashWay + m.DashWay + m.DotWay + m.DotWay + m.PauseWat

}

type Morse interface {
	Dot()
	Dash()
	Pause()
	InitData()
}

func (m *MorseWay) Explain(text string) {
	for _, v := range strings.Split(text, "") {
		m.Ticker(m.Col[v])
	}
	fmt.Println()
}

func (m *MorseWay) Ticker(morse string) {
	for _, v := range strings.Split(morse, "") {
		switch v {
		case "·":
			m.Dot()
			break
		case "-":
			m.Dash()
			break
		default:
			break
		}
	}
	m.Pause()
}

func main() {
	fmt.Println("请输入a-z A-Z ：")
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	m := new(MorseWay)
	m.InitData()
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			m.Explain(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func report(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
