package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

func main() {
	length := flag.Int("n", 32, "number of random bytes to generate")
	flag.Parse()

	if *length <= 0 {
		log.Fatalf("length must be positive, got %d", *length)
	}

	buf := make([]byte, *length)
	if _, err := rand.Read(buf); err != nil {
		log.Fatalf("rand.Read failed: %v", err)
	}

	fmt.Println(hex.EncodeToString(buf))
}
