// Package stringutil contains utility functions for working with strings
package stringutil

// function definition: func <name> (<argn> <argn_type> ...) <return_typen> { ... }
// reverses a string and returns a copy in reverse
func Reverse(s string) string {
    // rune is a function that operates on strings and returns a unicode array
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}