package auth_test

import (
	"testing"

	"github.com/atbu/slate/backend/auth"
)

func TestHashPassword(t *testing.T) {
	password := "aG{n5~[`eo4|3`FG"

	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if len(hash) == 0 {
		t.Errorf("Hash is empty")
	}

	if err := auth.VerifyPassword(hash, password); err != nil {
		t.Errorf("Expected hash %v for password %v", hash, password)
	}

	password = "password"
	if err := auth.VerifyPassword(hash, password); err == nil {
		t.Errorf("Expected hash %v to not match incorrect password %v", hash, password)
	}
}
