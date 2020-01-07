package package1

import (
	"fmt"
)

// Log stuff to standard output
func Log(s fmt.Stringer) {
	// to
	fmt.Printf("logging: %s\n", s)
}
