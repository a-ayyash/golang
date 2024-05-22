package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T)  {
  name := "Ayyash"
  want := regexp.MustCompile(`\b`+name+`\b`)
  msg, err := Hello("ayyash")
  if !want.MatchString(msg) || err != nil {
    t.Fatalf(`Hello("Ayyash") = %q, %v, want match for %#q, nil`, msg, err, want)
  }

  
}
