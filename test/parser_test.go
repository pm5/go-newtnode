package test

import (
	".."
	"testing"
)

func TestNewParser(t *testing.T) {
	p := newtnode.NewParser()
	if p == nil {
		t.Fatalf("Parser creation failed")
	}
}
