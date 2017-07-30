package stringutil

import "testing"

func TestReverseEquality(t *testing.T) {
    // inline definition of an array of an anonymous struct with two string members: in want want
    // then it inline defines the set of members of the array
    cases := [] struct {
        in, want string
    } {
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    // this is iterating over the range of the members of cases
    // got = the reverse of the in member of the given struct
    // got is compared to the want member of the struct (the correct output). 
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}