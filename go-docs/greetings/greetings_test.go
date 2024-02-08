package greetings

import (
	"regexp"
	"testing"
)

// to run tests, enter this command from same directory. -v flag for verbose.
// % go test -v
// happy path
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") =%q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// sad path
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
