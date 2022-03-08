package main_test

import (
	"fmt"
	"testing"

	diff "github.com/Iiqbal2000/stringDissimilarity"
)

func TestGetDiffChar(t *testing.T) {
	cases := []struct{
		s, t string
		want string
	} {
		{s: "", t: "y", want: "y"},
		{s: "annqalff", t: "fqlnannaf", want: "n"},
		{s: "bxcn", t: "abncx", want: "a"},
		{s: "abk", t: "abkw", want: "w"},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test-%d", i+1), func(t *testing.T) {
			got := diff.GetDiffChar(c.s, c.t)
			if got != c.want {
				t.Errorf("Got %s, want %s", got, c.want)
			}
		})
	}
}