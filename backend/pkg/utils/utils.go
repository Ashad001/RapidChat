package utils

import (
	"math/rand"
	"time"
)

const (
	Turquoise = "#1ABC9C"
	Orange    = "#E67E2A"
	Red       = "#E92750"
	Blue      = "#3498DB"
	Green     = "#2ECC71"
	Yellow    = "#F1C40F"
	Purple    = "#9B59B6"
	Gray      = "#514652"
)

func GetRandomColor() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	colors := []string{Turquoise, Orange, Red, Blue, Green, Yellow, Purple, Gray}
	randomIndex := rng.Intn(len(colors))
	return colors[randomIndex]
}
