package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	words := strings.Fields("400 500 600 700 800 900")

	fmt.Println(words)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	fmt.Println("after", words)
}
