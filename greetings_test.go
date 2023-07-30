package grettings

import (
	"regexp"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Merlote"
	wants := regexp.MustCompile(`\b` + name + `\b`) //must match
	got, err := Greet(name)

	if !wants.MatchString(got) || err != nil { //if not match
		//t.Errorf("got %q want %q", got, wants)
		t.Fatalf("funtion Great() = %q, %v, want match for %#q, nil", got, err, wants)
	}
}

func TestGreetEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" || err == nil {
		t.Fatalf("funtion Greet() = %q, %v, want empty string or nil", msg, err)
	}
}
