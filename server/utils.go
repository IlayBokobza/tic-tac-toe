package main

import (
	"math/rand"
	"time"
)

func createId() string {
	rand.Seed(time.Now().UnixNano())
	var id []byte

	for i := 0; i < 6; i++ {
		randInt := rand.Intn(122-97) + 97
		id = append(id, byte(randInt))
	}

	return string(id)
}
