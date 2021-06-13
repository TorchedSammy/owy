package owy

import (
	"fmt"
	"time"
)

type Spinner struct {
	Text string
	style Style
	stop chan bool
}

func NewSpinner(text string, style *Style) *Spinner {
	if style == nil {
		style = &LineStyle
	}

	return &Spinner{
		Text: text,
		style: *style,
		stop: make(chan bool, 1),
	}
}

func (s *Spinner) Start() {
	fmt.Printf("\033[?25l")
	go func() {
		for {
			for i := 0; i < len(s.style.Stages); i++ {
				select {
				case <-s.stop:
					return
				default:
					fmt.Printf("\r%s%s\033[0m %s ", "\033[34m", s.style.Stages[i], s.Text)
					time.Sleep(time.Duration(s.style.Interval) * time.Millisecond)
				}
			}
		}
	}()
}

func (s *Spinner) Stop() {
	fmt.Println("\r\033[2K")
	fmt.Printf("\033[?25h")
	s.stop <- true
}

