package main

import (
	"time"

	"github.com/Rosettea/owy"
)

func main() {
	s := owy.NewSpinner("Hello world!", &owy.DotStyle)
	s.Start()
	time.Sleep(5 * time.Second)
	s.Stop()
}

