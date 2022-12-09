package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

const (
	width  = 20
	height = 20

	maxTurns = 200
)

var (
	snake    = []int{width / 2, height / 2}
	fruit    = []int{rand.Intn(width), rand.Intn(height)}
	dir      = "RIGHT"
	score    = 0
	gameOver = false
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	encoding.Register()
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Println("An error occurred:", e)
		return
	}

	if e = s.Init(); e != nil {
		fmt.Println("An error occurred:", e)
		return
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()

	for !gameOver {
		draw(s)

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'w':
					dir = "UP"
				case 's':
					dir = "DOWN"
				case 'a':
					dir = "LEFT"
				case 'd':
					dir = "RIGHT"
				}
			case tcell.KeyEscape:
				gameOver = true
			}
		}

		move()

		if gameOver {
			fmt.Println("Game over! Your score was", score)
		}
	}

	s.Fini()
}

func draw(s tcell.Screen) {
	s.Clear()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == fruit[0] && y == fruit[1] {
				s.SetContent(x, y, 'F', nil, tcell.StyleDefault)
				continue
			}

			if x == snake[0] && y == snake[1] {
				s.SetContent(x, y, 'S', nil, tcell.StyleDefault)
				continue
			}

			s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		}
	}

	s.Show()
}

func move() {
	switch dir {
	case "UP":
		snake[1]--
	case "DOWN":
		snake[1]++
	case "LEFT":
		snake[0]--
	case "RIGHT":
		snake[0]++
	}

	if snake[0] < 0 || snake[0] >= width || snake[1] < 0 || snake[1] >= height {
		gameOver = true
		return
	}

	if snake[0] == fruit[0] && snake[1] == fruit[1] {
		score++
		fruit = []int{rand.Intn(width), rand.Intn(height)}
		return
	}

	gameOver = maxTurns <= 0
}
