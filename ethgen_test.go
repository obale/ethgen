package main

import (
	"regexp"
	"testing"
)

func TestGenerateAddress(t *testing.T) {
	var (
		addr    string
		privKey string
		err     error
		match   bool
	)

	if addr, privKey, err = GenerateAddress(); err != nil {
		t.Error(err)
	}

	if match, err = regexp.MatchString("0x.*", addr); err != nil {
		t.Error(err)
	}
	if !match {
		t.Errorf("address is not valid: %v", err)
	}

	var privKeyLength = 64
	if len(privKey) != privKeyLength {
		t.Errorf("private key is not valid, length is %d instead of %d", len(privKey), privKeyLength)
	}
}
