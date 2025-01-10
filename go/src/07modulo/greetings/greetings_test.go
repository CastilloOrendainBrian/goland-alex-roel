package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Juan")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, want %q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
