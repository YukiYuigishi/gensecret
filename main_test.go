package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
)

func TestGenerateSecretEncodings(t *testing.T) {
	input := []byte{0, 1, 2, 3, 4, 5, 250, 251, 252, 253, 254, 255}
	reader := bytes.NewReader(input)

	got, err := generateSecret(len(input), "hex", reader)
	if err != nil {
		t.Fatalf("hex generateSecret returned error: %v", err)
	}
	want := hex.EncodeToString(input)
	if got != want {
		t.Fatalf("hex output mismatch: got %q want %q", got, want)
	}

	reader = bytes.NewReader(input)
	got, err = generateSecret(len(input), "base64", reader)
	if err != nil {
		t.Fatalf("base64 generateSecret returned error: %v", err)
	}
	want = base64.StdEncoding.EncodeToString(input)
	if got != want {
		t.Fatalf("base64 output mismatch: got %q want %q", got, want)
	}

	reader = bytes.NewReader(input)
	got, err = generateSecret(len(input), "base64url", reader)
	if err != nil {
		t.Fatalf("base64url generateSecret returned error: %v", err)
	}
	want = base64.RawURLEncoding.EncodeToString(input)
	if got != want {
		t.Fatalf("base64url output mismatch: got %q want %q", got, want)
	}
}

func TestGenerateSecretValidation(t *testing.T) {
	_, err := generateSecret(0, "hex", bytes.NewReader([]byte{1}))
	if err == nil || !strings.Contains(err.Error(), "length must be positive") {
		t.Fatalf("expected length validation error, got %v", err)
	}

	_, err = generateSecret(1, "nope", bytes.NewReader([]byte{1}))
	if err == nil || !strings.Contains(err.Error(), "unknown encoding") {
		t.Fatalf("expected unknown encoding error, got %v", err)
	}
}
