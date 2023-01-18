package greetings

import (
    "testing"
    "regexp"
)

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if (msg != "" || err == nil) {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
    }
}