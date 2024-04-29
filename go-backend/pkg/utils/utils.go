package utils

import (
	"math/rand"
	"time"
)

const (
	//TODO: Make it automated and generate random colors for more users...
    Turquoise = "#1ABC9C"
    Orange    = "#E67E2A"
    Red       = "#E92750"
    Blue      = "#3498DB"
    Green     = "#2ECC71"
    Yellow    = "#F1C40F"
    Purple    = "#9B59B6"
)

var colors = []string{Turquoise, Orange, Red, Blue, Green, Yellow, Purple}
var shuffledColors = make([]string, len(colors))
var index = 0
var rng *rand.Rand

func init() {
    rng = rand.New(rand.NewSource(time.Now().UnixNano()))
    shuffleColors()
}

func shuffleColors() {
    copy(shuffledColors, colors)

    rng.Shuffle(len(shuffledColors), func(i, j int) {
        shuffledColors[i], shuffledColors[j] = shuffledColors[j], shuffledColors[i]
    })
}

func GetRandomColor() string {
    if index >= len(colors) {
        shuffleColors()
        index = 0
    }

    color := shuffledColors[index]
    index++
    return color
}
