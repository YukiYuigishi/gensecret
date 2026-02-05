package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
)

func generateSecret(length int, encoding string, r io.Reader) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive, got %d", length)
	}
	if r == nil {
		return "", errors.New("reader must not be nil")
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", fmt.Errorf("read failed: %w", err)
	}

	switch strings.ToLower(encoding) {
	case "hex":
		return hex.EncodeToString(buf), nil
	case "base64":
		return base64.StdEncoding.EncodeToString(buf), nil
	case "base64url":
		return base64.RawURLEncoding.EncodeToString(buf), nil
	default:
		return "", fmt.Errorf("unknown encoding %q (expected: hex|base64|base64url)", encoding)
	}
}

func main() {
	length := flag.Int("n", 32, "number of random bytes to generate")
	encoding := flag.String("enc", "hex", "output encoding: hex|base64|base64url")
	flag.Parse()

	out, err := generateSecret(*length, *encoding, rand.Reader)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(out)
}
