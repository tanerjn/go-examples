package main

import (
    "testing"
    "unicode/utf8"
)

func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev, revErr := Reverse(tc.in)
        if rev != tc.want {
	    t.Errorf("Reverse: %q, want %q, revErr: %v", rev, tc.want, revErr)
        }
    }
}
//Create seed corpus and fuzz test
func FuzzReverse(f *testing.F){
	testcases := []string{"Hello, world", " ", "12345", "aa32", "adsd", "adada", "adad"}
	for _, tc := range testcases {
		f.Add(tc) // f.Add to add seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)
		if revErr != nil{
		    t.Skip()
		    }
		doubleRev, doubleRevErr := Reverse(rev)
		if doubleRevErr != nil{
		    t.Skip()
		    }
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
	    t.Errorf("Before: %q, after: %q", orig, doubleRev )
	}
	if utf8.ValidString(orig) && !utf8.ValidString(rev) {
	    t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}


		

