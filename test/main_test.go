package test

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if "Hello World" == "Hello World" {
		t.Errorf("Expected Hello World")
	}
}
